package blocks

type Block []byte

type Blockpair struct {
	A Block
	B Block
}

func (bp Blockpair) Distance() int {
	return HammingDistance([]byte(bp.A), []byte(bp.B))
}

func GetBlocks(src []byte, size int) []Block {
	blocks := make([]Block, 0, len(src)%size)
	for i := 0; i+size < len(src); i += size {
		blocks = append(blocks, src[i:i+size])
	}
	return blocks
}

func Transpose(blocks []Block) []Block {
	blockSize := len(blocks[0])
	numBlocks := len(blocks)
	transposed := make([]Block, blockSize)
	for i := 0; i < blockSize; i++ {
		transposed[i] = make(Block, numBlocks)
	}
	for i := 0; i < blockSize; i++ {
		for j := 0; j < numBlocks; j++ {
			transposed[i][j] = blocks[j][i]
		}
	}

	return transposed
}
