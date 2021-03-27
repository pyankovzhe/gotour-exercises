package lrucache

type List interface {
	Len() int
	First() *ListElement
	Last() *ListElement
	PushFront(v interface{}) *ListElement
	PushBack(v interface{}) *ListElement
	Values() []interface{}
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

func (l list) Values() []interface{} {
	values := make([]interface{}, l.size, l.size)
	for i, el := 0, l.first; el != nil; i, el = i+1, el.Next {
		values[i] = el.Value
	}

	return values
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

// func main() {
// 	list := NewList()
// 	fmt.Println(list)
// 	fmt.Println(list.Len())

// 	elements := []struct {
// 		name  string
// 		Value string
// 	}{
// 		{name: "el1", Value: "foo"},
// 		{name: "el2", Value: "boo"},
// 		{name: "el3", Value: "bar"},
// 		{name: "el4", Value: "bar4"},
// 	}

// 	for _, el := range elements[:2] {
// 		newEl := list.PushFront(el)
// 		fmt.Printf("new pushed el: %v\n", newEl.Value)
// 	}

// 	penultEl := list.PushBack(elements[len(elements)-2])
// 	fmt.Printf("new pushed last el: %v\n", penultEl.Value)
// 	lastEl := list.PushBack(elements[len(elements)-1])
// 	fmt.Printf("new pushed last el: %v\n", lastEl)
// 	fmt.Printf("list: %v\n", list)

// 	list.Remove(penultEl)
// 	fmt.Printf("list: %v\n", list)
// 	fmt.Println(list.Len())
// 	fmt.Printf("last el: %v\n", lastEl.Prev != penultEl)
// 	list.Remove(lastEl)
// 	fmt.Println(list.Len())

// 	fmt.Printf("list values: %v\n", list.Values())
// }
