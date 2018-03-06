package plasma

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

type Transaction struct {
	Blknum1   *big.Int
	Txindex1  *big.Int
	Oindex1   *big.Int
	Blknum2   *big.Int
	Txindex2  *big.Int
	Oindex2   *big.Int
	Newowner1 common.Address
	Amount1   *big.Int
	Newowner2 common.Address
	Amount2   *big.Int
	Fee       *big.Int
	Sig1      []byte
	Sig2      []byte

	spent1 bool
	spent2 bool
}

func (tx *Transaction) EncodeUnsigned() []byte {
	rst, _ := rlp.EncodeToBytes([]interface{}{
		tx.Blknum1, tx.Txindex1, tx.Oindex1,
		tx.Blknum2, tx.Txindex2, tx.Oindex2,
		tx.Newowner1, tx.Amount1,
		tx.Newowner2, tx.Amount2,
		tx.Fee,
	})
	return rst
}

func (tx *Transaction) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(tx)
}

func (tx *Transaction) Sign(key1 *ecdsa.PrivateKey, key2 *ecdsa.PrivateKey) (err error) {
	raw := tx.EncodeUnsigned()
	hash := crypto.Keccak256(raw)
	tx.Sig1, err = crypto.Sign(hash, key1)
	if err != nil {
		return err
	}
	if key2 == nil {
		return nil
	}
	tx.Sig2, err = crypto.Sign(hash, key2)
	return err
}

func NewDeposit(depositor common.Address, value *big.Int) *Transaction {
	// no idea why but vyper fails to parse rlp lists with additional null bytes
	return &Transaction{
		Txindex1:  big.NewInt(1),
		Oindex1:   big.NewInt(1),
		Txindex2:  big.NewInt(1),
		Oindex2:   big.NewInt(1),
		Newowner1: depositor,
		Amount1:   value,
	}
}

type UTXO struct {
	Block  *big.Int
	Tx     *big.Int
	OIndex *big.Int
	Amount *big.Int
}
