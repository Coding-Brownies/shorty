package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

// implementation of our shortener algorithm and the choices behind the design decisions in hash function and byte to text encoder

// We will be using  SHA256 to hash the initial inputs during the process
func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortLink(initialLink string) string {
	// Hashing  initialUrl + userId url with sha256.  Here userId is added to prevent providing similar shortened urls to separate users in case they want to shorten exact same link, it's a design decision, so some implementations do this differently
	urlHashBytes := sha256Of(initialLink)
	// Derive a big integer number from the hash bytes generated during the hasing
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	// Apply base58  on the derived big integer value and pick the first 8 characters
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}
