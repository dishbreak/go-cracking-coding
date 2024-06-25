package ch3_test

import (
	"testing"

	"github.com/dishbreak/go-cracking-coding/ch3"
	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
	t.Run("Empty stack behaves correctly", func(t *testing.T) {
		s := ch3.MinStack{}
		v, ok := s.Pop()
		assert.Equal(t, 0, v)
		assert.Equal(t, false, ok)

		v, ok = s.Min()
		assert.Equal(t, 0, v)
		assert.Equal(t, false, ok)
	})

	t.Run("Popping items off a stack follows LIFO order", func(t *testing.T) {
		s := ch3.MinStack{}
		s.Push(5)
		s.Push(7)
		s.Push(3)
		s.Push(4)

		results := []int{4, 3, 7, 5}
		for _, expected := range results {
			v, ok := s.Pop()
			assert.True(t, ok)
			assert.Equal(t, expected, v)
		}

		v, ok := s.Pop()
		assert.False(t, ok)
		assert.Equal(t, 0, v)
	})

	t.Run("Min updates correctly as items pop", func(t *testing.T) {
		s := ch3.MinStack{}
		s.Push(5)
		s.Push(7)
		s.Push(3)
		s.Push(9)
		s.Push(4)

		v, ok := s.Min()
		assert.True(t, ok)
		assert.Equal(t, 3, v)

		v, ok = s.Pop()
		assert.True(t, ok)
		assert.Equal(t, 4, v)

		v, ok = s.Min()
		assert.True(t, ok)
		assert.Equal(t, 3, v)

		v, ok = s.Pop()
		assert.True(t, ok)
		assert.Equal(t, 9, v)

		v, ok = s.Min()
		assert.True(t, ok)
		assert.Equal(t, 3, v)

		v, ok = s.Pop()
		assert.True(t, ok)
		assert.Equal(t, 3, v)

		v, ok = s.Min()
		assert.True(t, ok)
		assert.Equal(t, 5, v)

		v, ok = s.Pop()
		assert.True(t, ok)
		assert.Equal(t, 7, v)

		v, ok = s.Min()
		assert.True(t, ok)
		assert.Equal(t, 5, v)

		v, ok = s.Pop()
		assert.True(t, ok)
		assert.Equal(t, 5, v)

		v, ok = s.Pop()
		assert.False(t, ok)
		assert.Equal(t, 0, v)

	})
}

func TestMinStackV2(t *testing.T) {
	t.Run("an empty stack returns zero values on Pop and Min.", func(t *testing.T) {
		m := ch3.NewMinStackV2()
		assert.Equal(t, 0, m.Min())
		assert.Equal(t, 0, m.Pop())
	})

	t.Run("Stack returns the correct min throughout", func(t *testing.T) {
		m := ch3.NewMinStackV2()
		m.Push(5)
		m.Push(7)
		m.Push(3)
		m.Push(9)
		m.Push(4)

		assert.Equal(t, 3, m.Min())
		assert.Equal(t, 4, m.Pop())
		assert.Equal(t, 3, m.Min())
		assert.Equal(t, 9, m.Pop())
		assert.Equal(t, 3, m.Min())
		assert.Equal(t, 3, m.Pop())
		assert.Equal(t, 5, m.Min())
		assert.Equal(t, 7, m.Pop())
		assert.Equal(t, 5, m.Min())
		assert.Equal(t, 5, m.Pop())
		assert.Equal(t, 0, m.Min())
		assert.Equal(t, 0, m.Pop())
	})
}
