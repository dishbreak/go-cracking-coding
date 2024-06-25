package linkedlist

type Node struct {
	Value int
	Next  *Node
}

func FromSlice(v []int) *Node {
	if len(v) == 0 {
		return nil
	}

	head := &Node{
		Value: v[0],
	}

	cond := head
	for i := 1; i < len(v); i++ {
		cond.Next = &Node{
			Value: v[i],
		}
		cond = cond.Next
	}
	return head
}

func ToSlice(head *Node) []int {
	result := make([]int, 0)
	for c := head; c != nil; c = c.Next {
		result = append(result, c.Value)
	}
	return result
}

func Index(head *Node, i int) *Node {
	c := head
	for j := 0; j < i; j++ {
		if c == nil {
			return nil
		}
		c = c.Next
	}
	return c
}

func Copy(h *Node) *Node {
	if h == nil {
		return nil
	}

	var r *Node

	for c := h; h != nil; h = h.Next {
		n := &Node{
			Value: c.Value,
		}
		if r == nil {
			r = n
			continue
		}
		r.Next = n
		r = r.Next
	}

	return r
}

func Reverse(h *Node) *Node {
	if h == nil || h.Next == nil {
		return nil
	}

	r := Reverse(h.Next)
	h.Next.Next = h
	h.Next = nil
	return r
}

func Tail(h *Node) *Node {
	for ; h.Next != nil; h = h.Next {
	}

	return h
}
