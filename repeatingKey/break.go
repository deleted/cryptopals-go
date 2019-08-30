package repeatingKey

import (
	"github.com/deleted/cryptopals-go/basics"
	blocklib "github.com/deleted/cryptopals-go/blocks"
	"fmt"
	"sort"
)

func Break(cypher []byte) (key []byte, plaintext []byte) {
	keySize := findKeysize(&cypher, 2, 40)
	fmt.Printf("Most likely keysize: %d\n", keySize)
	blocks := blocklib.GetBlocks(cypher, keySize)
	fmt.Printf("%d blocks of size %d\n", len(blocks), len(blocks[0]))

	transposed := blocklib.Transpose(blocks)
	key = solveSubkeyBlocks(transposed)

	plaintext = RepeatingXor(cypher, key)
	return
}

type block = blocklib.Block
type blockpair = blocklib.Blockpair

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
		blocks := blocklib.GetBlocks(*cypher, keySize)
		numBlocks := len(blocks)
		pairs := make([]blockpair, 0, numBlocks*numBlocks)
		pairs = getAllPairs(blocks, pairs)
		sum := 0
		for _, pair := range pairs {
			sum += pair.Distance()
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

// Solve each of the blocks seperately as single-byte xor
// and recover the key
func solveSubkeyBlocks(blocks []block) (key []byte) {
	key = make([]byte, len(blocks))

	for i, block := range blocks {
		attempts := basics.BruteForceXorCrack(block)
		key[i] = attempts[0].Key
	}

	return key
}
