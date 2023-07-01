package test

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestSignature(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		log.Fatal("private_key: ", err)
	}

	message := []byte("www.example.com")
	hash := crypto.Keccak256Hash(message)

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal("signature : ", err)
	}

	signatureHex := hexutil.Encode(signature)
	fmt.Println("ECDSA Signature:", signatureHex)
}
