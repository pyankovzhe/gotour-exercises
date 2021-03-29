package main

import "fmt"

func sum(s []int, ch chan int) {
	var sum int

	for _, v := range s {
		sum += v
	}

	ch <- sum
}
func main() {
	s := []int{1, -5, 3, 4, 9}

	ch := make(chan int)
	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)

	x, y := <-ch, <-ch

	fmt.Println(x, y, x+y)
}
