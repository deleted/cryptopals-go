package basics

import (
	"testing"
)

func TestHex2Base64(t *testing.T) {
	hexStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	val := Hex2base64(hexStr)
	if val != expected {
		t.Errorf("Expected %s; got %s", expected, val)
	}
}

func TestFixedXOR(t *testing.T) {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"
	val := FixedXOR(input1, input2)
	if val != expected {
		t.Errorf("Expected %s, got %s", expected, val)
	}
}
