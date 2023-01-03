package main

import (
	"fmt"
	"sync"
	"time"
)

func Work(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d strting\n", i)
	time.Sleep(10 * time.Microsecond)
	fmt.Printf("Worker %d finished\n", i)
}

func main() {
	wg := &sync.WaitGroup{}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		Work(i, wg)
	}

	wg.Wait()
	fmt.Println("Success")
}
