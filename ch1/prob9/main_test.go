package prob9_test

import (
	"fmt"
	"testing"

	"github.com/dishbreak/go-cracking-coding/ch1/prob9"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	s1, s2 string
	result bool
}

var testCases = []testCase{
	{"waterbottle", "erbottlewat", true},
	{"doo", "foo", false},
}

func TestIsRotation(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.result, prob9.IsRotation(tc.s1, tc.s2))
		})
	}
}
