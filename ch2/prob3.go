package ch2

import (
	"errors"

	"github.com/dishbreak/go-cracking-coding/lib/linkedlist"
)

func DeleteMiddle(n *linkedlist.Node) {
	if n.Next == nil {
		panic(errors.New("can't delete end of list"))
	}

	n.Value = n.Next.Value
	n.Next = n.Next.Next
}
