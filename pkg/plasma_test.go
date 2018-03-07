package plasma

import (
	"context"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/dshulyak/plasma/pkg/merkle"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

var RLPDECODERSENDER = [20]byte{57, 186, 8, 60, 48, 252, 229, 152, 131, 119, 95, 199, 41, 187, 225, 249, 222, 77, 238, 17}

const RLPDECODERHEX = "0xf9035b808506fc23ac0083045f788080b903486103305660006109ac5260006109cc527f0100000000000000000000000000000000000000000000000000000000000000600035046109ec526000610a0c5260006109005260c06109ec51101515585760f86109ec51101561006e5760bf6109ec510336141558576001610a0c52610098565b60013560f76109ec51036020035260005160f66109ec510301361415585760f66109ec5103610a0c525b61022060016064818352015b36610a0c511015156100b557610291565b7f0100000000000000000000000000000000000000000000000000000000000000610a0c5135046109ec526109cc5160206109ac51026040015260016109ac51016109ac5260806109ec51101561013b5760016109cc5161044001526001610a0c516109cc5161046001376001610a0c5101610a0c5260216109cc51016109cc52610281565b60b86109ec5110156101d15760806109ec51036109cc51610440015260806109ec51036001610a0c51016109cc51610460013760816109ec5114156101ac5760807f01000000000000000000000000000000000000000000000000000000000000006001610a0c5101350410151558575b607f6109ec5103610a0c5101610a0c5260606109ec51036109cc51016109cc52610280565b60c06109ec51101561027d576001610a0c51013560b76109ec510360200352600051610a2c526038610a2c5110157f01000000000000000000000000000000000000000000000000000000000000006001610a0c5101350402155857610a2c516109cc516104400152610a2c5160b66109ec5103610a0c51016109cc516104600137610a2c5160b66109ec5103610a0c510101610a0c526020610a2c51016109cc51016109cc5261027f565bfe5b5b5b81516001018083528114156100a4575b5050601f6109ac511115155857602060206109ac5102016109005260206109005103610a0c5261022060016064818352015b6000610a0c5112156102d45761030a565b61090051610a0c516040015101610a0c51610900516104400301526020610a0c5103610a0c5281516001018083528114156102c3575b50506109cc516109005101610420526109cc5161090051016109005161044003f35b61000461033003610004600039610004610330036000f31b2d4f"

func TestSubmitBlock(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	addr := crypto.PubkeyToAddress(key.PublicKey)
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{addr: {Balance: big.NewInt(10000000000000)}})
	opts := bind.NewKeyedTransactor(key)

	_, _, contract, err := DeployPlasma(opts, backend)
	require.NoError(t, err)
	backend.Commit()

	authority, err := contract.Authority(nil)
	require.NoError(t, err)
	blocknum, err := contract.Last_parent_block(nil)
	require.NoError(t, err)
	assert.Equal(t, addr, authority)
	assert.Equal(t, int64(1), blocknum.Int64())

	for i := 0; i < 6; i++ {
		backend.Commit()
	}

	hash := make([]byte, 32)
	rand.Read(hash)
	root := [32]byte{}
	copy(root[:], hash[:])
	opts.GasLimit = 81689
	_, err = contract.SubmitBlock(opts, root)
	require.NoError(t, err)
	backend.Commit()
	blocknum, err = contract.Last_parent_block(nil)
	assert.NoError(t, err)
	assert.Equal(t, int64(8), blocknum.Int64())

	childnum, err := contract.Last_child_block(nil)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), childnum.Int64())

	childChainRoot, err := contract.Child_chain__root(nil, big.NewInt(1))
	assert.NoError(t, err)
	assert.Equal(t, root, childChainRoot)
}

func TestDeposit(t *testing.T) {
	rlpTx := new(types.Transaction)
	txbytes, err := hexutil.Decode(RLPDECODERHEX)
	require.NoError(t, err)
	require.NoError(t, rlp.DecodeBytes(txbytes, rlpTx))

	key2, err := crypto.GenerateKey()
	require.NoError(t, err)
	addr2 := crypto.PubkeyToAddress(key2.PublicKey)
	assert.NotEqual(t, addr2, common.Address{})

	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	addr := crypto.PubkeyToAddress(key.PublicKey)
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		addr:             {Balance: big.NewInt(10000000000000)},
		addr2:            {Balance: big.NewInt(100000000000000)},
		RLPDECODERSENDER: {Balance: big.NewInt(85975200000000000)},
	})
	opts := bind.NewKeyedTransactor(key)
	_, _, contract, err := DeployPlasma(opts, backend)
	require.NoError(t, err)
	backend.Commit()

	require.NoError(t, backend.SendRawTransaction(context.TODO(), rlpTx))
	backend.Commit()

	value := big.NewInt(10000000000000)
	opts = bind.NewKeyedTransactor(key2)
	opts.Value = value
	opts.GasLimit = 96273
	tx := NewDeposit(addr2, value)
	encoded := tx.EncodeUnsigned()
	_, err = contract.Deposit(opts, encoded)
	assert.NoError(t, err)
	backend.Commit()
	blocknum, err := contract.Last_parent_block(nil)
	assert.NoError(t, err)
	assert.Equal(t, int64(3), blocknum.Int64())

	childnum, err := contract.Last_child_block(nil)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), childnum.Int64())

	encoded = append(encoded, tx.Sig1...)
	encoded = append(encoded, tx.Sig2...)
	tree := merkle.New(16, encoded)
	// test that root hash is expected
	root, err := contract.Child_chain__root(nil, big.NewInt(1))
	assert.NoError(t, err)
	assert.Equal(t, tree.Root(), common.Hash(root))
}
