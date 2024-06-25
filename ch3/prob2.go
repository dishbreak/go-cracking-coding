package ch3

import "github.com/dishbreak/go-cracking-coding/lib/stack"

type minStackFrame struct {
	value int
	next  *minStackFrame
	min   *minStackFrame
}

type MinStack struct {
	head *minStackFrame
}

func (m *MinStack) Push(i int) {
	f := &minStackFrame{
		value: i,
	}
	f.min = f
	if m.head == nil {
		m.head = f
		return
	}

	if m.head.min.value <= f.value {
		f.min = m.head.min
	}
	f.next = m.head
	m.head = f
}

func (m *MinStack) Pop() (int, bool) {
	if m.head == nil {
		return 0, false
	}

	value := m.head.value
	m.head = m.head.next

	return value, true
}

func (m *MinStack) Min() (int, bool) {
	if m.head == nil {
		return 0, false
	}

	return m.head.min.value, true
}

type MinStackV2 struct {
	min  *stack.IntStack
	vals *stack.IntStack
}

func NewMinStackV2() *MinStackV2 {
	s := &MinStackV2{
		min:  &stack.IntStack{},
		vals: &stack.IntStack{},
	}

	return s
}

func (m *MinStackV2) Push(i int) {
	m.vals.Push(i)
	if m.min.Empty() || m.min.Peek() > i {
		m.min.Push(i)
	}
}

func (m *MinStackV2) Pop() int {
	v := m.vals.Pop()
	if m.min.Peek() == v {
		m.min.Pop()
	}
	return v
}

func (m *MinStackV2) Min() int {
	return m.min.Peek()
}
