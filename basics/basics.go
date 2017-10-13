package basics

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func Hex2Bytes(input string) []byte {
	bytes := make([]byte, hex.DecodedLen(len(input)))
	hex.Decode(bytes, []byte(input))
	return bytes
}

func Bytes2Hex(input []byte) string {
	return hex.EncodeToString(input)
}

func Hex2base64(hexInput string) string {
	bytes := Hex2Bytes(hexInput)
	return base64.StdEncoding.EncodeToString(bytes)
}

func FixedXOR(hexA, hexB string) string {
	bytesA := Hex2Bytes(hexA)
	bytesB := Hex2Bytes(hexB)
	if len(bytesA) != len(bytesB) {
		log.Fatal("Inputs need to be the same length")
	}
	outBytes := make([]byte, len(bytesA))
	for i := 0; i < len(bytesA); i++ {
		outBytes[i] = bytesA[i] ^ bytesB[i]
	}
	return hex.EncodeToString(outBytes)
}
