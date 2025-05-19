package main

import (
	"fmt"
)

func main() {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}


	fmt.Printf("%5s", "a\\b")

	for _, el := range numbers {
		fmt.Printf("%5d", el)
	}

	fmt.Println("\n") 

	for _, a := range numbers {
		fmt.Printf("%5d", a)
		for _, b := range numbers {
			fmt.Printf("%5d", a * b)
		}
		fmt.Println()
	}
}

