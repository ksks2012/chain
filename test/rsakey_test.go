package rsakey

import (
	"encoding/hex"
	"log"
	"testing"

	"github.com/block-chain/pkg/rsakey"
)

func TestRSAKey(t *testing.T) {
	prvKey, pubKey := rsakey.GenRsaKey()
	log.Printf("prv key: %x", prvKey)
	log.Printf("--------------------------------------------------------------------------------------")
	log.Printf("pub key: %x", pubKey)

	var data = "Hello RSA"
	log.Printf("sign data")
	signData := rsakey.RsaSignWithSha256([]byte(data), prvKey)
	log.Printf("signed msg： %x", hex.EncodeToString(signData))
	log.Printf("\nchecck signed msg")
	if rsakey.RsaVerySignWithSha256([]byte(data), signData, pubKey) {
		log.Printf("correct signed msg")
	}

	log.Printf("-------------------------------encode, decode-----------------------------------------")
	ciphertext := rsakey.RsaEncrypt([]byte(data), pubKey)
	log.Printf("encode by public key： %x", hex.EncodeToString(ciphertext))
	sourceData := rsakey.RsaDecrypt(ciphertext, prvKey)
	log.Printf("decode by private key： %v", string(sourceData))
}

// func TestChainWithRSA(t *testing.T) {
// 	blocker.GenAddress()
// 	for true {

// 	}
// }
