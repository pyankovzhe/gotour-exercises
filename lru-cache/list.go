package lrucache

type List interface {
	Len() int
	First() *ListElement
	Last() *ListElement
	PushFront(v interface{}) *ListElement
	PushBack(v interface{}) *ListElement
	Remove(el *ListElement)
	MoveToFront(i *ListElement)
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

func (l *list) PushBack(v interface{}) *ListElement {
	el := &ListElement{Value: v, Prev: l.last}

	if l.size == 0 {
		l.first = el
		l.last = el
	} else {
		l.last.Next = el
		l.last = el
	}
	l.size++

	return el
}

func (l *list) MoveToFront(el *ListElement) {
	if el == l.first {
		return
	} else if el == l.last {
		l.last = el.Prev
		el.Prev.Next = nil
	} else {
		el.Prev.Next, el.Next.Prev = el.Next, el.Prev
	}

	l.first.Prev = el
	l.first = el
	el.Prev = nil
}

func (l *list) Remove(el *ListElement) {
	if el == l.last {
		l.last = el.Prev
		el.Prev.Next = nil
	} else if el == l.first {
		l.first = el.Next
		el.Next.Prev = nil
	} else {
		el.Prev.Next, el.Next.Prev = el.Next, el.Prev
	}

	l.size--
}
