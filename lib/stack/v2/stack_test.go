package stack_test

import (
	"testing"

	"github.com/dishbreak/go-cracking-coding/lib/stack/v2"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	// custom type for some generics testing
	type animal struct {
		age  int
		name string
	}
	t.Run("Base int stack returns zero values when empty for peek and pop", func(t *testing.T) {
		s := &stack.Stack[int]{}
		assert.Zero(t, s.Peek())
		assert.Zero(t, s.Pop())
		assert.Zero(t, s.Len())
		assert.True(t, s.Empty())
	})
	t.Run("Zero value contract holds on empty stack even when using a struct type", func(t *testing.T) {

		s := &stack.Stack[animal]{}
		assert.Zero(t, s.Peek())
		assert.Zero(t, s.Pop())
		assert.Zero(t, s.Len())
		assert.True(t, s.Empty())

	})
	t.Run("items pushed to the stack appear on peek", func(t *testing.T) {
		s := &stack.Stack[animal]{}
		s.Push(animal{age: 4, name: "leo"})
		assert.Equal(t, animal{age: 4, name: "leo"}, s.Peek())
		assert.False(t, s.Empty()) // peek keeps the item on the stack
	})

	t.Run("stack obeys LIFO semantics", func(t *testing.T) {
		s := &stack.Stack[int]{}
		s.Push(7)
		s.Push(4)
		s.Push(3)

		assert.Equal(t, 3, s.Pop())
		assert.Equal(t, 4, s.Pop())
		assert.Equal(t, 7, s.Pop())
	})
}
