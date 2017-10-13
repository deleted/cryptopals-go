package repeatingKey

import (
	"fmt"
	"sort"
)

func Break(cypher []byte) string {
	keySize := findKeysize(&cypher, 2, 40)
	fmt.Printf("Most likely keysize: %d\n", keySize)
	return "Dunno."
}

type keySizeTry struct {
	keysize            int
	normalizedDistance float64
}

func findKeysize(cypher *[]byte, minSize int, maxSize int) int {
	tries := make([]keySizeTry, 0, maxSize-minSize)
	for i := minSize; i <= maxSize; i++ {
		chunkA := (*cypher)[0:i]
		chunkB := (*cypher)[i : 2*i]
		dist := HammingDistance(string(chunkA), string(chunkB))
		normalizedDistance := float64(dist) / float64(i)
		tries = append(tries, keySizeTry{i, normalizedDistance})
	}
	sort.Slice(tries, func(i, j int) bool {
		return tries[i].normalizedDistance < tries[j].normalizedDistance
	})
	// fmt.Println(tries)
	return tries[0].keysize
}
