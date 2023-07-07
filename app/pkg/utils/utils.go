package utils

import (
	"example.com/m/app/config"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"strconv"
)

type signature struct{}

var Signature = signature{}

func (sig *signature) StringSign(message string) (string, error) {
	privateKey, err := crypto.HexToECDSA(config.AppConfig.PrivateKeyBackend)
	if err != nil {
		return "", err
	}
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%s%s", strconv.Itoa(len(message)), message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}
	signatureBytes[64] += 27
	return hexutil.Encode(signatureBytes), nil
}

func (sig *signature) BytesSign(message string) ([]byte, error) {
	privateKey, err := crypto.HexToECDSA(config.AppConfig.PrivateKeyBackend)
	if err != nil {
		return nil, err
	}
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%s%s", strconv.Itoa(len(message)), message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}
	signatureBytes[64] += 27
	// return hexutil.Encode(signatureBytes)
	return signatureBytes, nil
}

func (sig *signature) Decode(input string) ([]byte, error) {
	return hexutil.Decode(input)
}
