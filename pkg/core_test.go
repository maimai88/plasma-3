package plasma

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddBlock(t *testing.T) {
	chain := NewChain(nil, false)
	key1, err := crypto.GenerateKey()
	require.NoError(t, err)
	addr1 := crypto.PubkeyToAddress(key1.PublicKey)

	key2, err := crypto.GenerateKey()
	require.NoError(t, err)
	addr2 := crypto.PubkeyToAddress(key2.PublicKey)

	deposit1 := NewDeposit(addr1, big.NewInt(100))
	deposit2 := NewDeposit(addr2, big.NewInt(200))
	// this is just for test purposes, on main chain deposits
	// will be in the separate blocks
	chain.AddBlock(NewBlock(deposit1, deposit2))

	utxos := chain.FindUTXOs(addr1)
	require.Len(t, utxos, 1)
	assert.Equal(t,
		UTXO{big.NewInt(1), big.NewInt(0), big.NewInt(0), big.NewInt(100)}, utxos[0])

	tx := NewTransaction(utxos[0], UTXO{}, addr1, addr2, big.NewInt(45), big.NewInt(50), big.NewInt(5))
	require.True(t, chain.ValidateTransaction(tx))
	chain.AddBlock(NewBlock(tx))
	utxos = chain.FindUTXOs(addr2)
	require.Len(t, utxos, 2)
	assert.Equal(t, utxos[0], UTXO{big.NewInt(1), big.NewInt(1), big.NewInt(0), big.NewInt(200)})
	assert.Equal(t, utxos[1], UTXO{big.NewInt(2), big.NewInt(0), big.NewInt(1), big.NewInt(50)})

	utxos = chain.FindUTXOs(addr1)
	require.Len(t, utxos, 1)
	assert.Equal(t, utxos[0], UTXO{big.NewInt(2), big.NewInt(0), big.NewInt(0), big.NewInt(45)})
}
