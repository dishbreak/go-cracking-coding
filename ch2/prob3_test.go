package ch2_test

import (
	"testing"

	"github.com/dishbreak/go-cracking-coding/ch2"
	"github.com/dishbreak/go-cracking-coding/lib/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestDeleteMiddle(t *testing.T) {
	head := linkedlist.FromSlice([]int{1, 3, 3, 4, 5, 3, 4})
	n := linkedlist.Index(head, 3)
	assert.Equal(t, 4, n.Value)
	ch2.DeleteMiddle(n)
	assert.Equal(t, []int{1, 3, 3, 5, 3, 4}, linkedlist.ToSlice(head))
}
