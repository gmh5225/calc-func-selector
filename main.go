package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"strings"

	"golang.org/x/crypto/sha3"
)

func calculateFunctionSelector(signature string, verbose bool) string {
	// Create a Keccak-256 hash object
	hasher := sha3.NewLegacyKeccak256()

	// Write the function signature
	hasher.Write([]byte(signature))

	// Calculate the hash
	hash := hasher.Sum(nil)

	if verbose {
		fmt.Printf("hash[0]: %x\n", hash[0])
		fmt.Printf("hash[1]: %x\n", hash[1])
		fmt.Printf("hash[2]: %x\n", hash[2])
		fmt.Printf("hash[3]: %x\n", hash[3])
	}

	// Take the first 4 bytes of the hash and convert to hexadecimal
	selector := hex.EncodeToString(hash[:4])

	return "0x" + selector
}

func main() {

	// Example function signature
	signature := flag.String("signature", "decimals()", "The function signature to calculate the selector for")
	verbose := flag.Bool("verbose", false, "Print verbose output")

	flag.Parse()

	// Remove the prefix "function "
	*signature = strings.TrimPrefix(*signature, "function ")

	// Calculate the function selector
	selector := calculateFunctionSelector(*signature, *verbose)

	if *verbose {
		fmt.Printf("Function signature: %s\n", *signature)
		fmt.Printf("Function selector: %s\n", selector)
	}

	fmt.Println(selector)
}
