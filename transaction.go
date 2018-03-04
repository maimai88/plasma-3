package plasma

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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
