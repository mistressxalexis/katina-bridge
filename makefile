gen:
	npx hardhat export-abi

gen-sepo:
	npx hardhat export-abi --network sepolia

gen-ftm:
	npx hardhat export-abi --network ftm

abigen:
	abigen --abi ./app/pkg/resource/T55NFTs.json --pkg model --type T55NFTs --out ./app/pkg/model/T55NFTs.go
	abigen --abi ./app/pkg/resource/T54NFTs.json --pkg model --type T54NFTs --out ./app/pkg/model/T54NFTs.go
	abigen --abi ./app/pkg/resource/T55Bridge.json --pkg model --type T55Bridge --out ./app/pkg/model/T55Bridge.go
	abigen --abi ./app/pkg/resource/T54Bridge.json --pkg model --type T54Bridge --out ./app/pkg/model/T54Bridge.go

	abigen --abi ./app/pkg/resource/LEO.json --pkg nft --type Leopard --out ./app/pkg/model/NFT/leopard.go
	abigen --abi ./app/pkg/resource/SEN.json --pkg nft --type Sentinels --out ./app/pkg/model/NFT/sentinels.go
	abigen --abi ./app/pkg/resource/LeopardBridge.json --pkg bridge --type LeopardBridge --out ./app/pkg/model/bridge/LeopardBridge.go
	abigen --abi ./app/pkg/resource/SentinelsBridge.json --pkg bridge --type SentinelsBridge --out ./app/pkg/model/bridge/SentinelsBridge.go

run:
	go run cmd/main.go

dep-bsc:
	npx hardhat run scripts/deploy.ts --network bsctestnet

dep-sepo:
	npx hardhat run scripts/deploy.ts --network sepolia

dep-ftm:
	npx hardhat run scripts/deploy.ts --network ftm

docker-run:
	docker-compose up --build