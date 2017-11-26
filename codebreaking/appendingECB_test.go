package codebreaking

import "testing"
import "cryptopals/oracle"

func TestDetectBlockSize(t *testing.T) {
	oracle := oracle.MakeAppendingOracle(16)
	blocksize := detectBlockSize(oracle)
	if blocksize != 16 {
		t.Error("detectBlockSize is broken")
	}
}
