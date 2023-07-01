package services

import (
	"example.com/m/app/pkg/core"
	model2 "example.com/m/app/pkg/model"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func MintT54NFTs(address string, url string) error {
	_to := common.HexToAddress(address)

	cores := core.NewCoreT54()
	t55cor := core.NewCoreT55()

	leo, err := model2.NewT54NFTs(cores.NFTs, cores.Cli)
	if err != nil {
		return err
	}
	t55, err := model2.NewT55NFTs(t55cor.NFTs, t55cor.Cli)
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

	token, err := leo.GetTokenID(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(token);

	_, err = t55.Sync(t55cor.T55NewKeyedTransactor(), token)
	if err != nil {
		return err
	}
	return nil
}

func DepositT54(pkey string, tokenID int64) error {
	// owner := common.HexToAddress(address)
	key, err := crypto.HexToECDSA(pkey)
	if err != nil {
		log.Fatal(err)
	}
	cors := core.NewCoreT54()

	bridge, err := model2.NewT54Bridge(cors.BridgeAddr, cors.Cli)
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

func ClaimT54NFTs(addr string, tokenID int64) error {
	cors := core.NewCoreT54()
	bridgeInstance, err := model2.NewT54Bridge(cors.BridgeAddr, cors.Cli)
	if err != nil {
		return err
	}
	_, err = bridgeInstance.ClaimNFTs(cors.T54NewKeyedTransactor(), common.HexToAddress(addr), big.NewInt(tokenID))
	if err != nil {
		return nil
	}
	return nil
}
