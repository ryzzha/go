package main

import (
	"fmt"
	// "math/big"
	"strings"
	"flag"
	"bufio"
	"os"

	"cipherer/pkg/cipherer"
)

var mode = flag.String("mode", "cipher", "Set to 'cipher' or 'decipher'. Default is 'cipher'.")
var secretKey = flag.String("secret", "", "Your secret key. Must contain at least 1 character")

func main() {
	flag.Parse()

	if len(*secretKey) < 1 {
		fmt.Fprintln(os.Stderr, "No secret is provided! Exiting now...")
		os.Exit(1)
	}

	switch *mode {
	case "cipher":
		input := getUserInput("Enter the text to cipher: ");

		cipheredText, err := cipherer.Cipher(input, *secretKey)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error encrypting text: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(cipheredText)
	case "decipher":
		cipheredText := getUserInput("Enter your ciphered data to decipher: ")

		decipheredText, err := cipherer.Decipher(cipheredText, *secretKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error decrypting text: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(decipheredText)
	default:
		fmt.Println("Invalid mode. Use 'cipher' or 'decipher'.")
		os.Exit(1)	

	}
}

func getUserInput(msg string) string {
	fmt.Println(msg)

	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("An error occurred while reading the entered text! Please try again...")
			continue
		}

		return strings.TrimRight(input, "\r\n")
	}
}