package basics

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"unicode"
)

// Alphabet gives the order of the letters
const Alphabet string = "abcdefghijklmnopqrstuvwxyz"

// Read a json file of english letter frequencies and return a slice of 26 floats, giving relative frequency in alphabetical order
func LoadLetterFrequencies(filename string) []float64 {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panic("Can't open frequencies file.")
	}
	freqMap := make(map[string]float64)
	json.Unmarshal(dat, &freqMap)
	frequencies := make([]float64, len(Alphabet))
	for i := 0; i < len(Alphabet); i++ {
		frequencies[i] = freqMap[string(Alphabet[i])]
	}
	return frequencies
}

/*
 AnalyzeLetterFrequency analyzes a string for letter frequencies and outputs a normalized 26-vector
 corresponding to letters of the alphabet.
*/
func AnalyzeLetterFrequency(input string) []float64 {
	letterCounter := make(map[rune]int)
	for _, r := range input {
		if unicode.IsLetter(r) {
			letterCounter[unicode.ToLower(r)]++
		}
	}

	total := 0.0
	for _, count := range letterCounter {
		total += float64(count)
	}

	frequencies := make([]float64, 26)
	for i, r := range Alphabet {
		var f float64
		f = float64(letterCounter[r]) / total
		frequencies[i] = f
	}
	return frequencies
}

func CosineSimilarity(a []float64, b []float64) (cosine float64, err error) {
	count := 0
	lengthA := len(a)
	lengthB := len(b)
	if lengthA > lengthB {
		count = lengthA
	} else {
		count = lengthB
	}
	sumA := 0.0
	s1 := 0.0
	s2 := 0.0
	for k := 0; k < count; k++ {
		if k >= lengthA {
			// s2 += math.Pow(b[k], 2)
			s2 += b[k] * b[k]
			continue
		}
		if k >= lengthB {
			// s1 += math.Pow(a[k], 2)
			s1 += a[k] * a[k]
			continue
		}
		sumA += a[k] * b[k]
		// s1 += math.Pow(a[k], 2)
		// s2 += math.Pow(b[k], 2)
		s1 += a[k] * a[k]
		s2 += b[k] * b[k]
	}
	if s1 == 0 || s2 == 0 {
		return 0.0, errors.New("Vectors should not be null (all zeros)")
	}
	if s1 < 0 || s2 < 0 {
		return 0.0, errors.New(fmt.Sprintf("One of these is < 0: %f %f", s1, s2))
	}
	result := sumA / (math.Sqrt(s1) * math.Sqrt(s2))
	if math.IsNaN(result) {
		return 0.0, nil
	}
	return result, nil
}

func ComputeFrequencyScore(input string, letterFrequencies *[]float64) float64 {
	frequencies := AnalyzeLetterFrequency(input)
	sum := 0.0
	for _, f := range frequencies {
		sum += f
	}
	if sum == 0.0 {
		// this string contains no letters.
		return -1.0
	}
	similarity, err := CosineSimilarity(frequencies, *letterFrequencies)
	if err != nil {
		log.Fatal(err)
	}
	return similarity
}
