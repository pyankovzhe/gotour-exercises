package main

import (
	"fmt"
)

func read(ch <-chan int) {
	res := <-ch
	fmt.Println(res)
}

func main() {
	ch := make(chan int)
	go read(ch)
	ch <- 5
}
