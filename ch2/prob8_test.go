package ch2_test

import (
	"testing"

	"github.com/dishbreak/go-cracking-coding/ch2"
	"github.com/dishbreak/go-cracking-coding/lib/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestLooopedPoint(t *testing.T) {
	t.Run("finds loop when one exists", func(t *testing.T) {
		a := linkedlist.FromSlice([]int{1, 2, 5, 6, 4, 3})
		looped := linkedlist.Index(a, 3)
		tail := linkedlist.Tail(a)
		tail.Next = looped

		assert.True(t, looped == ch2.LoopedPoint(a))
	})

	t.Run("finds no loop when there is none", func(t *testing.T) {
		a := linkedlist.FromSlice([]int{1, 2, 5, 6, 4, 3})
		assert.Nil(t, ch2.LoopedPoint(a))
	})
}
