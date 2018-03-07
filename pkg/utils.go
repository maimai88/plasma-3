package plasma

import (
	"fmt"

	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/whisper/whisperv5"
)

func NewNode(config *node.Config) (*node.Node, error) {
	stack, err := node.New(config)
	if err != nil {
		return nil, err
	}
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return whisperv5.New(nil), nil
	}); err != nil {
		return nil, fmt.Errorf("can't register whisper: %v", err)
	}
	return stack, nil
}
