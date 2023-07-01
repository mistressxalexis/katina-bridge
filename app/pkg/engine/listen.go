package engine

import (
	"example.com/m/app/pkg/core"
	model2 "example.com/m/app/pkg/model"
	"log"
)

func ListenT55Bridge() {
	ct55 := core.NewCoreT55()
	ct54 := core.NewCoreT54()
	logs := make(chan *model2.T55BridgeDepositor)
	t55br, _ := model2.NewT55Bridge(ct55.BridgeAddr, ct55.Cli)
	t54br, err := model2.NewT55Bridge(ct54.BridgeAddr, ct54.Cli)
	// t54nfts, _ := model.NewT54NFTs(ct54.NFTs, ct54.Cli)
	if err != nil {
		log.Fatal("Error from Listen BridgeSepo", err)
	}
	sub, err := t55br.WatchDepositor(nil, logs)
	if err != nil {
		log.Fatal("WatchDepositor ERROR: ", err)
	}
	// addr, _ := t54br.GetBridgeAddr(nil)
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			log.Println("Received event from Bridge Sepolia: ", vLog.TokenID)
			log.Println("url: ", vLog.Url)
			log.Println("Token ID: ", vLog.TokenID)
			log.Println("From: ", vLog.From)
			log.Println("To: ", vLog.To)

			_, err := t54br.SyncNFTs(ct54.T54NewKeyedTransactor(), vLog.From, vLog.TokenID, vLog.Url)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("NFT locked successfully")
		}
	}
}

func ListenT54Bridge() {
	ct55 := core.NewCoreT55()
	ct54 := core.NewCoreT54()
	logs := make(chan *model2.T54BridgeDepositor)
	t54br, err := model2.NewT54Bridge(ct54.BridgeAddr, ct54.Cli)
	t55br, _ := model2.NewT54Bridge(ct55.BridgeAddr, ct55.Cli)
	// t55nfts, _ := model.NewT55NFTs(ct55.NFTs, ct55.Cli)
	if err != nil {
		log.Fatal("Error from Listen BridgeSepo", err)
	}
	sub, err := t54br.WatchDepositor(nil, logs)
	if err != nil {
		log.Fatal("WatchDepositor ERROR: ", err)
	}
	// addr, _ := t55br.GetBridgeAddr(nil)
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			log.Println("Received event from Bridge ARB: ", vLog.TokenID)
			log.Println("url: ", vLog.Url)
			log.Println("Token ID: ", vLog.TokenID)
			log.Println("From: ", vLog.From)
			log.Println("To: ", vLog.To)

			_, err := t55br.SyncNFTs(ct55.T55NewKeyedTransactor(), vLog.From, vLog.TokenID, vLog.Url)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("NFT locked successfully")

		}
	}
}

func ListenT55MintNFTs() {
	ct55 := core.NewCoreT55()
	ct54 := core.NewCoreT54()

	logs := make(chan *model2.T55NFTsMintNFTs)
	t55instance, err := model2.NewT55NFTs(ct55.NFTs, ct55.Cli)
	if err != nil {
		log.Fatal(err)
	}
	t54instance, err := model2.NewT54NFTs(ct54.NFTs, ct54.Cli)
	if err != nil {
		log.Fatal(err)
	}

	sub, _ := t55instance.WatchMintNFTs(nil, logs)
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vlog := <-logs:
			log.Println("Sync from t55 nft url: ", vlog.Url)
			log.Println("Sync from t55 nft tokenID: ", vlog.TokenID)
			log.Println("Sync from t55 nft played: ", vlog.Player)

			_, err := t54instance.Sync(ct54.T54NewKeyedTransactor(), vlog.TokenID)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("NFTs sync from t55 to t54 successfully")
		}
	}
}

func ListenT54MintNFTs() {
	ct55 := core.NewCoreT55()
	ct54 := core.NewCoreT54()

	logs := make(chan *model2.T54NFTsMintNFTs)
	t55instance, err := model2.NewT55NFTs(ct55.NFTs, ct55.Cli)
	if err != nil {
		log.Fatal(err)
	}
	t54instance, err := model2.NewT54NFTs(ct54.NFTs, ct54.Cli)
	if err != nil {
		log.Fatal(err)
	}

	sub, _ := t54instance.WatchMintNFTs(nil, logs)
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vlog := <-logs:
			log.Println("Sync from t54 nft url: ", vlog.Url)
			log.Println("Sync from t54 nft tokenID: ", vlog.TokenID)
			log.Println("Sync from t54 nft played: ", vlog.Player)

			_, err := t55instance.Sync(ct55.T55NewKeyedTransactor(), vlog.TokenID)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("NFTs sync from t54 to t55 successfully")
		}
	}
}

func ListenClaimT55NFTs() {
	ct55 := core.NewCoreT55()
	// ct54 := core.NewCoreT54()

	logs := make(chan *model2.T55BridgeClaimNFTs)
	t55instance, err := model2.NewT55Bridge(ct55.BridgeAddr, ct55.Cli)
	if err != nil {
		log.Fatal(err)
	}
	t55nfts, err := model2.NewT55NFTs(ct55.NFTs, ct55.Cli)
	if err != nil {
		log.Fatal(err)
	}

	sub, _ := t55instance.WatchClaimNFTs(nil, logs)
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vlog := <-logs:
			log.Println("claim from t55 nft url: ", vlog.Url)
			log.Println("claim from t55 nft tokenID: ", vlog.TokenID)
			log.Println("claim from t55 nft to: ", vlog.To)

			_, err := t55nfts.MintTo(ct55.T55NewKeyedTransactor(), vlog.To, vlog.TokenID, vlog.Url)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("NFTs mint from t54 to t55 successfully")
		}
	}
}

func ListenClaimT54NFTs() {
	ct54 := core.NewCoreT54()
	// ct54 := core.NewCoreT54()

	logs := make(chan *model2.T54BridgeClaimNFTs)
	t54instance, err := model2.NewT54Bridge(ct54.BridgeAddr, ct54.Cli)
	if err != nil {
		log.Fatal(err)
	}
	t54nfts, err := model2.NewT54NFTs(ct54.NFTs, ct54.Cli)
	if err != nil {
		log.Fatal(err)
	}

	sub, _ := t54instance.WatchClaimNFTs(nil, logs)
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vlog := <-logs:
			log.Println("claim from t54 nft url: ", vlog.Url)
			log.Println("claim from t54 nft tokenID: ", vlog.TokenID)
			log.Println("claim from t54 nft to: ", vlog.To)

			_, err := t54nfts.MintTo(ct54.T54NewKeyedTransactor(), vlog.To, vlog.TokenID, vlog.Url)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("NFTs mint from t55 to t54 successfully")
		}
	}
}
