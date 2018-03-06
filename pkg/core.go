package plasma

import (
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

func NewChain(
	contract *Plasma,
	isAuthority bool,
) *Chain {
	return &Chain{
		plasmaContract: contract,
		txQueue:        make(chan *Transaction, 100),
		blockQueue:     make(chan *Block, 20),
		newBlocks:      make([]chan<- *Block, 0, 1),
		isAuthority:    isAuthority,
		blocks:         []*Block{},
	}
}

// Chain manages transactions state.
type Chain struct {
	plasmaContract *Plasma

	isAuthority bool

	// newBlocks used to advertise blocks that were minted by current node
	subMu     sync.RWMutex
	newBlocks []chan<- *Block

	// queues used for objects received from network
	txQueue    chan *Transaction
	blockQueue chan *Block

	// used to synchronize start and stop
	// start and stop should be executed on the same thread
	wg   sync.WaitGroup
	quit chan struct{}

	stateMu sync.RWMutex
	blocks  []*Block
}

func (c *Chain) Start() {
	c.quit = make(chan struct{})
	if c.isAuthority {
		c.wg.Add(1)
		go func() {
			c.stateLoop()
			c.wg.Done()
		}()
	}
}

func (c *Chain) Stop() {
	close(c.quit)
	c.wg.Wait()
}

func (c *Chain) NotifyTx(tx *Transaction) {
	select {
	case c.txQueue <- tx:
	case <-c.quit:
	}
}

func (c *Chain) SubscribeBlocks(blockCh chan<- *Block) {
	c.subMu.Lock()
	defer c.subMu.Unlock()
	c.newBlocks = append(c.newBlocks, blockCh)
}

func (c *Chain) FindUTXOs(address common.Address) []UTXO {
	c.stateMu.RLock()
	defer c.stateMu.RUnlock()
	rst := []UTXO{}
	for i, block := range c.blocks {
		for j, tx := range block.transactions {
			if !tx.spent1 && tx.Newowner1 == address {
				rst = append(rst, UTXO{
					big.NewInt(int64(i) + 1), big.NewInt(int64(j)),
					big.NewInt(0), tx.Amount1})
			}
			if !tx.spent2 && tx.Newowner2 == address {
				rst = append(rst, UTXO{
					big.NewInt(int64(i) + 1), big.NewInt(int64(j)),
					big.NewInt(1), tx.Amount2})
			}
		}
	}
	return rst
}

func (c *Chain) AddBlock(block *Block) {
	c.stateMu.Lock()
	defer c.stateMu.Unlock()
	c.blocks = append(c.blocks, block)
	for _, tx := range block.transactions {
		if tx.Blknum1 != nil {
			c.blocks[tx.Blknum1.Int64()-1].SetSpent(tx.Txindex1, tx.Oindex1)
		}
		if tx.Blknum2 != nil {
			c.blocks[tx.Blknum2.Int64()-1].SetSpent(tx.Txindex2, tx.Oindex2)
		}
	}
}

func (c *Chain) ValidateTransaction(tx *Transaction) bool {
	c.stateMu.Lock()
	defer c.stateMu.Unlock()
	inputs := big.NewInt(0)
	if tx.Blknum1 != nil {
		if c.blocks[tx.Blknum1.Int64()-1].IsSpent(tx.Txindex1, tx.Oindex1) {
			log.Info("invalid transaction")
			return false
		}
		inputs.Add(c.blocks[tx.Blknum1.Int64()-1].Amount(tx.Txindex1, tx.Oindex1), inputs)
	}
	if tx.Blknum2 != nil {
		if c.blocks[tx.Blknum2.Int64()-1].IsSpent(tx.Txindex2, tx.Oindex2) {
			log.Info("invalid transaction")
			return false
		}
		inputs.Add(c.blocks[tx.Blknum1.Int64()-1].Amount(tx.Txindex2, tx.Oindex2), inputs)
	}
	outputs := big.NewInt(0)
	outputs.Add(tx.Amount1, tx.Amount2)
	outputs.Add(outputs, tx.Fee)
	if inputs.Cmp(outputs) != 0 {
		log.Info("tx invalid: inputs", inputs, "!=", "ouputs", outputs)
		return false
	}
	return true
}

// FIXME only autority needs to run this loop
func (c *Chain) stateLoop() {
	// FIXME concurrent deposits and regular blocks
	pendingTransactions := make([]*Transaction, 0, 100)
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	depositLogs := make(chan *PlasmaDeposit, 10)
	// FIXME wrap whole loop with retry on network errors
	sub, _ := c.plasmaContract.WatchDeposit(nil, depositLogs)
	defer sub.Unsubscribe()
	for {
		select {
		case <-c.quit:
			return
		case tx := <-c.txQueue:
			// this should not be used if node isn't authority
			if c.ValidateTransaction(tx) {
				pendingTransactions = append(pendingTransactions, tx)
			}
		case <-ticker.C:
			if len(pendingTransactions) != 0 {
				block := NewBlock(pendingTransactions...)
				c.AddBlock(block)
				c.subMu.RLock()
				for _, ch := range c.newBlocks {
					select {
					case ch <- block:
					case <-time.After(100 * time.Microsecond):
					}
				}
				c.subMu.RUnlock()
			}
			pendingTransactions = make([]*Transaction, 0, 100)
		case deposit := <-depositLogs:
			block := NewBlock(NewDeposit(deposit.Depositor, deposit.Value))
			c.AddBlock(block)
			c.subMu.RLock()
			for _, ch := range c.newBlocks {
				select {
				case ch <- block:
				case <-time.After(100 * time.Microsecond):
				}
			}
			c.subMu.RUnlock()
		}
	}
}
