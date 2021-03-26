package main

import "fmt"

type List interface {
	Len() int
	First() *ListElement
	Last() *ListElement
	PushFront(v interface{}) *ListElement
	Values() []interface{}
	// PushBack(v interface{}) *ListElement
	// Remove(i *ListElement)
	// MoveToFront(i *ListElement)
}

type list struct {
	size  int
	first *ListElement
	last  *ListElement
}

type ListElement struct {
	Value interface{}
	Next  *ListElement
	Prev  *ListElement
}

func NewList() List {
	return &list{}
}

func (l list) Len() int {
	return l.size
}

func (l list) First() *ListElement {
	return l.first
}

func (l list) Last() *ListElement {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListElement {
	el := &ListElement{Value: v, Next: l.first}

	if l.size == 0 {
		l.first = el
		l.last = el
	} else {
		l.first.Prev = el
		l.first = el
	}
	l.size++

	return el
}

func (l list) Values() []interface{} {
	values := make([]interface{}, l.size, l.size)
	for i, el := 0, l.first; el != nil; i, el = i+1, el.Next {
		values[i] = el.Value
	}

	return values
}

func main() {
	list := NewList()
	fmt.Println(list)
	fmt.Println(list.Len())

	elements := []struct {
		name  string
		Value string
	}{
		{name: "el1", Value: "foo"},
		{name: "el2", Value: "boo"},
	}

	for _, el := range elements {
		newEl := list.PushFront(el)
		fmt.Printf("new pushed el: %v\n", newEl.Value)
	}

	fmt.Printf("list: %v\n", list)
	fmt.Println(list.Len())

	fmt.Printf("list values: %v\n", list.Values())
}
