package plasma

import (
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type PublicPlasmaAPI struct {
	contract *PlasmaTransactor
	network  NetworkClient
	chain    *Chain
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
