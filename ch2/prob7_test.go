package ch2_test

import (
	"testing"

	"github.com/dishbreak/go-cracking-coding/ch2"
	"github.com/dishbreak/go-cracking-coding/lib/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	t.Run("detects interestion between two lists of same length", func(t *testing.T) {
		a := linkedlist.FromSlice([]int{1, 2, 3, 5, 6, 8})
		intersection := linkedlist.Index(a, 3)
		b := linkedlist.FromSlice([]int{6, 7, 8})
		linkedlist.Index(b, 2).Next = intersection
		assert.True(t, ch2.Intersection(a, b) == intersection)
	})

	t.Run("detects interection between two lists of differing lengths", func(t *testing.T) {
		a := linkedlist.FromSlice([]int{2, 2, 5, 1, 2, 3, 5, 6, 8})
		intersection := linkedlist.Index(a, 6)
		b := linkedlist.FromSlice([]int{6, 7, 8})
		linkedlist.Index(b, 2).Next = intersection
		assert.True(t, ch2.Intersection(a, b) == intersection)
		assert.True(t, ch2.Intersection(b, a) == intersection)
	})

	t.Run("sees no intersction in distinct lists with common overlap.", func(t *testing.T) {
		a := linkedlist.FromSlice([]int{1, 2, 3, 5, 6, 8})
		b := linkedlist.FromSlice([]int{6, 7, 8, 5, 6, 8})

		assert.Nil(t, ch2.Intersection(a, b))
	})
}
