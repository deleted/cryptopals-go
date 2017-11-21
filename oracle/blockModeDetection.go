package oracle

import (
	"bytes"
	myAES "cryptopals/aes"
	"fmt"
	"log"
	"math/rand"
)

/*
An ECB/CBC detection oracle
Now that you have ECB and CBC working:

Write a function to generate a random AES key; that's just 16 random bytes.

Write a function that encrypts data under an unknown key --- that is, a function that generates a random key and encrypts under it.

The function should look like:

encryption_oracle(your-input)
=> [MEANINGLESS JIBBER JABBER]
Under the hood, have the function append 5-10 bytes (count chosen randomly) before the plaintext and 5-10 bytes after the plaintext.

Now, have the function choose to encrypt under ECB 1/2 the time, and under CBC the other half (just use random IVs each time for CBC). Use rand(2) to decide which to use.

Detect the block cipher mode the function is using each time. You should end up with a piece of code that, pointed at a block box that might be encrypting ECB or CBC, tells you which one is happening.
*/

const (
	ECB = 0
	CBC = 1
)

/*
EncryptionOracle encrypts the input plaintext with a random mode (ECB or CBC)
It then tests the cypher text against DetectECB and returns a boolean indicating
whether it was ablet to guess the mode correctly.
*/
func EncryptionOracle(input []byte) (success bool) {
	mode := rand.Intn(2)
	key := randomBytes(16)
	input = randomPad(input)
	var cypherbytes []byte

	switch mode {
	case ECB:
		cypherbytes = myAES.EncryptECB(input, key)
	case CBC:
		iv := randomBytes(len(key))
		cypherbytes = myAES.EncryptCBC(input, key, iv)
	default:
		log.Fatal(fmt.Sprintf("Unknown mode: %d", mode))
	}

	ecbDetected := myAES.DetectECB(cypherbytes)
	if ecbDetected == (mode == ECB) {
		success = true
		fmt.Printf("Mode %d detected successfully!\n", mode)
	} else {
		success = false
		fmt.Printf("I chose poorly. (mode %d)\n", mode)
	}
	return
}

func randomBytes(n int) []byte {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return bytes
}

func randomPad(input []byte) []byte {
	out := new(bytes.Buffer)
	out.Write(randomBytes(rand.Intn(5) + 5))
	out.Write(input)
	out.Write(randomBytes(rand.Intn(5) + 5))
	return out.Bytes()
}
