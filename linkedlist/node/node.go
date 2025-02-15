package node

type MyNode struct {
	element int
	next    *MyNode
}

func NewNode() *MyNode {
	return &MyNode{
		element: 0,
		next:    &MyNode{},
	}
}

func (n *MyNode) SetElement(element int) {
	n.element = element
}

func (n *MyNode) GetElement() int {
	return n.element
}

func (n *MyNode) SetNext(next *MyNode) {
	n.next = next
}

func (n *MyNode) GetNext() *MyNode {
	return n.next
}
