package merkle

import (
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type FixedMerkle struct {
	depth int

	leafes [][]common.Hash
}

func New(depth int, leafes ...[]byte) FixedMerkle {
	tree := FixedMerkle{depth: depth}
	tree.init(leafes...)
	return tree
}

func (tree *FixedMerkle) init(leafes ...[]byte) {
	tree.leafes = make([][]common.Hash, tree.depth+1)
	count := int64(math.Pow(2, float64(tree.depth)))
	tree.leafes[0] = make([]common.Hash, count)
	for i, l := range leafes {
		tree.leafes[0][i] = crypto.Keccak256Hash(l)
	}
	for level := 1; level <= tree.depth; level++ {
		tree.leafes[level] = make([]common.Hash, len(tree.leafes[level-1])/2)
		elem := 0
		for i := 0; i < len(tree.leafes[level-1]); i += 2 {
			tree.leafes[level][elem] = crypto.Keccak256Hash(tree.leafes[level-1][i][:], tree.leafes[level-1][i+1][:])
			elem++
		}
	}
}

func (tree *FixedMerkle) Root() common.Hash {
	return tree.leafes[tree.depth][0]
}

func (tree *FixedMerkle) GetProof(leaf []byte) []common.Hash {
	proof := make([]common.Hash, 0, tree.depth-1)
	hash := crypto.Keccak256Hash(leaf)
	var index int
	for i, l := range tree.leafes[0] {
		if l == hash {
			index = i
			break
		}
	}
	for i := 0; i < tree.depth; i++ {
		var sibling int
		if index%2 == 0 {
			sibling = index + 1
		} else {
			sibling = index - 1
		}
		index /= 2
		proof = append(proof, tree.leafes[i][sibling])
	}
	return proof
}

func (tree *FixedMerkle) Verify(leaf []byte, proof []common.Hash) bool {
	hash := crypto.Keccak256Hash(leaf)
	var index int
	for i, l := range tree.leafes[0] {
		if l == hash {
			index = i
			break
		}
	}
	for _, sibling := range proof {
		if index%2 == 0 {
			hash = crypto.Keccak256Hash(hash[:], sibling[:])
		} else {
			hash = crypto.Keccak256Hash(sibling[:], hash[:])
		}
		index /= 2
	}
	return hash == tree.Root()
}
