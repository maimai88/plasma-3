package plasma

import (
	"math/big"

	"github.com/dshulyak/plasma/pkg/merkle"
	"github.com/ethereum/go-ethereum/common"
)

const (
	signatureLength = 65
)

type Block struct {
	Transactions []*Transaction

	tree merkle.FixedMerkle
}

func NewBlock(txList ...*Transaction) *Block {
	leafs := [][]byte{}
	for _, tx := range txList {
		leafs = append(leafs, tx.Bytes())
	}
	tree := merkle.New(16, leafs...)
	return &Block{
		Transactions: txList,
		tree:         tree,
	}
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

func (b *Block) Root() common.Hash {
	return b.tree.Root()
}
