package ch3

import "github.com/dishbreak/go-cracking-coding/lib/stack"

type stackStackFrame struct {
	val  *stack.IntStack
	next *stackStackFrame
}

const (
	StackLimit = 5
)

type StackOfStacks struct {
	head *stackStackFrame
}

func (s *StackOfStacks) Empty() bool {
	return s.head == nil
}

func (s *StackOfStacks) topStack() *stack.IntStack {
	if s.head == nil {
		return nil
	}
	return s.head.val
}

func (s *StackOfStacks) Pop() int {
	if s.Empty() {
		return 0
	}

	v := s.topStack().Pop()
	if s.topStack().Empty() {
		s.head = s.head.next
	}

	return v
}

func (s *StackOfStacks) Peek() int {
	if s.Empty() {
		return 0
	}

	return s.topStack().Peek()
}

func (s *StackOfStacks) Push(i int) {
	if s.Empty() || s.topStack().Len() == StackLimit {
		f := &stackStackFrame{
			val:  &stack.IntStack{},
			next: s.head,
		}
		s.head = f
	}
	s.topStack().Push(i)
}
