package rtda

type Stack struct {
	maxSize int
	size    int
	_top    *Frame
}

func newStack(maxSize int) *Stack {
	return &Stack{maxSize: maxSize}
}

func (s *Stack) push(frame *Frame) {
	if s.size > s.maxSize {
		panic("java.lang.StackoverflowError")
	}
	if s._top != nil {
		frame.lower = s._top
	}
	s._top = frame
	s.size++
}

func (s *Stack) pop() *Frame {
	if s._top == nil {
		panic("stack is empty")
	}
	top := s._top
	s._top = top.lower
	top.lower = nil
	s.size--
	return top
}

func (s *Stack) top() *Frame {
	if s._top == nil {
		panic("stack is empty")
	}
	return s._top
}
