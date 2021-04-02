package main

import (
	"fmt"
	"time"
)

func worker(w int, jobs <-chan int, results chan<- int) {
	fmt.Printf("Starting worker %d\n", w)

	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", w, j)
		time.Sleep(60 * time.Millisecond)
		results <- j * 2
		fmt.Printf("Worker %d finished processing job %d\n", w, j)
	}
}

func main() {
	const jobNum = 5
	jobs := make(chan int, jobNum)
	results := make(chan int, jobNum)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for i := 1; i <= jobNum; i++ {
		jobs <- i
	}
	close(jobs)

	for r := 1; r <= jobNum; r++ {
		fmt.Printf("Result: %d\n", <-results)
	}

	fmt.Println("Success")
}
