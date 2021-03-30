package btree

type Tree struct {
	Root *Node
}

type Node struct {
	Value int
	Right *Node
	Left  *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(data int) *Tree {
	if t.Root == nil {
		t.Root = &Node{
			Value: data,
			Left:  nil,
			Right: nil,
		}
	} else {
		t.Root.insert(data)
	}

	return t
}

func (node *Node) insert(data int) {
	if node == nil {
		return
	}

	if data <= node.Value {
		if node.Left == nil {
			node.Left = &Node{Value: data, Right: nil, Left: nil}
		} else {
			node.Left.insert(data)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Value: data, Right: nil, Left: nil}
		} else {
			node.Right.insert(data)
		}
	}
}
