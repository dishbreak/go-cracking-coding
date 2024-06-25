package stack

type stackFrame[T any] struct {
	data T
	next *stackFrame[T]
}

type Stack[T any] struct {
	count int
	head  *stackFrame[T]
}

func (s *Stack[T]) Peek() T {
	if s.head == nil {
		var zero T
		return zero
	}
	return s.head.data
}

func (s *Stack[T]) Pop() T {
	if s.head == nil {
		var zero T
		return zero
	}
	v := s.head.data
	s.count--
	s.head = s.head.next
	return v
}

func (s *Stack[T]) Push(v T) {
	f := &stackFrame[T]{
		data: v,
		next: s.head,
	}
	s.count++
	s.head = f
}

func (s *Stack[T]) Len() int {
	return s.count
}

func (s *Stack[T]) Empty() bool {
	return s.head == nil
}
