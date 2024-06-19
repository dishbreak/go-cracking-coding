package ch2

import "github.com/dishbreak/go-cracking-coding/lib/linkedlist"

func Dedupe(head *linkedlist.Node) *linkedlist.Node {
	result := &linkedlist.Node{
		Value: head.Value,
	}

	seen := make(map[int]bool)
	seen[result.Value] = true

	r := result
	for c := head.Next; c != nil; c = c.Next {
		if seen[c.Value] {
			continue
		}
		r.Next = &linkedlist.Node{
			Value: c.Value,
		}
		seen[c.Value] = true
		r = r.Next
	}

	return result
}

func DedupeNoSpace(head *linkedlist.Node) *linkedlist.Node {
	for c := head; c != nil; c = c.Next {
		for r := c; r.Next != nil; {
			if r.Next.Value == c.Value {
				r.Next = r.Next.Next
				continue
			}
			r = r.Next
		}
	}
	return head
}
