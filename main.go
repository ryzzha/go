package main

import (
	"fmt"
	"math/big"
	"strings"
	// "log"
	// "net/http"
	// "os"
	// "time"
)

func main() {
	fmt.Println("Enter hex number or 'stop' to exit:")

	var input string

	for {
		fmt.Scanln(&input)

		input = strings.ToLower(input)

		if input == "stop" {
			fmt.Println("Exiting...")
			break
		}

		i := new(big.Int)

		if _, ok := i.SetString(processHex(input), 16); !ok {
			fmt.Println("Invalid hexadecimal number!")
			continue
		}

		fmt.Println(i)
	}
}

func processHex(hexStr string) string {
	return strings.TrimPrefix(hexStr, "0x")
}