package main

import (
	"flag"
	"fmt"

	"github.com/felipeazsantos/dsa-go/stack/stackimpl"
)

func main() {
	var number int
	flag.IntVar(&number, "entry number", 172, "decimal number to convert to binary")
	flag.Parse()

	stack := stackimpl.NewStack()

	for number != 0 {
		rest := number % 2
		stack.Push(rest)
		number /= 2
	}

	for !stack.IsEmpty() {
		fmt.Print(stack.Pop())
	}

}
