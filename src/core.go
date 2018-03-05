package plasma

import (
	"sync"
	"time"
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
	}
}

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
	c.wg.Add(1)
	go func() {
		c.stateLoop()
		c.wg.Done()
	}()
}

func (c *Chain) Stop() {
	close(c.quit)
	c.wg.Wait()
}

func (c *Chain) NotifyTx(tx *Transaction) {
	select {
	case c.txQueue <- tx:
	case c.quit:
	}
}

func (c *Chain) NotifyBlock(block *Block) {
	select {
	case c.blockQueue <- block:
	case c.quit:
	}
}

func (c *Chain) SubscribeBlocks(blockCh chan<- *Block) {
	c.subMu.Lock()
	defer c.subMu.Unlock()
	c.newBlocks = append(c.newBlocks, blockCh)
}

func (c *Chain) addBlock(block *Block) {
	c.stateMu.Lock()
	defer c.stateMu.Unlock()
	c.blocks = append(c.blocks, block)
}

func (c *Chain) stateLoop() {
	pendingTransactions := make([]*Transaction, 0, 100)
	var signerPeriod <-chan time.Time
	// this required better approach
	if c.isAuthority {
		ticker := time.NewTicker(100 * time.Millisecond)
		signerPeriod = ticker.C
		defer ticker.Stop()
	}
	depositLogs := make(chan *PlasmaDeposit, 10)
	// if minter exits with error wrap it with retry
	sub, _ := c.plasmaContract.WatchDeposit(nil, depositLogs)
	defer sub.Unsubscribe()
	for {
		select {
		case <-c.quit:
			return
		case tx := <-c.txQueue:
			// this will leak if node is not an authority
			pendingTransactions = append(pendingTransactions, tx)
		case <-signerPeriod:
			if len(pendingTransactions) != 0 {
				block := NewBlock(pendingTransactions)
				c.addBlock(block)
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
			block := NewBlock([]*Transaction{NewDeposit(deposit.Depositor, deposit.Value)})
			c.addBlock(block)
		case block := <-c.blockQueue:
			c.addBlock(block)
		}
	}
}
