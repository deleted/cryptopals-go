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

func padToBlockSize(input []byte, blockSize int) []byte {
	if len(input)%blockSize == 0 {
		return input
	}
	out := input
	for len(out)%blockSize != 0 {
		out = append(out, paddingByte)
	}
	return out
}

// Break src into blocks of the given size, padding the end if nessecary
func getBlocks(src []byte, size int) [][]byte {
	src = padToBlockSize(src, size)
	blocks := make([][]byte, 0, len(src)/size)
	for i := 0; i+size < len(src); i += size {
		blocks = append(blocks, src[i:i+size])
	}
	return blocks
}
