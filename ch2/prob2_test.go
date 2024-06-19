package ch2_test

import (
	"fmt"
	"testing"

	"github.com/dishbreak/go-cracking-coding/ch2"
	"github.com/dishbreak/go-cracking-coding/lib/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestKthToLast(t *testing.T) {
	type testCase struct {
		input  []int
		k      int
		result int
	}

	testCases := []testCase{
		{
			[]int{1, 2, 3, 4, 5, 5, 7},
			3, 5,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			assert.Equal(t, tc.result, ch2.KthToLast(linkedlist.FromSlice(tc.input), tc.k).Value)
		})
	}
}
