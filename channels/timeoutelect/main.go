package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "response 1"
	}()

	select {
	case msg := <-ch1:
		fmt.Println("res: ", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout after 1 second")
	}

	ch2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "response 2"
	}()

	select {
	case msg := <-ch2:
		fmt.Println("res 2:", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout after 3 seconds")
	}
}
