package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var cond *sync.Cond
var tasks []func()

func worker(w int) {
	var task func()

	fmt.Printf("Worker %d started\n", w)

	mu.Lock()
	for len(tasks) == 0 {
		fmt.Printf("Worker %d waiting\n", w)
		cond.Wait()
	}
	task, tasks = tasks[0], tasks[1:]
	mu.Unlock()

	fmt.Printf("Worker %d running task:\n", w)
	task()
}

func produce(task func()) {
	mu.Lock()
	fmt.Println("Add task")
	tasks = append(tasks, task)
	mu.Unlock()

	cond.Broadcast()
}

func main() {
	cond = sync.NewCond(&mu)

	for i := 1; i <= 5; i++ {
		go worker(i)
	}

	time.Sleep(time.Second)

	produce(func() { fmt.Println("1") })
	produce(func() { fmt.Println("2") })
	produce(func() { fmt.Println("3") })

	time.Sleep(time.Second)
}
