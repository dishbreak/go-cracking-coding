package linkedlist

type DoubleLinkedListNode[T any] struct {
	Data       T
	prev, next *DoubleLinkedListNode[T]
}

type DoubleLinkedList[T any] struct {
	head, tail *DoubleLinkedListNode[T]
	count      int
	zero       T
}

func (d *DoubleLinkedList[T]) PushLeft(v T) *DoubleLinkedListNode[T] {
	f := &DoubleLinkedListNode[T]{
		Data: v,
		next: d.head,
	}
	if d.Empty() {
		d.tail = f
	}
	d.head = f
	d.count++
	return f
}

func (d *DoubleLinkedList[T]) PopLeft() T {
	if d.Empty() {
		return d.zero
	}

	f := d.head

	d.head = d.head.next
	if d.head == nil {
		d.tail = nil
	}
	d.count--
	return f.Data
}

func (d *DoubleLinkedList[T]) PeekLeft() T {
	if d.Empty() {
		return d.zero
	}
	return d.head.Data
}

func (d *DoubleLinkedList[T]) PushRight(v T) *DoubleLinkedListNode[T] {
	f := &DoubleLinkedListNode[T]{
		Data: v,
		prev: d.tail,
	}
	d.count++
	if d.Empty() {
		d.head = f
	}
	d.tail = f
	return f
}

func (d *DoubleLinkedList[T]) PopRight() T {
	if d.Empty() {
		return d.zero
	}
	v := d.tail.Data
	d.tail = d.tail.prev
	d.count--
	if d.tail == nil {
		d.head = nil
	}
	return v
}

func (d *DoubleLinkedList[T]) PeekRight() T {
	if d.Empty() {
		return d.zero
	}

	return d.tail.Data
}

func (d *DoubleLinkedList[T]) Delete(n *DoubleLinkedListNode[T]) {
	if n.next == nil && n.prev == nil && d.head == n {
		d.head, d.tail = nil, nil
		d.count = 0
		return
	}

	if n.next != nil {
		n.next.prev = n.prev
	} else {
		d.tail = n.prev
	}

	if n.prev != nil {
		n.prev.next = n.next
	} else {
		d.head = n.next
	}
}

func (d *DoubleLinkedList[T]) Empty() bool {
	return d.head == nil
}

func (d *DoubleLinkedList[T]) Len() int {
	return d.count
}
