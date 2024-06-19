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
