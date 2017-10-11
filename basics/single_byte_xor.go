package basics

import (
	"fmt"
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

type runeFreq struct {
	r     rune
	count int
}

type runeCounter map[rune]int

func (rc runeCounter) consume(input string) {
	for _, r := range input {
		rc[r]++
	}
}

// Return the n most frequent runes
func (rc runeCounter) top(n int) []runeFreq {
	runeCounts := make([]runeFreq, 0, len(rc))
	for r, count := range rc {
		rf := runeFreq{r, count}
		runeCounts = append(runeCounts, rf)
	}
	sort.Slice(runeCounts, func(i, j int) bool { return runeCounts[i].count > runeCounts[j].count })
	if len(runeCounts) >= n {
		return runeCounts[0:n]
	}
	return runeCounts

}

func countRunes(input string) runeCounter {
	results := make(runeCounter)
	results.consume(input)
	return results
}

type Attempt struct {
	key   byte
	text  string
	freqs runeCounter
	score float64
}

func NewXorAttempt(cyphertext string, key byte, letterFreqs *[]float64) *Attempt {
	a := new(Attempt)
	a.key = key
	a.text = SingleByteXOR(cyphertext, key)
	a.freqs = make(runeCounter)
	a.freqs.consume(a.text)
	a.score = a.computeScore(letterFreqs)
	// a.score = ComputeFrequencyScore(a.text, letterFreqs)
	return a
}

func (a Attempt) computeScore(letterFreqs *[]float64) float64 {
	score := 0.0

	// Any non-ascii characters immediately disqualify
	if !isASCII(a.text) {
		return 0.0
	}

	// strings get points for having letters and spaces, lose points for having non-ascii characters
	for _, r := range a.text {
		if r > 'a' && r < 'z' {
			score++
		}
		if r == ' ' {
			score += 2
		}
	}

	score += 100 * ComputeFrequencyScore(a.text, letterFreqs)

	// boost score for having frequent lettsers in the top runes.
	// topRunes := make(map[rune]bool)
	// for _, rf := range a.freqs.top(10) {
	// 	topRunes[rf.r] = true
	// }
	// letterFreqs := "etaoinsrhldcumfpgwybvkxjqz" // most common letters in the english language
	// bonus := len(letterFreqs)
	// for _, l := range letterFreqs {
	// 	if topRunes[l] || topRunes[unicode.ToUpper(l)] {
	// 		score += bonus
	// 	}
	// 	bonus--
	// }
	return score
}

func BruteForceXorCrack(cyphertext string, letterFreqs *[]float64) {
	attempts := make([]*Attempt, 0, 256)
	for i := 0x00; i <= 0xff; i++ {
		a := NewXorAttempt(cyphertext, byte(i), letterFreqs)
		attempts = append(attempts, a)
	}
	sort.Slice(attempts, func(i, j int) bool { return attempts[i].score > attempts[j].score })
	for _, a := range attempts {
		fmt.Printf("%f: %q\n", a.score, a.text)
	}
}
