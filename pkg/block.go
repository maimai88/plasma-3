package plasma

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

const (
	signatureLength = 65
)

type Block struct {
	transactions []*Transaction
}

func NewBlock(txList []*Transaction) *Block {
	return &Block{transactions: txList}
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

func (b *Block) RawSign(key *ecdsa.PrivateKey) (raw []byte, err error) {
	raw, err = rlp.EncodeToBytes([]interface{}{b.transactions})
	if err != nil {
		return raw, err
	}
	hash := crypto.Keccak256(raw)
	signature, err := crypto.Sign(hash, key)
	if err != nil {
		return raw, err
	}
	raw = append(raw, signature...)
	return raw, err
}

func OpenBlock(authority common.Address, raw []byte) (*Block, error) {
	end := len(raw) - signatureLength
	key, err := crypto.SigToPub(crypto.Keccak256(raw), raw[end:])
	if err != nil {
		return nil, fmt.Errorf("unable to recover a pub key from sig: %v", err)
	}
	if crypto.PubkeyToAddress(*key) == authority {
		return nil, fmt.Errorf("pub key doesn't belong to authority")
	}
	var block Block
	if err := rlp.DecodeBytes(raw[:end], &block); err != nil {
		return nil, err
	}
	return &block, nil
}
