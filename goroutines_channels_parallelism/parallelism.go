package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	fmt.Printf("Worker %d started\n", id)

	for j := range jobs {
		fmt.Printf("Worker %d got job %d\n", id, j)
		time.Sleep(500 * time.Millisecond) // імітація обробки задачі

		results <- j * 2
		fmt.Printf("Worker %d done with job %d\n", id, j)
	}

	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	// if we set buffer size to 0, it program will start workers earlear and wait for jobs anotheway firstly we get jobs in channel buffer and then we start workers
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	time.Sleep(100 * time.Millisecond) // Give workers time to start

	for j := 1; j <= 5; j++ {
		fmt.Printf("Sending job %d\n", j)
		jobs <- j
	}
	close(jobs)


	for a := 1; a <= 5; a++ {
		fmt.Println(<-results)
	}
}