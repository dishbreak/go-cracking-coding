package ch2

import "github.com/dishbreak/go-cracking-coding/lib/linkedlist"

func LoopedPoint(h *linkedlist.Node) *linkedlist.Node {
	fast, slow := h, h

	for fast.Next != nil && fast.Next.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			break
		}
	}

	if fast != slow {
		return nil
	}

	slow = h

	for ; fast != slow; fast, slow = fast.Next, slow.Next {
	}

	return slow
}
