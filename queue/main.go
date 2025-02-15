package main

import (
	"fmt"

	"github.com/felipeazsantos/dsa-go/queue/queueimpl"
)

func main() {
	myQueue := queueimpl.NewQueue()
	myQueue.Queue(58)
	myQueue.Queue(175)
	myQueue.Queue(1578)
	myQueue.Queue(54)
	myQueue.Queue(1)
	myQueue.Queue(79)
	myQueue.Queue(984)
	myQueue.Queue(99)

	fmt.Println("Welcome to BankApp")
	fmt.Println("The customer always comes first")
	fmt.Println("Treatment------------------------")
	for !myQueue.IsEmpty() {
		fmt.Printf("Password to be answered at the counter: %d\n", myQueue.Dequeue())
	}
}
