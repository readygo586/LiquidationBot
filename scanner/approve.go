package scanner

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

func (s *Scanner) doApprove() (*types.Transaction, error) {
	publicKey := s.PrivateKey.Public()
	vai := s.vai
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	repayer := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := s.c.PendingNonceAt(context.Background(), repayer)
	if err != nil {
		return nil, err
	}

	gasPrice, err := s.c.SuggestGasPrice(context.Background())
	if err != nil {
		gasPrice = big.NewInt(5000000000)
	}

	gasLimit := uint64(3000000)
	chainId, err := s.c.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(s.PrivateKey, chainId)
	auth.Value = big.NewInt(0)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = gasLimit

	tx, err := vai.Approve(auth, s.vaiControllerAddr, big.NewInt(-1))
	if err != nil {
		return nil, err
	}

	return tx, nil
}
