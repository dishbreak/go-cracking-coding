package ch2_test

import (
	"fmt"
	"testing"

	"github.com/dishbreak/go-cracking-coding/ch2"
	"github.com/dishbreak/go-cracking-coding/lib/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestAddReversed(t *testing.T) {
	type testCase struct {
		a, b, result []int
	}

	testCases := []testCase{
		{
			[]int{7, 1, 6}, []int{5, 9, 2}, []int{2, 1, 9},
		},
		{
			[]int{7, 1}, []int{5, 9, 2, 1, 5}, []int{2, 1, 3, 1, 5}, // 51313
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			a, b := linkedlist.FromSlice(tc.a), linkedlist.FromSlice(tc.b)
			assert.Equal(t, tc.result, linkedlist.ToSlice(ch2.SumReversed(a, b)))
		})
	}
}

func TestAddForwards(t *testing.T) {
	type testCase struct {
		a, b, result []int
	}

	testCases := []testCase{
		{
			[]int{7, 1, 6}, []int{5, 9, 2}, []int{1, 3, 0, 8},
		},
		{
			[]int{7, 1}, []int{5, 9, 2, 1, 5}, []int{5, 9, 2, 8, 6},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			a, b := linkedlist.FromSlice(tc.a), linkedlist.FromSlice(tc.b)
			assert.Equal(t, tc.result, linkedlist.ToSlice(ch2.SumForwards(a, b)))
		})
	}
}
