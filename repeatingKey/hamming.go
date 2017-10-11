package repeatingKey

import (
	"log"
	"math/bits"
)

func HammingDistance(x, y string) int {
	if len(x) != len(y) {
		log.Fatal("inputs must be the same length")
	}
	dist := 0
	for i := 0; i < len(x); i++ {
		xor := x[i] ^ y[i]
		dist += bits.OnesCount8(uint8(xor))
	}
	return dist
}
