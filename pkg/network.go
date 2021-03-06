package plasma

import (
	"context"
	"crypto/ecdsa"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/whisper/shhclient"
	"github.com/ethereum/go-ethereum/whisper/whisperv5"
)

const (
	blockTopic = "blocks"
	txTopic    = "tx"
)

func NewNetwork(client *shhclient.Client, plasmaKeyID string) NetworkClient {
	return NetworkClient{client: client, plasmaSymKeyID: plasmaKeyID}
}

// NetworkClient is wrapper around shh
type NetworkClient struct {
	client *shhclient.Client

	plasmaSymKeyID string
}

// SendConfirmation sends a confirmation message to a recipient.
// message must be encrypted with public key of recipient
func (c *NetworkClient) SendConfirmation(key *ecdsa.PublicKey, conf Confirmation) error {
	payload, err := rlp.EncodeToBytes(conf)
	if err != nil {
		return err
	}
	return c.client.Post(context.TODO(), whisperv5.NewMessage{
		PublicKey: crypto.FromECDSAPub(key),
		Payload:   payload,
		TTL:       20,
		PowTarget: 10,
		PowTime:   20,
	})
}

// BroadcastBlock will be used by authority to broadcast newly
// added blocks.
func (c *NetworkClient) BroadcastBlock(payload []byte) error {
	return c.client.Post(context.TODO(), whisperv5.NewMessage{
		SymKeyID:  c.plasmaSymKeyID,
		Payload:   payload,
		TTL:       20,
		PowTarget: 10,
		PowTime:   20,
		Topic:     whisperv5.BytesToTopic([]byte(blockTopic)),
	})
}

// BroadcastTx will be used by any peer to broadcast any transaction
// can be encrypted by authority public key or sym key
func (c *NetworkClient) BroadcastTx(payload []byte) error {
	return c.client.Post(context.TODO(), whisperv5.NewMessage{
		SymKeyID:  c.plasmaSymKeyID,
		Payload:   payload,
		TTL:       20,
		PowTarget: 10,
		PowTime:   20,
		Topic:     whisperv5.BytesToTopic([]byte(txTopic)),
	})
}

func (c *NetworkClient) SubscribeBlock(ch chan<- *whisperv5.Message) (ethereum.Subscription, error) {
	return c.client.SubscribeMessages(context.TODO(), whisperv5.Criteria{
		SymKeyID: c.plasmaSymKeyID,
		Topics:   []whisperv5.TopicType{whisperv5.BytesToTopic([]byte(blockTopic))},
	}, ch)
}

func (c *NetworkClient) SubscribeTx(ch chan<- *whisperv5.Message) (ethereum.Subscription, error) {
	return c.client.SubscribeMessages(context.TODO(), whisperv5.Criteria{
		SymKeyID: c.plasmaSymKeyID,
		Topics:   []whisperv5.TopicType{whisperv5.BytesToTopic([]byte(txTopic))},
	}, ch)
}

type Confirmation struct {
	Blknum          int
	Txindex         int
	ConfirmationSig []byte
}
