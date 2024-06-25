package ch3_test

import (
	"testing"

	"github.com/dishbreak/go-cracking-coding/ch3"
	"github.com/stretchr/testify/assert"
)

func TestSortStack(t *testing.T) {

	s := ch3.NewSortStack()
	s.Push(5)
	s.Push(7)
	s.Push(3)
	s.Push(9)
	s.Push(4)

	assert.Equal(t, "[9 7 5 4 3]", s.String())
}
