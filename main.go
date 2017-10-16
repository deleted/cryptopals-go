package main

import (
	"bufio"
	"bytes"
	myAES "cryptopals/aes"
	"cryptopals/basics"
	"cryptopals/repeatingKey"
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func challenge_3(args ...string) {
	/*
		Single-byte XOR cipher
		The hex encoded string:

		1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
		... has been XOR'd against a single character. Find the key, decrypt the message.

		You can do this by hand. But don't: write code to do it for you.

		How? Devise some method for "scoring" a piece of English plaintext. Character frequency is a good metric. Evaluate each output and choose the one with the best score.
	*/
	cypher := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	results := basics.BruteForceXorCrack(basics.Hex2Bytes(cypher))
	fmt.Println(results[0].Text)
}

func challenge_4(args ...string) {
	datafilename := "./data/4.txt"
	result := basics.ParallelXorSolve(datafilename)
	fmt.Println(result)
}

func challenge_5(args ...string) {

	if len(args) < 2 {
		log.Fatal("Key required")
	}
	key := args[1]

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString("\n"[0])
	output := repeatingKey.Encrypt(input, key)
	fmt.Println(output)
}

func challenge_6(args ...string) {
	b64EncodedCypher, _ := ioutil.ReadFile("./data/6.txt")
	cypherBytes := make([]byte, len(b64EncodedCypher))
	_, _ = base64.StdEncoding.Decode(cypherBytes, b64EncodedCypher)
	key, solution := repeatingKey.Break(cypherBytes)
	fmt.Printf("Key: %s (%x) \n", string(key), string(key))
	fmt.Print(string(solution))
}

func challenge_7(args ...string) {
	b64EncodedCypher, err := ioutil.ReadFile("./data/7.txt")
	cypherBytes := make([]byte, len(b64EncodedCypher))
	bytesRead, err := base64.StdEncoding.Decode(cypherBytes, b64EncodedCypher)
	fmt.Printf("Read %d bytes\n", bytesRead)
	if err != nil {
		log.Fatal()
	}
	key := "YELLOW SUBMARINE"
	plainText := myAES.DecryptECB(cypherBytes, []byte(key))
	fmt.Println(string(plainText))
}

func challenge_8(args ...string) {
	filename := "./data/8.txt"
	for _, line := range readLines(filename) {
		cypher, err := hex.DecodeString(line)
		check(err)
		ecb := myAES.DetectECB(cypher)
		if ecb {
			fmt.Println("ECB Detected")
			fmt.Println(basics.Bytes2Hex(cypher))
		}
	}
}

func challenge_10(args ...string) {
	filename := "./data/10.txt"
	b64EncodedCypher, err := ioutil.ReadFile(filename)
	cypherBytes := make([]byte, len(b64EncodedCypher))
	bytesRead, err := base64.StdEncoding.Decode(cypherBytes, b64EncodedCypher)

	fmt.Printf("Read %d bytes\n", bytesRead)
	key := []byte("YELLOW SUBMARINE")
	iv := bytes.Repeat([]byte{0x00}, len(key))
	check(err)
	fmt.Printf("IV: %x\n", iv)
	clearBytes := myAES.DecryptCBC(cypherBytes, key, iv)
	fmt.Print(string(clearBytes))
}

func main() {
	fnMap := map[string]func(...string){
		"3":  challenge_3,
		"4":  challenge_4,
		"5":  challenge_5,
		"6":  challenge_6,
		"7":  challenge_7,
		"8":  challenge_8,
		"10": challenge_10,
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

	fcn(flag.Args()...)
}

func readLines(filename string) []string {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	str := string(dat)
	return strings.Split(str, "\n")

}
