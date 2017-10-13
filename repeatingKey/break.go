package repeatingKey

import (
	"cryptopals/basics"
	"fmt"
	"sort"
)

func Break(cypher []byte) (key []byte, plaintext []byte) {
	keySize := findKeysize(&cypher, 2, 40)
	fmt.Printf("Most likely keysize: %d\n", keySize)
	blocks := getBlocks(cypher, keySize)
	fmt.Printf("%d blocks of size %d\n", len(blocks), len(blocks[0]))

	transposed := transpose(blocks)
	key = solveSubkeyBlocks(transposed)

	plaintext = RepeatingXor(cypher, key)
	return
}

type block []byte

type blockpair struct {
	a block
	b block
}

func (bp blockpair) distance() int {
	return HammingDistance([]byte(bp.a), []byte(bp.b))
}

func getAllPairs(blocks []block, dest []blockpair) []blockpair {
	if len(blocks) == 1 {
		return dest
	}
	head := blocks[0]
	rest := blocks[1:]
	for _, block := range rest {
		dest = append(dest, blockpair{head, block})
	}
	dest = getAllPairs(rest, dest)

	return dest
}

type keySizeTry struct {
	keysize            int
	normalizedDistance float64
}

func findKeysize(cypher *[]byte, minSize int, maxSize int) int {
	tries := make([]keySizeTry, 0, maxSize-minSize)
	for keySize := minSize; keySize <= maxSize; keySize++ {
		blocks := getBlocks(*cypher, keySize)
		numBlocks := len(blocks)
		pairs := make([]blockpair, 0, numBlocks*numBlocks)
		pairs = getAllPairs(blocks, pairs)
		sum := 0
		for _, pair := range pairs {
			sum += pair.distance()
		}
		avgDistance := float64(sum) / float64(len(pairs))
		normalizedDistance := avgDistance / float64(keySize)
		tries = append(tries, keySizeTry{keySize, normalizedDistance})
	}
	sort.Slice(tries, func(i, j int) bool {
		return tries[i].normalizedDistance < tries[j].normalizedDistance
	})
	return tries[0].keysize
}

func getBlocks(src []byte, size int) []block {
	blocks := make([]block, 0, len(src)%size)
	for i := 0; i+size < len(src); i += size {
		blocks = append(blocks, src[i:i+size])
	}
	return blocks
}

func transpose(blocks []block) []block {
	blockSize := len(blocks[0])
	numBlocks := len(blocks)
	transposed := make([]block, blockSize)
	for i := 0; i < blockSize; i++ {
		transposed[i] = make(block, numBlocks)
	}
	for i := 0; i < blockSize; i++ {
		for j := 0; j < numBlocks; j++ {
			transposed[i][j] = blocks[j][i]
		}
	}

	return transposed
}

// Solve each of the blocks seperately as single-byte xor
// and recover the key
func solveSubkeyBlocks(blocks []block) (key []byte) {
	key = make([]byte, len(blocks))

	for i, block := range blocks {
		attempts := basics.BruteForceXorCrack(basics.Bytes2Hex(block))
		key[i] = attempts[0].Key
	}

	return key
}
