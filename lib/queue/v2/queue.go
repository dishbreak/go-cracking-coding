package queue

import "github.com/dishbreak/go-cracking-coding/lib/stack/v2"

type Queue[T any] struct {
	in, out *stack.Stack[T]
}

func (q *Queue[T]) Enqueue(v T) {
	q.in.Push(v)
}

func (q *Queue[T]) Dequeue() T {
	if !q.out.Empty() {
		return q.out.Pop()
	}

	if !q.in.Empty() {
		for q.in.Len() > 1 {
			q.out.Push(q.in.Pop())
		}
		return q.in.Pop()
	}

	var zero T
	return zero
}

func (q *Queue[T]) Peek() T {
	if !q.out.Empty() {
		return q.out.Peek()
	}

	if !q.in.Empty() {
		for !q.in.Empty() {
			q.out.Push(q.in.Pop())
		}
		return q.out.Peek()
	}

	var zero T
	return zero
}
