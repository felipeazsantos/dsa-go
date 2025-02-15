package stackimpl

type stack struct {
	values [10]int
	top    int
}

func NewStack() *stack {
	return &stack{
		values: [10]int{},
		top:    -1,
	}
}

func (s *stack) Push(element int) {
	s.top++
	s.values[s.top] = element
}

func (s *stack) IsEmpty() bool {
	return s.top == -1
}

func (s *stack) Pop() int {
	element := s.values[s.top]
	s.top--
	return element
}
