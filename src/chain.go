package plasma

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	signatureLength = 65
)

func NewChain(contract *Plasma, txQueue <-chan Transaction, blockQueue <-chan Block) *Chain {
	return &Chain{
		plasmaContract: contract,
		txQueue:        txQueue,
		blockQueue:     blockQueue,
	}
}

type Chain struct {
	plasmaContract *Plasma
	txQueue        <-chan Transaction
	blockQueue     <-chan Block
	isAuthority    bool

	authority common.Address

	wg   sync.WaitGroup
	quit chan struct{}

	stateMu sync.RWMutex
	blocks  []*Block
}

func (c *Chain) Start() error {
	// move this verification elsewhere, it doesn't make sense here
	authority, err := c.plasmaContract.Authority(nil)
	if err != nil {
		return err
	}
	c.wg.Add(1)
	go c.minter()
	c.authority = authority
	c.quit = make(chan struct{})
}

func (c *Chain) Stop() {
	close(c.quit)
	c.wg.Wait()
}

// will be used in 3 different cases
// 1. deposit event received from a contract
// 2. block mined by authority every 100ms
// 3. signed block received from a network
func (c *Chain) addBlock(block *Block) {
	c.stateMu.Lock()
	defer c.stateMu.Unlock()
	c.blocks = append(c.blocks, block)
}

func (c *Chain) minter() {
	pendingTransactions := []Transaction{}
	var minterPeriod <-chan time.Time
	if c.isAuthority {
		ticker := time.NewTicker(100 * time.Millisecond)
		minterPeriod = ticker.C
		defer ticker.Stop()
	}
	events := make(chan *PlasmaDeposit, 10)
	sub, _ := c.plasmaContract.WatchDeposit(nil, events)
	defer sub.Unsubscribe()
	defer func() {
		c.wg.Done()
	}()
	for {
		select {
		case <-c.quit:
			return
		case tx := <-c.txQueue:
			pendingTransactions = append(pendingTransactions, tx)
		case <-minterPeriod:
			if len(pendingTransactions) != 0 {
				block := NewBlock(pendingTransactions)
				c.addBlock(block)
			}
			// notify network
		case ev := <-events:
			block := NewBlock([]Transaction{NewDeposit(ev.Depositor, ev.Value)})
			c.addBlock(block)
		case block := <-c.blockQueue:
			c.addBlock(&block)
		}
	}
}
