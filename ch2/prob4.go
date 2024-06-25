package ch2

import "github.com/dishbreak/go-cracking-coding/lib/linkedlist"

func Partition(head *linkedlist.Node, pivot int) *linkedlist.Node {
	if head == nil {
		return head
	}

	n := &linkedlist.Node{
		Value: head.Value,
	}

	h, t := n, n

	for c := head; c != nil; c = c.Next {
		m := &linkedlist.Node{
			Value: c.Value,
		}
		if m.Value < pivot {
			m.Next = h
			h = m
		} else {
			t.Next = m
			t = m
		}
	}

	return h
}
