package ch2

import "github.com/dishbreak/go-cracking-coding/lib/linkedlist"

func KthToLast(head *linkedlist.Node, k int) *linkedlist.Node {
	r := head

	for i := 0; i < k; i++ {
		if r == nil {
			return nil
		}
		r = r.Next
	}

	c := head
	for r != nil {
		c = c.Next
		r = r.Next
	}

	return c
}
