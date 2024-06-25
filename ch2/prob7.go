package ch2

import "github.com/dishbreak/go-cracking-coding/lib/linkedlist"

func Intersection(a, b *linkedlist.Node) *linkedlist.Node {
	ac, bc := a, b
	for ; ac != nil && bc != nil; ac, bc = ac.Next, bc.Next {
	}

	if ac == nil && bc == nil {
		ac, bc = a, b
	} else if ac == nil {
		diff := remaining(bc)
		bc = advance(b, diff)
		ac = a
	} else {
		diff := remaining(ac)
		ac = advance(a, diff)
		bc = b
	}

	for ; ac != nil; ac, bc = ac.Next, bc.Next {
		if ac == bc {
			return ac
		}
	}

	return nil
}

func remaining(h *linkedlist.Node) int {
	i := 0

	for c := h; c != nil; c = c.Next {
		i++
	}

	return i
}

func advance(h *linkedlist.Node, c int) *linkedlist.Node {
	for i := 0; i < c; i++ {
		h = h.Next
	}

	return h
}
