package aes

import (
	"log"
)

const paddingByte byte = 0x04

func PadBlock(block []byte, size int) []byte {
	if len(block) > size {
		log.Fatal("Block must be <= to the desired length after padding. Refusing to truncate.")
	}

	if len(block) == size {
		return block
	}

	output := make([]byte, size, size)
	for i := 0; i < size; i++ {
		if i < len(block) {
			output[i] = block[i]
		} else {
			output[i] = paddingByte
		}
	}

	return output
}

func PKCS7Pad(input []byte, blockSize int) []byte {
	if len(input)%blockSize == 0 {
		return input
	}
	out := input
	paddingLength := blockSize - (len(input) % blockSize)
	paddingByte := byte(paddingLength)
	for i := 0; i < paddingLength; i++ {
		out = append(out, paddingByte)
	}
	if len(out)%blockSize != 0 {
		log.Fatal("This is fucked!")
	}
	return out

}
