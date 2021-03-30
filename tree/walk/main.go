package main

import (
	"github.com/pyankovzhe/go-tour/tree/btree"
)

func Walk(t *btree.Tree, ch chan int) {
	if t == nil {
		close(ch)
		return
	}
}

func main() {
	t := btree.NewTree()
	s := []int{10, -1, -10, 5, 6, 8, 4 - 2}
	for v := range s {
		t.Insert(v)
	}

	ch := make(chan int)
	Walk(t, ch)
}
