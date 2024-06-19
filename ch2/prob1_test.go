package ch2_test

import (
	"testing"

	"github.com/dishbreak/go-cracking-coding/ch2"
	"github.com/dishbreak/go-cracking-coding/lib/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestDedupe(t *testing.T) {
	head := linkedlist.FromSlice([]int{1, 3, 1, 2, 3, 3, 3, 5, 1, 4, 5})

	result := ch2.Dedupe(head)

	assert.Equal(t, []int{1, 3, 2, 5, 4}, linkedlist.ToSlice(result))
}

func TestDedupeNoSpace(t *testing.T) {
	head := linkedlist.FromSlice([]int{1, 3, 1, 2, 3, 3, 3, 5, 1, 4, 5})

	result := ch2.DedupeNoSpace(head)

	assert.Equal(t, []int{1, 3, 2, 5, 4}, linkedlist.ToSlice(result))
}
