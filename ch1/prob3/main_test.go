package prob3_test

import (
	"testing"

	"github.com/dishbreak/go-cracking-coding/ch1/prob3"
	"github.com/stretchr/testify/assert"
)

func TestURLify(t *testing.T) {
	assert.Equal(t, "Mr%20John%20Smith", prob3.URLify("Mr John Smith    "))
}
