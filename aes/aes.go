package aes

import (
	"bytes"
	"crypto/aes"
	"fmt"
	"log"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func EncryptECB(plainBytes, key []byte) []byte {
	blockSize := len(key)
	if len(plainBytes)%blockSize != 0 {
		plainBytes = PKCS7Pad(plainBytes, blockSize)
	}
	inBuf := bytes.NewBuffer(plainBytes)
	outBuf := new(bytes.Buffer)
	cypher, err := aes.NewCipher(key)
	check(err)

	dst := make([]byte, blockSize)
	for inBuf.Len() > 0 {
		cypher.Encrypt(dst, inBuf.Next(blockSize))
		outBuf.Write(dst)
	}
	return outBuf.Bytes()
}

func DecryptECB(cypherBytes, key []byte) []byte {
	cypher, _ := aes.NewCipher(key)
	plainText := make([]byte, len(cypherBytes))
	blockSize := cypher.BlockSize()

	// This is all there is to ECB mode, I guess...
	src := cypherBytes
	dst := plainText
	for len(src) > 0 {
		cypher.Decrypt(dst, src)
		src = src[blockSize:]
		dst = dst[blockSize:]
	}
	return plainText
}

func blockXOR(a, b []byte) []byte {
	if len(a) != len(b) {
		log.Fatal("input blocks must be the same size")
	}
	out := make([]byte, len(a))
	for i := range a {
		out[i] = a[i] ^ b[i]
	}
	return out
}

func EncryptCBC(clearBytes, key, iv []byte) (cypherBytes []byte) {
	blockSize := len(key)
	if len(iv) != blockSize {
		log.Fatal("Initialization vector must be the same size as the key.")
	}
	cypher, err := aes.NewCipher(key)
	check(err)
	prevBlock := iv
	dest := make([]byte, blockSize)
	inputBuf := bytes.NewBuffer(clearBytes)
	outputBuf := new(bytes.Buffer)

	for inputBuf.Len() > 0 {
		nextBlock := inputBuf.Next(blockSize)
		nextBlock = PKCS7Pad(nextBlock, blockSize)
		src := blockXOR(nextBlock, prevBlock)
		cypher.Encrypt(dest, src)
		prevBlock = dest
		n, err := outputBuf.Write(dest)
		check(err)
		if n != blockSize {
			log.Fatal("Wrote the wrong blocksize")
		}
	}

	cypherBytes = outputBuf.Bytes()
	return cypherBytes

}

func DecryptCBC(cypherBytes, key, iv []byte) []byte {
	blockSize := len(key)
	if len(iv) != blockSize {
		log.Fatal("Initialization vector must be the same size as the key.")
	}
	cypher, err := aes.NewCipher(key)
	check(err)

	inBuf := bytes.NewBuffer(cypherBytes)
	outBuf := new(bytes.Buffer)
	prevBlock := iv
	dest := make([]byte, blockSize)
	for inBuf.Len() > 0 {
		block := inBuf.Next(blockSize)
		if len(block) != blockSize {
			log.Fatal(fmt.Sprintf("Truncated input block of length %d (expected %d)", len(block), blockSize))
		}
		cypher.Decrypt(dest, block)
		outBuf.Write(blockXOR(dest, prevBlock))
		prevBlock = block
	}

	return outBuf.Bytes()
}
