package main

import "fmt"

type IntStack struct {
	stack []int
}

func (i *IntStack) Push(x int) {
	i.stack = append(i.stack, x)
}

func (i *IntStack) Pop() int {
	if len(i.stack) == 0 {
		return 0
	}

	last := i.stack[len(i.stack)-1]
	i.stack = i.stack[:len(i.stack)-1]
	return last
}

func main() {
	st := IntStack{}
	st.Push(10)
	st.Push(20)
	st.Push(30)
	fmt.Printf("shoud eq 30 %v \n", st.Pop())
	fmt.Printf("shoud eq 20 %v \n", st.Pop())
	fmt.Printf("shoud eq 10 %v \n", st.Pop())
	fmt.Printf("shoud eq 0 %v \n", st.Pop())
}
