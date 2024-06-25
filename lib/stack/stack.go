package stack

import "fmt"

type intStackFrame struct {
	value int
	next  *intStackFrame
}

type IntStack struct {
	head *intStackFrame
	len  int
}

func (s *IntStack) Push(i int) {
	f := &intStackFrame{
		value: i,
		next:  s.head,
	}

	s.len++
	s.head = f
}

func (s *IntStack) Pop() int {
	if s.head == nil {
		return 0
	}

	val := s.head.value
	s.head = s.head.next
	s.len--
	return val
}

func (s *IntStack) Peek() int {
	if s.Empty() {
		return 0
	}

	return s.head.value
}

func (s *IntStack) Len() int {
	return s.len
}

func (s *IntStack) Empty() bool {
	return s.head == nil
}

func (s *IntStack) String() string {
	var v []int

	for c := s.head; c != nil; c = c.next {
		v = append(v, c.value)
	}

	return fmt.Sprint(v)
}

type stackFrame struct {
	data interface{}
	next *stackFrame
}

type Stack struct {
	count int
	head  *stackFrame
}

func (s *Stack) Empty() bool {
	return s.head == nil
}

func (s *Stack) Len() int {
	return s.count
}

func (s *Stack) Push(d interface{}) {
	f := &stackFrame{
		data: d,
		next: s.head,
	}
	s.count++
	s.head = f
}

func (s *Stack) Peek() interface{} {
	if s.Empty() {
		return nil
	}

	return s.head.data
}

func (s *Stack) Pop() interface{} {
	if s.Empty() {
		return nil
	}

	d := s.head.data
	s.count--
	s.head = s.head.next
	return d
}
