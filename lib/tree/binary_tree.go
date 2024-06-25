package tree

type BinaryTreeNode[T any] struct {
	Data        T
	Left, Right *BinaryTreeNode[T]
}

func (n *BinaryTreeNode[T]) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}
