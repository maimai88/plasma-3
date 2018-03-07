package plasma

import (
	"math/big"
)

const (
	signatureLength = 65
)

type Block struct {
	Transactions []*Transaction
}

func NewBlock(txList ...*Transaction) *Block {
	return &Block{Transactions: txList}
}

func (b *Block) Amount(txindex, oindex *big.Int) *big.Int {
	if oindex.Int64() == 0 {
		return b.Transactions[txindex.Int64()].Amount1
	}
	return b.Transactions[txindex.Int64()].Amount2
}

func (b *Block) IsSpent(txindex, oindex *big.Int) bool {
	if oindex.Int64() == 0 {
		return b.Transactions[txindex.Int64()].spent1
	}
	return b.Transactions[txindex.Int64()].spent2
}

func (b *Block) SetSpent(txindex, oindex *big.Int) {
	if oindex.Int64() == 0 {
		b.Transactions[txindex.Int64()].spent1 = true
	} else {
		b.Transactions[txindex.Int64()].spent2 = true
	}

}
