package main

import (
	"fmt"
	"time"
)

func sayHello() {
	time.Sleep(2 * time.Second) // Simulate some work
	fmt.Println("Hello from goroutine")
}

func main() {
	go sayHello() // Start a goroutine
	go printNumbers() // Start a goroutine

	fmt.Println("Main finished")
	time.Sleep(10 * time.Second) // Wait for goroutine to finish
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(2 * time.Second)
	}
}