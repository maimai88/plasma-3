package plasma

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactionSign(t *testing.T) {
	key1, err := crypto.GenerateKey()
	require.NoError(t, err)
	addr1 := crypto.PubkeyToAddress(key1.PublicKey)
	key2, err := crypto.GenerateKey()
	require.NoError(t, err)
	addr2 := crypto.PubkeyToAddress(key2.PublicKey)
	tx := NewTransaction(UTXO{}, UTXO{}, addr1, addr2, big.NewInt(15), big.NewInt(10), big.NewInt(5))
	require.NoError(t, tx.Sign(key1, key2))
	hash := crypto.Keccak256(tx.EncodeUnsigned())
	pub1, err := crypto.SigToPub(hash, tx.Sig1)
	require.NoError(t, err)
	assert.Equal(t, addr1, crypto.PubkeyToAddress(*pub1))
	pub2, err := crypto.SigToPub(hash, tx.Sig2)
	require.NoError(t, err)
	assert.Equal(t, addr2, crypto.PubkeyToAddress(*pub2))

	encoded, err := tx.Encode()
	require.NoError(t, err)
	var newTx Transaction
	require.NoError(t, rlp.DecodeBytes(encoded, &newTx))
	assert.Equal(t, tx, &newTx)
}

func TestEncodeDecodeDeposit(t *testing.T) {
	deposit := NewDeposit(common.Address{253}, big.NewInt(10))
	encoded, err := deposit.Encode()
	require.NoError(t, err)
	var tx Transaction
	require.NoError(t, rlp.DecodeBytes(encoded, &tx))
	assert.Equal(t, deposit, &tx)
}
