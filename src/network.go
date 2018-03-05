package plasma

import (
	"github.com/ethereum/go-ethereum/whisper/shhclient"
)

// NetworkClient is wrapper around shh
type NetworkClient struct {
	*shhclient.Client
}

// NotifyRecipient sends a confirmation message to a recipient
// message must be encrypted with public key of recipient
func (c *NetworkClient) NotifyRecipient() {

}

// BroadcastBlock will be used by authority to broadcast newly
// added blocks. Message should be signed by authority private key
// and encrypted by symmetric key
func (c *NetworkClient) BroadcastBlock() {
}

// BroadcastTx will be used by any peer to broadcast any transaction
// can be encrypted by authority public key or sym key
func (c *NetworkClient) BroadcastTx() {
}
