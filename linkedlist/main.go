package main

import (
	"fmt"

	"github.com/felipeazsantos/dsa-go/linkedlist/linkedlistimpl"
)

func main() {
	ll := linkedlistimpl.NewLinkedList()
	ll.Insert(54)
	ll.Insert(875)
	ll.Insert(1578)
	ll.Insert(2)
	ll.Insert(158)
	ll.Insert(65)
	ll.Insert(879)
	ll.Insert(1455)

	ll.List()
	fmt.Println("------------------------------")

	for !ll.IsEmpty() {
		fmt.Printf("removing element from list: %d\n", ll.Remove())
	}
}
