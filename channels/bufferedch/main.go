package main

import (
	"fmt"
	"time"
)

// basic buffered channel example
func bufferedch() {
	ch := make(chan string, 2)

	ch <- "one"
	ch <- "two"
	// ch <- "two" - deadlock
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- "three"
	fmt.Println(<-ch)
}

// example how to use buffered channel for syncrinization
func sync() {
	doneCh := make(chan bool, 1)
	go worker(doneCh)
	<-doneCh
	fmt.Println("Sync finished.")
}

func worker(doneCh chan bool) {
	fmt.Println("Worker started.")
	fmt.Println("Doing some work...")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Worker finishing...")
	doneCh <- true
	fmt.Println("Worker finished")
}

// channel direction example
func pingpong() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "My msg")
	pong(pings, pongs)
	fmt.Printf("Pong: %v \n", <-pongs)
}

func ping(pings chan<- string, msg string) {
	fmt.Printf("Ping: %v \n", msg)
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	// bufferedch()
	// sync()
	pingpong()
}
