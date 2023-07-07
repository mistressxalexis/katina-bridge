package config

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var sepoliaCli *ethclient.Client = nil
var arbitrumCli *ethclient.Client = nil

func Config() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	appConig()
}

var AppConfig struct {
	SepoliaChainID    int64
	ARBChainID        int64
	URLSepolia        string
	URLARB            string
	LEO_ADDRESS       string
	SEN_ADDRESS       string
	SEN_BRIDGE_ADDR   string
	LEO_BRIDGE_ADDR   string
	FWallet           string
	SWallet           string
	PrivateKeySepolia string
	PrivateKeyARB     string
	WSSSEPO           string
	WSSARB            string
	DB_URI            string
	PrivateKeyBackend string
}

type TransactOptsConfig struct {
	PrivateKey string `json:"private_key"`
	ChainID    int64  `json:"chain_id"`
}

func appConig() {
	AppConfig.SepoliaChainID = 11155111
	AppConfig.ARBChainID = 421613
	AppConfig.URLSepolia = os.Getenv("URL_SEPOLIA")
	AppConfig.URLARB = os.Getenv("URL_ARB")
	AppConfig.LEO_ADDRESS = os.Getenv("LEO_ADDRESS")
	AppConfig.SEN_ADDRESS = os.Getenv("SEN_ADDRESS")
	AppConfig.SEN_BRIDGE_ADDR = os.Getenv("SEN_BRIDGE_ADDR")
	AppConfig.LEO_BRIDGE_ADDR = os.Getenv("LEO_BRIDGE_ADDR")
	AppConfig.FWallet = os.Getenv("FWALLET")
	AppConfig.SWallet = os.Getenv("SWALLET")
	AppConfig.PrivateKeySepolia = os.Getenv("PRIVATE_KEY")
	AppConfig.PrivateKeyARB = os.Getenv("PRIVATE_KEY")
	AppConfig.WSSSEPO = os.Getenv("WSC_SEPO")
	AppConfig.WSSARB = os.Getenv("WSC_ARB")
	AppConfig.DB_URI = os.Getenv("DB_URI")
	AppConfig.PrivateKeyBackend = os.Getenv("PRIVATE_KEY")
}

func SepoliaConfig() *TransactOptsConfig {
	return &TransactOptsConfig{
		PrivateKey: AppConfig.PrivateKeySepolia,
		ChainID:    AppConfig.SepoliaChainID,
	}
}

func ArbitrumConfig() *TransactOptsConfig {
	return &TransactOptsConfig{
		PrivateKey: AppConfig.PrivateKeyARB,
		ChainID:    AppConfig.ARBChainID,
	}
}

func SepoliaClient() *ethclient.Client {
	if sepoliaCli == nil {
		sepoliaCli, err := ethclient.Dial(AppConfig.WSSSEPO)
		if err != nil {
			log.Fatal("NewCoreSEPO ERROR: ", err)
		}
		return sepoliaCli
	}
	return sepoliaCli
}

func ArbitrumClient() *ethclient.Client {
	if arbitrumCli == nil {
		arbitrumCli, err := ethclient.Dial(AppConfig.WSSARB)
		if err != nil {
			log.Fatal("NewCoreARB ERROR: ", err)
		}
		return arbitrumCli
	}
	return arbitrumCli
}
