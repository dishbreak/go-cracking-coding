package ch3

import "github.com/dishbreak/go-cracking-coding/lib/stack"

type MyQueue struct {
	in, out *stack.IntStack
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		in:  &stack.IntStack{},
		out: &stack.IntStack{},
	}
}

func (m *MyQueue) Enqueue(i int) {
	m.in.Push(i)
}

func (m *MyQueue) Dequeue() int {
	if !m.out.Empty() {
		return m.out.Pop()
	}

	if m.in.Empty() {
		return 0
	}

	for !m.in.Empty() {
		m.out.Push(m.in.Pop())
	}

	return m.out.Pop()
}
