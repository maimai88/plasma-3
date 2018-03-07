package plasma

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/whisper/shhclient"
	"github.com/ethereum/go-ethereum/whisper/whisperv5"
)

var (
	plasmaSymKey = crypto.Keccak256(crypto.Keccak256([]byte("randomkeyhere")))
)

type Backend struct {
	// plasma contract address
	address common.Address
	// FIXME for convenience it is fine
	isAuthority bool

	stateMu sync.Mutex
	quit    chan struct{}
	wg      sync.WaitGroup
	chain   *Chain
	network NetworkClient
}

func NewBackend(address common.Address, isAuthority bool) *Backend {
	return &Backend{
		address:     address,
		isAuthority: isAuthority,
	}
}

func (b *Backend) Start(shh *shhclient.Client, backend bind.ContractBackend) error {
	b.stateMu.Lock()
	defer b.stateMu.Unlock()
	contract, err := NewPlasma(b.address, backend)
	if err != nil {
		return err
	}
	symID, err := shh.AddSymmetricKey(context.TODO(), []byte(plasmaSymKey))
	if err != nil {
		return err
	}
	b.network = NewNetwork(shh, symID)
	b.chain = NewChain(contract, b.isAuthority)
	b.chain.Start()
	b.quit = make(chan struct{})
	b.wg.Add(1)
	if b.isAuthority {
		go func() {
			if err := b.autorityLoop(b.chain, b.network); err != nil {
				log.Error("authority exited", "error", err)
			}
			b.wg.Done()
		}()
	} else {
		go func() {
			if err := b.peerLoop(b.chain, b.network); err != nil {
				log.Error("peer exited", "error", err)
			}
			b.wg.Done()
		}()
	}
	return nil
}

func (b *Backend) Stop() {
	b.stateMu.Lock()
	defer b.stateMu.Unlock()
	b.chain.Stop()
	close(b.quit)
	b.wg.Wait()
}

func (b *Backend) peerLoop(chain *Chain, network NetworkClient) error {
	log.Info("running peer loop")
	blocks := make(chan *whisperv5.Message, 20)
	sub, err := network.SubscribeBlock(blocks)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()
	for {
		select {
		case msg := <-blocks:
			var block Block
			if err := rlp.DecodeBytes(msg.Payload, &block); err != nil {
				log.Error("decoding block", "error", err)
			}
			log.Info("received", "block", block)
			chain.AddBlock(&block)
		case <-b.quit:
			return nil
		}
	}
}

func (b *Backend) autorityLoop(chain *Chain, network NetworkClient) error {
	txs := make(chan *whisperv5.Message, 1000)
	sub, err := network.SubscribeTx(txs)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()
	addedBlocks := make(chan *Block, 20)
	chain.SubscribeBlocks(addedBlocks)
	for {
		select {
		case msg := <-txs:
			var tx Transaction
			if err := rlp.DecodeBytes(msg.Payload, &tx); err != nil {
				log.Error("decoding transaction", "error", err, "payload", msg.Payload)
				continue
			}
			log.Info("received", "tx", tx)
			chain.NotifyTx(&tx)
		case block := <-addedBlocks:
			payload, err := rlp.EncodeToBytes(block)
			if err != nil {
				log.Error("encoding block", "error", err)
				continue
			}
			log.Info("signed", "block", block)
			if err := network.BroadcastBlock(payload); err != nil {
				log.Error("broadcasting block", "error", err)
			}
		case <-b.quit:
			return nil
		}
	}
}
