package test

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"testing"

	"example.com/m/app/config"
	"example.com/m/app/pkg/model/bridge"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// senbr nft deployed:  0xEcbDeCBBc722B6E2Bcd6bAc68e6392883196F242
// sen nft deployed:  0x1ee14deD878B2d98001132F8B72e99A82dEf749C

var (
	BridgeAddr = "0x52D9c9973CC07566126e20E4f679E67b1a333150"
	NFTs       = "0xB23B14f3eF282d33C2F5f3a8E9334385D99acCa9"
)

var appConfig = config.AppConfig

func PersonalSign(message string, privateKey *ecdsa.PrivateKey) []byte {
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%s%s", strconv.Itoa(len(message)), message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	signatureBytes, _ := crypto.Sign(hash.Bytes(), privateKey)
	// if err != nil {
	// 	return "", err
	// }
	// signatureBytes[64] += 27
	// return hexutil.Encode(signatureBytes), nil
	return signatureBytes
}

func TestSignature(t *testing.T) {
	config.Config()
	pvkey, _ := crypto.GenerateKey()
	privateKey, _ := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	pubkey := common.HexToAddress("0x8f9d9aA7B313cf9360d4E61D1Ae809443f97aCad")

	// message := []byte("www.example.com")
	// hash := crypto.Keccak256Hash(message)

	// signature, _ := crypto.Sign(hash.Bytes(), privateKey)
	signature := PersonalSign("www.example.com", privateKey)

	cli, err := ethclient.Dial(os.Getenv("URL_SEPOLIA"))
	if err != nil {
		t.Error(err)
	}
	senbridge, err := bridge.NewSentinelsBridge(common.HexToAddress(BridgeAddr), cli)
	if err != nil {
		t.Error(err)
	}
	// key, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	// if err != nil {
	// 	t.Error(err)
	// }
	_, errAuth := bind.NewKeyedTransactorWithChainID(pvkey, big.NewInt(appConfig.SepoliaChainID))
	if errAuth != nil {
		t.Error(errAuth)
	}
	verify, err := senbridge.Verify(nil, "www.example.com", signature)
	if err != nil {
		t.Error(err)
	}

	address, _ := senbridge.Recover(nil, "www.example.com", signature)
	sender, _ := senbridge.GetSender(&bind.CallOpts{From: pubkey})
	// hexData := hexutil.Encode(data)

	fmt.Println("Public key from private key: ", privateKey)
	fmt.Println("address: ", address.Hex())
	fmt.Println("sender: ", sender.Hex())
	fmt.Println("Verify", verify)

}
