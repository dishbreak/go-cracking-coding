package prob6_test

import (
	"fmt"
	"testing"

	"github.com/dishbreak/go-cracking-coding/ch1/prob6"
	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	type testCase struct {
		input, result string
	}

	testCases := []testCase{
		{"aabcccccaaa", "a2b1c5a3"},
		{"abcabcabc", "abcabcabc"},
		{"aabbcc", "aabbcc"},
		{"ab", "ab"},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			assert.Equal(t, tc.result, prob6.Compress(tc.input))
		})
	}
}
