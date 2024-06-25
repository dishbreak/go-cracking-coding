package ch3_test

import (
	"testing"

	"github.com/dishbreak/go-cracking-coding/ch3"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := ch3.NewMyQueue()

	assert.Equal(t, 0, q.Dequeue())
	q.Enqueue(5)
	q.Enqueue(7)
	q.Enqueue(9)
	assert.Equal(t, 5, q.Dequeue())
	q.Enqueue(8)
	q.Enqueue(3)

	for _, v := range []int{7, 9, 8, 3} {
		assert.Equal(t, v, q.Dequeue())
	}

	assert.Equal(t, 0, q.Dequeue())
}
