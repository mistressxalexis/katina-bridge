package test

import (
	"example.com/m/app/config"
	"log"
	"math/big"
	"os"
	"testing"

	nft "example.com/m/app/pkg/model/NFT"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var nfts = "0x3a94a68266fB1c9EAc522077e1C9901684670B0B"

func TestMint(t *testing.T) {
	config.Config()
	cli, err := ethclient.Dial(os.Getenv("URL_SEPOLIA"))
	if err != nil {
		t.Error("1", err)
	}
	leopard, err := nft.NewLeopard(common.HexToAddress(nfts), cli)
	if err != nil {
		t.Error(err)
	}
	key, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	auth, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(11155111))
	_, err = leopard.MintItem(auth, "www.example.com")
	if err != nil {
		t.Log(err)
	}
}
