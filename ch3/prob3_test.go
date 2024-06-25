package ch3_test

import (
	"testing"

	"github.com/dishbreak/go-cracking-coding/ch3"
	"github.com/stretchr/testify/assert"
)

func TestStackOfStacks(t *testing.T) {
	s := &ch3.StackOfStacks{}

	for _, v := range []int{3, 4, 5, 5, 4, 3, 2, 3, 4, 7, 9, 1} {
		s.Push(v)
	}

	for _, v := range []int{1, 9, 7, 4, 3, 2, 3, 4, 5, 5, 4, 3} {
		assert.Equal(t, v, s.Peek())
		assert.Equal(t, v, s.Pop())
	}
}
