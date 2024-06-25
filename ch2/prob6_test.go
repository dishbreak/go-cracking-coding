package ch2_test

import (
	"fmt"
	"testing"

	"github.com/dishbreak/go-cracking-coding/ch2"
	"github.com/dishbreak/go-cracking-coding/lib/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	type testCase struct {
		input  []int
		result bool
	}

	testCases := []testCase{
		{[]int{4, 5, 5, 6, 5, 5, 4}, true},
		{[]int{1, 3, 3, 1}, true},
		{[]int{1, 2, 3, 5, 7, 8}, false},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			a := linkedlist.FromSlice(tc.input)
			// assert.Equal(t, tc.result, ch2.IsPalindromeV1(a))
			assert.Equal(t, tc.result, ch2.IsPalindromeV2(a))
		})
	}

}
