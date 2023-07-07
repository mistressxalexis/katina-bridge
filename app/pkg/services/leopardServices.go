package services

import (
	"example.com/m/app/config"
	"example.com/m/app/pkg/core"
	nft "example.com/m/app/pkg/model/NFT"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

func LeopardMintNFT(private string, uri string) {
	go func() {
		addr := common.HexToAddress(config.AppConfig.LEO_ADDRESS)
		sentinels, err := nft.NewLeopard(addr, config.SepoliaClient())
		if err != nil {
			log.Fatal(err)
		}
		_, err = sentinels.MintItem(
			core.NewTransactOpts(&config.TransactOptsConfig{
				PrivateKey: private,
				ChainID:    config.AppConfig.SepoliaChainID,
			}),
			uri,
		)
		if err != nil {
			log.Fatal(err)
		}
	}()
}
