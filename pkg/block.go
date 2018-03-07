package plasma

import (
	"math/big"
)

const (
	signatureLength = 65
)

type Block struct {
	transactions []*Transaction
}

func NewBlock(txList ...*Transaction) *Block {
	return &Block{transactions: txList}
}

func (b *Block) Amount(txindex, oindex *big.Int) *big.Int {
	if oindex.Int64() == 0 {
		return b.transactions[txindex.Int64()].Amount1
	}
	return b.transactions[txindex.Int64()].Amount2
}

func (b *Block) IsSpent(txindex, oindex *big.Int) bool {
	if oindex.Int64() == 0 {
		return b.transactions[txindex.Int64()].spent1
	}
	return b.transactions[txindex.Int64()].spent2
}

func (b *Block) SetSpent(txindex, oindex *big.Int) {
	if oindex.Int64() == 0 {
		b.transactions[txindex.Int64()].spent1 = true
	} else {
		b.transactions[txindex.Int64()].spent2 = true
	}

}
