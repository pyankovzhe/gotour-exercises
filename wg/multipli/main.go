package main

import (
	"fmt"
	"sync"
)

func printTable(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d\n", i, n, n*i)
	}

}

func main() {
	wg := &sync.WaitGroup{}

	for number := 2; number <= 12; number++ {
		wg.Add(1)
		go printTable(number, wg)
	}

	wg.Wait()
}
