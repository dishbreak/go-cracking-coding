package prob1_test

import (
	"fmt"
	"testing"

	"github.com/dishbreak/go-cracking-coding/ch1/prob1"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input  string
	result bool
}

var testCases = []testCase{
	{"abba", false},
	{"asbtri", true},
	{"a", true},
	{"aa", false},
	{"asbtria", false},
}

func TestHasUniqueCharacters(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.result, prob1.HasUniqueChars(tc.input))
		})
	}
}

func TestHasUniqueCharactersNoDS(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.result, prob1.HasUniqueCharsNoDataStructure(tc.input))
		})
	}
}
