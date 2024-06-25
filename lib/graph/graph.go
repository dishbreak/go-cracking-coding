package graph

type Node[T any] struct {
	Data      T
	Neighbors []*Node[T]
}
