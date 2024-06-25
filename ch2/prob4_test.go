package ch2_test

import (
	"testing"

	"github.com/dishbreak/go-cracking-coding/ch2"
	"github.com/dishbreak/go-cracking-coding/lib/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	head := linkedlist.FromSlice([]int{3, 5, 8, 5, 10, 2, 1})
	assert.Equal(t, []int{1, 2, 3, 3, 5, 8, 5, 10}, linkedlist.ToSlice(ch2.Partition(head, 5)))
}
