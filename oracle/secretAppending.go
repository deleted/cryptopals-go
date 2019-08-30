package oracle

/*
The oracle here is related to challenge 12, byte-at-a time ECB decryption
https://cryptopals.com/sets/2/challenges/12
*/

import (
	myAES "github.com/deleted/cryptopals-go/aes"
	"encoding/base64"
	"log"
	"math/rand"
)

const secretB64Encoded string = "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"

/*
MakeAppendingOracle will decode the secret string, generaate a random secret key,
and create an oracle that can encrypt a string of text using these secrets.
*/
func MakeAppendingOracle(blockSize int) AppendingOracle {
	secretBytes, err := base64.StdEncoding.DecodeString(secretB64Encoded)
	if err != nil {
		log.Fatal(err)
	}
	var secretKey = make([]byte, blockSize)
	rand.Read(secretKey)
	return AppendingOracle{
		secretKey:     secretKey,
		appendingText: secretBytes,
	}

}

type AppendingOracle struct {
	secretKey     []byte
	appendingText []byte
}

func (ao AppendingOracle) Encrypt(plaintext []byte) []byte {
	augmentedPlaintext := append(plaintext, ao.appendingText...)
	cypher := myAES.EncryptECB(augmentedPlaintext, ao.secretKey)
	return cypher
}
