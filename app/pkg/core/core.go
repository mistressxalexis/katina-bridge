package core

import (
	"crypto/ecdsa"
	"example.com/m/app/config"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	SepoliaChainID    int64 = 11155111
	ARBChainID        int64 = 421613
	URLSepolia        string
	URLARB            string
	T55NFT            string
	T54NFT            string
	T54Bridge         string
	T55Bridge         string
	FWallet           string
	SWallet           string
	PrivateKeySepolia string
	PrivateKeyARB     string
	WSSSEPO           string
	WSSARB            string
)

func constructor() {
	URLSepolia = os.Getenv("URL_SEPOLIA")
	URLARB = os.Getenv("URL_ARB")
	T55NFT = os.Getenv("T55_ADDRESS")
	T54NFT = os.Getenv("T54_ADDRESS")
	T54Bridge = os.Getenv("T54BRIDGE_ADDR")
	T55Bridge = os.Getenv("T55BRIDGE_ADDR")
	FWallet = os.Getenv("FWALLET")
	SWallet = os.Getenv("SWALLET")
	PrivateKeySepolia = os.Getenv("PRIVATE_KEY_SEPOLIA")
	PrivateKeyARB = os.Getenv("PRIVATE_KEY_ARB")
	WSSSEPO = os.Getenv("WSC_SEPO")
	WSSARB = os.Getenv("WSC_ARB")
}

type Core struct {
	Cli        *ethclient.Client
	Wallet     common.Address
	BridgeAddr common.Address
	NFTs       common.Address
	PrivateKey *ecdsa.PrivateKey
	ChainID    *big.Int
}

func NewCoreT55() *Core {
	constructor()
	cli, err := ethclient.Dial(WSSSEPO)
	if err != nil {
		log.Fatal("NewCoreSepolia ERROR: ", err)
	}

	key, err := crypto.HexToECDSA(PrivateKeySepolia)
	if err != nil {
		log.Fatal(err)
	}
	return &Core{
		Cli:        cli,
		Wallet:     common.HexToAddress(FWallet),
		NFTs:       common.HexToAddress(T55NFT),
		BridgeAddr: common.HexToAddress(T55Bridge),
		PrivateKey: key,
		ChainID:    big.NewInt(SepoliaChainID),
	}
}

func NewCoreT54() *Core {
	constructor()
	cli, err := ethclient.Dial(WSSARB)
	if err != nil {
		log.Fatal("NewCoreARB ERROR: ", err)
	}

	key, err := crypto.HexToECDSA(PrivateKeySepolia)
	if err != nil {
		log.Fatal(err)
	}
	return &Core{
		Cli:        cli,
		Wallet:     common.HexToAddress(SWallet),
		NFTs:       common.HexToAddress(T54NFT),
		BridgeAddr: common.HexToAddress(T54Bridge),
		PrivateKey: key,
		ChainID:    big.NewInt(ARBChainID),
	}
}

func (c *Core) T54NewKeyedTransactor() *bind.TransactOpts {
	constructor()
	key, err := crypto.HexToECDSA(PrivateKeySepolia)
	if err != nil {
		log.Fatal(err)
	}
	auth, errAuth := bind.NewKeyedTransactorWithChainID(key, big.NewInt(ARBChainID))
	if errAuth != nil {
		log.Fatal(errAuth)
	}
	return auth
}

func (c *Core) T55NewKeyedTransactor() *bind.TransactOpts {
	constructor()
	key, err := crypto.HexToECDSA(PrivateKeySepolia)
	if err != nil {
		log.Fatal(err)
	}
	auth, errAuth := bind.NewKeyedTransactorWithChainID(key, big.NewInt(SepoliaChainID))
	if errAuth != nil {
		log.Fatal(errAuth)
	}
	return auth
}

func NewTransactOpts(config *config.TransactOptsConfig) *bind.TransactOpts {
	key, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(config.ChainID))
	if err != nil {
		log.Fatal(err)
	}
	return auth
}
