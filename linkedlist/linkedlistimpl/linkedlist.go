package linkedlistimpl

import (
	"fmt"

	"github.com/felipeazsantos/dsa-go/linkedlist/node"
)

type linkedList struct {
	beginNode *node.MyNode
}

func NewLinkedList() *linkedList {
	return &linkedList{
		beginNode: nil,
	}
}

func (ll *linkedList) Insert(element int) {
	newNode := node.NewNode()
	newNode.SetElement(element)
	newNode.SetNext(nil)

	if ll.IsEmpty() {
		ll.beginNode = newNode
	} else {
		auxNode := ll.beginNode
		for auxNode.GetNext() != nil {
			auxNode = auxNode.GetNext()
		}
		auxNode.SetNext(newNode)
	}
}

func (ll *linkedList) Remove() int {
	if ll.IsEmpty() {
		fmt.Println("cannot remove, list is empty ")
	} else {
		element := ll.beginNode.GetElement()
		ll.beginNode = ll.beginNode.GetNext()
		return element
	}
	return -1
}

func (ll *linkedList) List() {
	if ll.IsEmpty() {
		fmt.Println("cannot list, list is empty")
	} else {
		auxNode := ll.beginNode
		for auxNode != nil {
			fmt.Printf("listing element: %d\n", auxNode.GetElement())
			auxNode = auxNode.GetNext()
		}
	}
}

func (ll *linkedList) IsEmpty() bool {
	return ll.beginNode == nil
}
