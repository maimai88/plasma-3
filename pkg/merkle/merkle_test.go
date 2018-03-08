package merkle

import (
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestMerkleRoot(t *testing.T) {
	tree := New(3)
	assert.Equal(t, tree.Root(),
		crypto.Keccak256Hash(
			tree.leafes[tree.depth-1][0][:],
			tree.leafes[tree.depth-1][1][:]))
}

func TestMerkleProof(t *testing.T) {
	rand.Seed(time.Now().Unix())
	data := make([]byte, 20)
	rand.Read(data)
	tree := New(16, data)
	rst := crypto.Keccak256Hash(data)
	for _, sibling := range tree.GetProof(data) {
		rst = crypto.Keccak256Hash(rst[:], sibling[:])
	}
	assert.Equal(t, tree.Root(), rst)
}

func TestMerkleVerifyProof(t *testing.T) {
	rand.Seed(time.Now().Unix())

	data1 := make([]byte, 20)
	rand.Read(data1)
	data2 := make([]byte, 20)
	rand.Read(data2)
	tree := New(3, data1, data2)
	proof := []common.Hash{
		crypto.Keccak256Hash(data1),
		tree.leafes[1][1],
		tree.leafes[2][1],
	}
	assert.True(t, tree.Verify(data2, proof))
}
