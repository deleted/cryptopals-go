package aes

import (
	"bytes"
	blocklib "cryptopals/blocks"
	"fmt"
	"sort"
)

func DetectECB(input []byte) (isECB bool) {
	for blockSize := 16; blockSize >= 4; blockSize /= 2 {
		bc := tallyBlocks(input, blockSize)
		if anyRepeats(bc) {
			return true
		}
	}
	return false
}

type blockCount struct {
	block blocklib.Block
	count int
}

func (bc blockCount) String() string {
	return fmt.Sprintf("[%d %x]", bc.count, bc.block)
}

func tallyBlocks(input []byte, blockSize int) []blockCount {
	blocks := blocklib.GetBlocks(input, blockSize)
	blockCounts := make([]blockCount, 0, len(blocks))
	for _, block := range blocks {
		count := bytes.Count(input, []byte(block))
		blockCounts = append(blockCounts, blockCount{block, count})
	}

	sort.Slice(blockCounts, func(i, j int) bool { return blockCounts[i].count > blockCounts[j].count })
	return blockCounts
}

func anyRepeats(counts []blockCount) bool {
	for _, bc := range counts {
		if bc.count > 1 {
			return true
		}
	}
	return false
}
