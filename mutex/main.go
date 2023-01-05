package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.counters[key]++
}

func main() {
	c := Container{
		// zero value of a mutex is usable as-is, so no initialization is required here.
		counters: map[string]int{"a": 0, "b": 0},
	}

	wg := &sync.WaitGroup{}

	doIncrement := func(key string, n int) {
		for i := 0; i < n; i++ {
			c.inc(key)
		}
		wg.Done()
	}

	wg.Add(4)
	go doIncrement("a", 100)
	go doIncrement("a", 1050)
	go doIncrement("b", 500)
	go doIncrement("b", 550)

	wg.Wait()
	fmt.Println(c.counters)
}
