package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

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
}

func appConig() {
	AppConfig.SepoliaChainID = 11155111
	AppConfig.ARBChainID = 421613
	AppConfig.URLSepolia = os.Getenv("URL_SEPOLIA")
	AppConfig.URLARB = os.Getenv("URL_ARB")
	AppConfig.T55NFT = os.Getenv("T55_ADDRESS")
	AppConfig.T54NFT = os.Getenv("T54_ADDRESS")
	AppConfig.T54Bridge = os.Getenv("T54BRIDGE_ADDR")
	AppConfig.T55Bridge = os.Getenv("T55BRIDGE_ADDR")
	AppConfig.FWallet = os.Getenv("FWALLET")
	AppConfig.SWallet = os.Getenv("SWALLET")
	AppConfig.PrivateKeySepolia = os.Getenv("PRIVATE_KEY_SEPOLIA")
	AppConfig.PrivateKeyARB = os.Getenv("PRIVATE_KEY_ARB")
	AppConfig.WSSSEPO = os.Getenv("WSC_SEPO")
	AppConfig.WSSARB = os.Getenv("WSC_ARB")
}
