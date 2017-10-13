package basics

import (
	"cryptopals/frequency"
	"sort"
	"unicode"
)

/*
Single-byte XOR cipher
The hex encoded string:

1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
... has been XOR'd against a single character. Find the key, decrypt the message.

You can do this by hand. But don't: write code to do it for you.

How? Devise some method for "scoring" a piece of English plaintext. Character frequency is a good metric. Evaluate each output and choose the one with the best score.
*/

const cyphertext string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

func SingleByteXOR(cypher string, key byte) string {
	bytes := Hex2Bytes(cypher)
	decoded := make([]byte, len(bytes))
	for i := 0; i < len(bytes); i++ {
		decoded[i] = bytes[i] ^ key
	}
	return string(decoded)
}

func isASCII(str string) bool {
	for _, r := range str {
		if r > unicode.MaxASCII {
			return false
		}
	}
	return true
}

type Attempt struct {
	Score  float64
	Key    byte
	Text   string
	Cypher string
}

func NewXorAttempt(cyphertext string, key byte) *Attempt {
	a := new(Attempt)
	a.Key = key
	a.Text = SingleByteXOR(cyphertext, key)
	a.Cypher = cyphertext
	a.Score = a.computeScore()
	return a
}

func (a Attempt) computeScore() float64 {
	score := 0.0

	// Any non-ascii characters immediately disqualify
	if !isASCII(a.Text) {
		score--
	}

	// strings get points for having letters and spaces, lose points for having non-ascii characters
	for _, r := range a.Text {
		if r > 'a' && r < 'z' {
			score++
		}
		if r == ' ' {
			score += 2
		}
	}

	score += 100 * frequency.ComputeFrequencyScore(a.Text)

	return score
}

func BruteForceXorCrack(cyphertext string) []*Attempt {
	attempts := make([]*Attempt, 0, 256)
	for i := 0x00; i <= 0xff; i++ {
		a := NewXorAttempt(cyphertext, byte(i))
		attempts = append(attempts, a)
	}

	// Sort by score
	sort.Slice(attempts, func(i, j int) bool { return attempts[i].Score > attempts[j].Score })
	return attempts
}
