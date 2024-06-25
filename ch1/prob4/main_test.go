package prob4_test

import (
	"testing"

	"github.com/dishbreak/go-cracking-coding/ch1/prob4"
	"github.com/stretchr/testify/assert"
)

func TestHasPalindromePermutation(t *testing.T) {
	assert.True(t, prob4.HasPalindromePermutation("Tact Coa"))
	assert.False(t, prob4.HasPalindromePermutation("Tact Coals"))
}
