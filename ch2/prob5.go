package ch2

import "github.com/dishbreak/go-cracking-coding/lib/linkedlist"

func SumReversed(a, b *linkedlist.Node) *linkedlist.Node {
	var result *linkedlist.Node
	var tail *linkedlist.Node
	carryover := 0
	ac, bc := a, b
	for ; ac != nil && bc != nil; ac, bc = ac.Next, bc.Next {
		digit := ac.Value + bc.Value + carryover
		if digit >= 10 {
			carryover = 1
			digit = digit % 10
		} else {
			carryover = 0
		}

		n := &linkedlist.Node{
			Value: digit,
		}
		if result == nil {
			result = n
			tail = n
		} else {
			tail.Next = n
			tail = n
		}
	}

	var remainder *linkedlist.Node
	if ac == nil && bc == nil {
		return result
	} else if ac != nil {
		remainder = ac
	} else {
		remainder = bc
	}

	for r := remainder; r != nil; r = r.Next {
		digit := r.Value + carryover
		if digit >= 10 {
			carryover = 1
			digit = digit % 10
		} else {
			carryover = 0
		}
		n := &linkedlist.Node{
			Value: digit,
		}
		tail.Next = n
		tail = n
	}

	return result
}

func normalize(a, b *linkedlist.Node) (*linkedlist.Node, *linkedlist.Node) {
	ac, bc := a, b
	for ; ac != nil && bc != nil; ac, bc = ac.Next, bc.Next {
	}

	if ac == nil && bc == nil {
		return a, b
	}

	padding := 0
	var longer, shorter, tail *linkedlist.Node
	if ac != nil {
		tail = ac
		longer = a
		shorter = b
	} else {
		tail = bc
		longer = b
		shorter = a
	}

	for ; tail != nil; tail = tail.Next {
		padding++
	}

	for i := 0; i < padding; i++ {
		n := &linkedlist.Node{
			Value: 0,
		}
		n.Next = shorter
		shorter = n
	}

	return longer, shorter
}

func SumForwards(a, b *linkedlist.Node) *linkedlist.Node {
	a, b = normalize(a, b)
	head, carryover := sumRecursive(a, b)
	if carryover > 0 {
		n := &linkedlist.Node{
			Value: carryover,
			Next:  head,
		}
		head = n
	}
	return head
}

func sumRecursive(a, b *linkedlist.Node) (result *linkedlist.Node, carryover int) {
	result = &linkedlist.Node{}
	if a.Next != nil {
		var tail *linkedlist.Node
		tail, carryover = sumRecursive(a.Next, b.Next)
		result.Next = tail
	}
	result.Value = a.Value + b.Value + carryover
	if result.Value >= 10 {
		carryover = 1
		result.Value = result.Value % 10
	}
	return
}
