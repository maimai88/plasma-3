package plasma

import (
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type PublicPlasmaAPI struct {
	contract *Plasma
	network  NetworkClient
	chain    *Chain
}

func NewPlasmaApi(backend *Backend) *PublicPlasmaAPI {
	// FIXME figure out how to improve it
	return &PublicPlasmaAPI{
		contract: backend.chain.plasmaContract,
		network:  backend.network,
		chain:    backend.chain,
	}
}

// FIXME check how geth handles selected private key
func (api *PublicPlasmaAPI) Deposit(key *ecdsa.PrivateKey, value *big.Int) (*types.Transaction, error) {
	opts := bind.NewKeyedTransactor(key)
	opts.Value = value
	deposit := NewDeposit(opts.From, value)
	return api.contract.Deposit(opts, deposit.EncodeUnsigned())
}

func (api *PublicPlasmaAPI) UtxoBalance(address common.Address) []UTXO {
	return api.chain.FindUTXOs(address)
}

func (api *PublicPlasmaAPI) SendTransaction(tx *Transaction, key1, key2 *ecdsa.PrivateKey) error {
	if err := tx.Sign(key1, key2); err != nil {
		return err
	}
	payload, err := tx.Encode()
	if err != nil {
		return err
	}
	if !api.chain.ValidateTransaction(tx) {
		return errors.New("tx is not valid")
	}
	return api.network.BroadcastTx(payload)
}

func (api *PublicPlasmaAPI) SendConfirmation(recipient *ecdsa.PublicKey, block, tx int, key1, key2 *ecdsa.PrivateKey) error {
	hash, err := api.chain.ConfirmationHash(block, tx)
	if err != nil {
		return err
	}
	confirmSig, err := crypto.Sign(hash[:], key1)
	if err != nil {
		return err
	}
	if key2 != nil {
		sig2, err := crypto.Sign(hash[:], key2)
		if err != nil {
			return err
		}
		confirmSig = append(confirmSig, sig2...)
	}
	conf := Confirmation{
		Blknum:          block,
		Txindex:         tx,
		ConfirmationSig: confirmSig,
	}
	return api.network.SendConfirmation(recipient, conf)
}
