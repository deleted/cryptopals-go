package codebreaking

import "testing"
import "github.com/deleted/cryptopals-go/oracle"

func TestDetectBlockSize(t *testing.T) {
	oracle := oracle.MakeAppendingOracle(16)
	blocksize := detectBlockSize(oracle)
	if blocksize != 16 {
		t.Error("detectBlockSize is broken")
	}
}
