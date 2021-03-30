// Exercise: Equivalent Binary Trees
// https://tour.golang.org/concurrency/7

package main

import (
	"fmt"

	"github.com/pyankovzhe/go-tour/tree/btree"
)

func WalkTree(t *btree.Tree, ch chan int) {
	if t == nil {
		close(ch)
		return
	}

	walk(t.Root, ch)
	close(ch)
}

func walk(node *btree.Node, ch chan int) {
	if node == nil {
		return
	}
	ch <- node.Value

	walk(node.Left, ch)
	walk(node.Right, ch)
}

func main() {
	t := btree.NewTree()
	s := []int{10, -1, -10, 5, 6, 8, 4 - 2}
	for _, v := range s {
		t.Insert(v)
	}

	ch := make(chan int, len(s))
	go WalkTree(t, ch)
	for v := range ch {
		fmt.Println(v)
	}
}
