package prob5_test

import (
	"fmt"
	"testing"

	"github.com/dishbreak/go-cracking-coding/ch1/prob5"
	"github.com/stretchr/testify/assert"
)

func TestIsOneAway(t *testing.T) {
	type testCase struct {
		a, b   string
		result bool
	}

	testCases := []testCase{
		{"pale", "ple", true},
		{"pales", "pale", true},
		{"pale", "bale", true},
		{"pale", "bake", false},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			assert.Equal(t, tc.result, prob5.IsOneAway(tc.a, tc.b))
		})
	}
}
