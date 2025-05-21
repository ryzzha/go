package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() { ch1 <- "hello from channel_1" }()
	// go func() { ch2 <- "hello from channel_2" }()

	// we use select to wait for multiple channel operations
	select {
	case msg1 := <- ch1:
		fmt.Println(msg1)
	case msg2 := <- ch2:
		fmt.Println(msg2)
	}

	// if we try to get data from channel which is not closed, it will block the program
	doDeadlock := <- ch2
	fmt.Println(doDeadlock)

	time.Sleep(2 * time.Second) // Wait for goroutines to finish
}