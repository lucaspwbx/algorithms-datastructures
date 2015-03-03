package stack

type Stack struct {
	stack  []interface{}
	length int
}

func NewStack() *Stack {
	s := &Stack{}
	s.stack = make([]interface{}, 0)
	s.length = 0

	return s
}

func (s *Stack) Length() int {
	return s.length
}

func (s *Stack) IsEmpty() bool {
	return s.length == 0
}

func (s *Stack) Push(el interface{}) {
	s.stack = append(s.stack, el)
	s.length++
}

func (s *Stack) Pop() interface{} {
	index := s.length - 1
	el := s.stack[index]
	s.stack = s.stack[:index]
	s.length--

	return el
}

func (s *Stack) Peek() interface{} {
	return s.stack[s.length-1]
}
