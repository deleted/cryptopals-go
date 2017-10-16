package aes

import "testing"
import "bytes"

func TestPadByte(t *testing.T) {
	testBlock := []byte("YELLOW SUBMARINE")
	expectedResult := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")
	result := PadBlock(testBlock, 20)
	if !bytes.Equal(result, expectedResult) {
		t.Fail()
	}
}

func TestCBCRoundTrip(t *testing.T) {
	testTxt := []byte("scoobidy do derp dip flight attendante prepare for arrival etc. doo ya!!!")
	key := []byte("YELLOW SUBMARINE")
	zero := []byte("\x00")
	iv := bytes.Repeat(zero, len(key))

	cypherBytes := EncryptCBC(testTxt, key, iv)
	plainBytes := DecryptCBC(cypherBytes, key, iv)
	// fmt.Println(plainBytes)
	if !bytes.Equal(plainBytes, padToBlockSize(testTxt, len(key))) {
		t.Error(string(plainBytes))
	}
}
