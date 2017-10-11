package main

import (
	"cryptopals/basics"
	"flag"
	"log"
)

func challenge_3() {
	/*
		Single-byte XOR cipher
		The hex encoded string:

		1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
		... has been XOR'd against a single character. Find the key, decrypt the message.

		You can do this by hand. But don't: write code to do it for you.

		How? Devise some method for "scoring" a piece of English plaintext. Character frequency is a good metric. Evaluate each output and choose the one with the best score.
	*/
	letterFrequencies := basics.LoadLetterFrequencies("./data/letter_frequencies.json")
	cypher := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	basics.BruteForceXorCrack(cypher, &letterFrequencies)
}

func main() {
	fnMap := map[string]func(){
		"3": challenge_3,
	}
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("I require a subcommand id")
	}
	subcommand := flag.Arg(0)
	fcn, exists := fnMap[subcommand]
	if !exists {
		log.Fatal("Unrecognised subcommand")
	}

	fcn()
	// challenge_3()
}