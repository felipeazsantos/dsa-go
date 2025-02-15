package dynamicstackimpl

import (
	"fmt"

	"github.com/felipeazsantos/dsa-go/linkedlist/node"
)

type dynamicStack struct {
	topNode *node.MyNode
}

func NewDynamicStack() *dynamicStack {
	return &dynamicStack{
		topNode: nil,
	}
}

func (d *dynamicStack) Push(element int) {
	newNode := node.NewNode()
	newNode.SetElement(element)

	if d.IsEmpty() {
		d.topNode = newNode
	} else {
		auxNode := d.topNode
		d.topNode = newNode
		d.topNode.SetNext(auxNode)
	}
}

func (d *dynamicStack) Pop() int {
	if d.IsEmpty() {
		fmt.Println("cannot pop, stack is empty")
		return -1
	}

	element := d.topNode.GetElement()
	d.topNode = d.topNode.GetNext()
	return element
}

func (d *dynamicStack) IsEmpty() bool {
	return d.topNode == nil
}
