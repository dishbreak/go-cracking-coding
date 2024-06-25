package prob2_test

import (
	"fmt"
	"testing"

	"github.com/dishbreak/go-cracking-coding/ch1/prob2"
	"github.com/stretchr/testify/assert"
)

func TestArePermutations(t *testing.T) {
	type testCase struct {
		one, other string
		result     bool
	}

	testCases := []testCase{
		{"thought", "hottghu", true},
		{"cat", "bake", false},
		{"taco", "book", false},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			assert.Equal(t, tc.result, prob2.ArePermutations(tc.one, tc.other))
		})
	}
}
