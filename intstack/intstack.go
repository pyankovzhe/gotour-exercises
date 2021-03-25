package intstack

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
