package codebreaking

import (
	"bytes"
	myAES "cryptopals/aes"
	"cryptopals/blocks"
	"cryptopals/oracle"
	"log"
)

type encrypter interface {
	Encrypt([]byte) []byte
}

func detectBlockSize(oracle encrypter) int {
	nullLen := len(oracle.Encrypt(make([]byte, 0)))
	var cypherLen int
	var text []byte
	a := []byte("A")
	i := 1
	for {
		text = bytes.Repeat(a, i)
		cypherLen = len(oracle.Encrypt(text))
		if cypherLen > nullLen {
			break
		}
		i++
	}
	blockSize := cypherLen - nullLen
	return blockSize
}

func BreakAppendingOracle() []byte {
	oracle := oracle.MakeAppendingOracle(16)
	a := []byte("A")
	isECB := myAES.DetectECB(oracle.Encrypt(bytes.Repeat(a, 500)))
	if !isECB {
		log.Fatal("This cypher does not appear to be ECB mode.")
	}

	blockSize := detectBlockSize(oracle)
	secretLen := len(oracle.Encrypt(make([]byte, 0)))
	log.Printf("Detected block size %d and secret length %d\n", blockSize, secretLen)

	plainText := make([]byte, 0)
	for len(plainText) < secretLen {
		nextByte := recoverNextByte(oracle, blockSize, plainText)
		plainText = append(plainText, nextByte)
		// log.Printf("Known bytes: %s\n", string(plainText))
	}

	return plainText
}

func recoverNextByte(oracle encrypter, blockSize int, knownBytes []byte) byte {
	a := []byte("A")
	var numBytesToPrepend int
	if len(knownBytes) == 0 {
		numBytesToPrepend = blockSize - 1
	} else {
		numBytesToPrepend = blockSize - (len(knownBytes) % blockSize) - 1
	}
	prefix := bytes.Repeat(a, numBytesToPrepend)

	activeBlock := len(knownBytes) / blockSize

	cypherToByte := make(map[string]byte)
	var testSlice []byte
	var testCypher []byte
	for i := byte(0); i < 255; i++ {
		testSlice = append(prefix, knownBytes...)
		testSlice = append(testSlice, i)
		testCypher = oracle.Encrypt(testSlice)
		testBlock := blocks.GetBlocks(testCypher, blockSize)[activeBlock]
		cypherToByte[string(testBlock)] = i
	}
	allBlocks := blocks.GetBlocks(oracle.Encrypt(prefix), blockSize)
	if activeBlock == len(allBlocks) {
		// On the very last byte, we wind up one block to the left (no prefix padding)
		activeBlock--
	}
	encypheredBlock := allBlocks[activeBlock]
	solution := cypherToByte[string(encypheredBlock)]
	return solution
}
