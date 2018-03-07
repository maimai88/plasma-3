package plasma

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/whisper/shhclient"
	"github.com/stretchr/testify/suite"
)

// e2e test for a plasma chain
// deploy a plasma contract on a chain
// make deposits from two users
// send transaction from user1 to user2 + do several rounds of such transactions
// send a confirmation over whisper
// withdraw funds + successful finalize
// withdraw funds and challenge exit
func TestSimulation(t *testing.T) {
	suite.Run(t, new(SimulationTestSuite))
}

type account struct {
	key     *ecdsa.PrivateKey
	address common.Address
}

type SimulationTestSuite struct {
	suite.Suite

	backend        *backends.SimulatedBackend
	accounts       []account
	plasmaAddress  common.Address
	nodes          []*node.Node
	plasmaBackends []*Backend
}

func (s *SimulationTestSuite) SetupTest() {
	log.Root().SetHandler(log.StderrHandler)
	s.accounts = make([]account, 3)
	s.nodes = make([]*node.Node, 3)
	s.plasmaBackends = make([]*Backend, 3)
	for i := range s.accounts {
		port := 8870
		key, err := crypto.GenerateKey()
		s.Require().NoError(err)
		s.accounts[i] = account{key, crypto.PubkeyToAddress(key.PublicKey)}
		cfg := node.Config{
			Name: fmt.Sprintf("node-%d", i),
			P2P: p2p.Config{
				NoDiscovery: true,
				MaxPeers:    20,
				PrivateKey:  key,
				ListenAddr:  fmt.Sprintf(":%d", port+i),
			},
		}
		node, err := NewNode(&cfg)
		s.Require().NoError(err)
		s.nodes[i] = node
	}

	s.backend = backends.NewSimulatedBackend(core.GenesisAlloc{
		s.accounts[0].address: {Balance: big.NewInt(10000000000000)},
		s.accounts[1].address: {Balance: big.NewInt(10000000000000)},
		s.accounts[2].address: {Balance: big.NewInt(10000000000000)},
		RLPDECODERSENDER:      {Balance: big.NewInt(85975200000000000)},
	})

	rlpTx := new(types.Transaction)
	txbytes, err := hexutil.Decode(RLPDECODERHEX)
	s.Require().NoError(err)
	s.Require().NoError(rlp.DecodeBytes(txbytes, rlpTx))

	opts := bind.NewKeyedTransactor(s.accounts[0].key)
	address, _, _, err := DeployPlasma(opts, s.backend)
	s.Require().NoError(err)
	s.plasmaAddress = address

	s.plasmaBackends[0] = NewBackend(s.plasmaAddress, true)
	s.plasmaBackends[1] = NewBackend(s.plasmaAddress, false)
	s.plasmaBackends[2] = NewBackend(s.plasmaAddress, false)

	s.Require().NoError(s.backend.SendRawTransaction(context.TODO(), rlpTx))
	// commit plasma and rlp decoder
	s.backend.Commit()
	for i := range s.nodes {
		s.Require().NoError(s.nodes[i].Start())
	}
}

func (s *SimulationTestSuite) TestWorkflow() {
	rpc, err := s.nodes[0].Attach()
	s.Require().NoError(err)
	shh := shhclient.NewClient(rpc)
	s.Require().NoError(s.plasmaBackends[0].Start(shh, s.backend))

	rpc2, err := s.nodes[1].Attach()
	s.Require().NoError(err)
	shh2 := shhclient.NewClient(rpc2)
	s.Require().NoError(s.plasmaBackends[1].Start(shh2, s.backend))

	s.nodes[1].Server().AddPeer(s.nodes[0].Server().Self())

	api := NewPlasmaApi(s.plasmaBackends[0])
	api.Deposit(s.accounts[0].key, big.NewInt(1000))
	s.backend.Commit()
	time.Sleep(time.Second)
	api2 := NewPlasmaApi(s.plasmaBackends[1])
	s.Equal(api.UtxoBalance(s.accounts[0].address), api2.UtxoBalance(s.accounts[0].address))
	api2.Deposit(s.accounts[1].key, big.NewInt(2000))
	s.backend.Commit()
	time.Sleep(time.Second)
	s.Equal(api.UtxoBalance(s.accounts[1].address), api2.UtxoBalance(s.accounts[1].address))
	utxos := api2.UtxoBalance(s.accounts[1].address)
	s.Require().Len(utxos, 1)
	tx := NewTransaction(utxos[0], EmptyUTXO(), s.accounts[0].address, common.Address{}, utxos[0].Amount, big.NewInt(0), big.NewInt(0))
	s.Require().NoError(api2.SendTransaction(tx, s.accounts[1].key, nil))
	time.Sleep(time.Second)
	s.Len(api.UtxoBalance(s.accounts[0].address), 2)
	s.Len(api2.UtxoBalance(s.accounts[1].address), 0)
}
