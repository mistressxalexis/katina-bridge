package services

import (
	"example.com/m/app/pkg/core"
	"example.com/m/app/pkg/model"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func MintT55NFTs(address string, url string) error {
	_to := common.HexToAddress(address)

	cores := core.NewCoreT55()

	leo, err := model.NewT55NFTs(cores.NFTs, cores.Cli)
	if err != nil {
		return err
	}
	auth, errAuth := bind.NewKeyedTransactorWithChainID(cores.PrivateKey, cores.ChainID)
	if errAuth != nil {
		return err
	}
	_, err = leo.Mint(auth, _to, url)
	if err != nil {
		return err
	}
	return nil
}

func DepositT55(pkey string, tokenID int64) error {
	// owner := common.HexToAddress(address)
	key, err := crypto.HexToECDSA(pkey)
	if err != nil {
		log.Fatal(err)
	}
	cors := core.NewCoreT55()

	bridge, err := model.NewT55Bridge(cors.BridgeAddr, cors.Cli)
	if err != nil {
		return err
	}
	auth, errAuth := bind.NewKeyedTransactorWithChainID(key, cors.ChainID)
	if errAuth != nil {
		return errAuth
	}
	_, err = bridge.Deposit(auth, big.NewInt(tokenID))
	return err
}

func ClaimT55NFTs(addr string, tokenID int64) error {
	cors := core.NewCoreT55()
	bridgeInstance, err := model.NewT55Bridge(cors.BridgeAddr, cors.Cli)
	if err != nil {
		return err
	}
	_, err = bridgeInstance.ClaimNFTs(cors.T55NewKeyedTransactor(), common.HexToAddress(addr), big.NewInt(tokenID))
	if err != nil {
		return nil
	}
	return nil
}
