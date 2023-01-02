package main

import "fmt"

func sendValues(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Sending to channel value: %v \n", i)
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	// defer close(ch)
	go sendValues(ch)

	// Example with open flag:
	// for i := 0; i < 6; i++ {
	// 	value, open := <-ch
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Printf("Reading from channel: %v \n", value)
	// }

	// Example with range:
	for value := range ch {
		fmt.Printf("Reading from channel: %v \n", value)
	}
}
