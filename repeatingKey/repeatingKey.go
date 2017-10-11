package repeatingKey

import (
	"encoding/hex"
)

func repeatingXor(input []byte, key []byte) []byte {
	output := make([]byte, len(input))
	keyLength := len(key)
	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ key[i%keyLength]
	}
	return output
}

func Encrypt(plaintext string, key string) string {
	textBytes := []byte(plaintext)
	keyBytes := []byte(key)
	cypher := repeatingXor(textBytes, keyBytes)
	return hex.EncodeToString(cypher)
}
