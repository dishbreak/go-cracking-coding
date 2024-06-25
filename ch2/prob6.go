package ch2

import "github.com/dishbreak/go-cracking-coding/lib/linkedlist"

func IsPalindromeV1(h *linkedlist.Node) bool {
	r := linkedlist.Copy(h)
	r = linkedlist.Reverse(r)

	for c, d := h, r; c != nil; c, d = c.Next, d.Next {
		if c.Value != d.Value {
			return false
		}
	}

	return true
}

func IsPalindromeV2(h *linkedlist.Node) bool {
	mid := getMidpoint(h)
	stack := make([]int, 0)
	for c := mid; c != nil; c = c.Next {
		stack = append([]int{c.Value}, stack...)
	}

	for c := h; len(stack) != 0; stack, c = stack[1:], c.Next {
		if stack[0] != c.Value {
			return false
		}
	}

	return true
}

func getMidpoint(h *linkedlist.Node) *linkedlist.Node {
	mid := h

	counter := 1
	for ; h != nil; counter, h = counter+1, h.Next {
		if counter%2 == 0 {
			mid = mid.Next
		}
	}

	return mid
}
