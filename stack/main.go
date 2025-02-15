package main

import (
	"flag"
	"fmt"

	"github.com/felipeazsantos/dsa-go/stack/dynamicstack/dynamicstackimpl"
	"github.com/felipeazsantos/dsa-go/stack/stackimpl"
)

type IStack interface {
	Push(element int)
	Pop() int
	IsEmpty() bool
}

func main() {
	var number, whichStack int
	var stack IStack
	flag.IntVar(&number, "number", 172, "decimal number to convert to binary")
	flag.IntVar(&whichStack, "stack", 0, "entry which stack you want to use (0 - Static Stack, 1 - Dynamic Stack)")
	flag.Parse()

	if whichStack == 0 {
		stack = stackimpl.NewStack()
	} else {
		stack = dynamicstackimpl.NewDynamicStack()
	}

	for number != 0 {
		rest := number % 2
		stack.Push(rest)
		number /= 2
	}

	for !stack.IsEmpty() {
		fmt.Print(stack.Pop())
	}

}
