package main

import "fmt"

func doPanic() {
	panic("exception!")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered; %v; %T", r, r)
		}
	}()

	doPanic()
	fmt.Println("After")
}
