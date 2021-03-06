package aes

import "testing"
import "bytes"

func TestPadByte(t *testing.T) {
	testBlock := []byte("YELLOW SUBMARINE")
	expectedResult := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")
	result := PKCS7Pad(testBlock, 20)
	if !bytes.Equal(result, expectedResult) {
		t.Fail()
	}
}

func TestEBCRoundTrip(t *testing.T) {
	testTxt := []byte("scoobidy do derp dip flight attendante prepare for arrival etc. doo ya!!!")
	key := []byte("YELLOW SUBMARINE")
	cypherBytes := EncryptECB(testTxt, key)
	plainBytes := DecryptECB(cypherBytes, key)
	if !bytes.Equal(plainBytes, PKCS7Pad(testTxt, len(key))) {
		t.Errorf("Wrong result: %s", plainBytes)
	}
}

func TestCBCRoundTrip(t *testing.T) {
	testTxt := []byte("scoobidy do derp dip flight attendante prepare for arrival etc. doo ya!!!")
	key := []byte("YELLOW SUBMARINE")
	zero := []byte{0x00}
	iv := bytes.Repeat(zero, len(key))

	cypherBytes := EncryptCBC(testTxt, key, iv)
	plainBytes := DecryptCBC(cypherBytes, key, iv)
	// fmt.Println(plainBytes)
	if !bytes.Equal(plainBytes, PKCS7Pad(testTxt, len(key))) {
		t.Error(string(plainBytes))
	}
}

func TestPKCS7Padding(t *testing.T) {
	testBlock := []byte("abcdef")
	var padded []byte
	padded = PKCS7Pad(testBlock, 8)
	if !bytes.Equal(padded, []byte("abcdef\x02\x02")) {
		t.Error("Unexpected padding result", string(padded))
	}
	padded = PKCS7Pad(testBlock, 10)
	if !bytes.Equal(padded, []byte("abcdef\x04\x04\x04\x04")) {
		t.Error("Unexpected padding result", string(padded))
	}
	padded = PKCS7Pad(testBlock, 6)
	if !bytes.Equal(padded, []byte("abcdef")) {
		t.Error("Unexpected padding result", string(padded))
	}
}
