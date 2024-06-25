package ch3

import (
	"github.com/dishbreak/go-cracking-coding/lib/linkedlist"
	"github.com/dishbreak/go-cracking-coding/lib/queue/v2"
)

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d *Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c *Cat) Speak() string {
	return "Meow!"
}

type Horse struct{}

func (h *Horse) Speak() string {
	return "Neigh!"
}

type AnimalQueue struct {
	q   *linkedlist.DoubleLinkedList[Animal]
	dog queue.Queue[*linkedlist.DoubleLinkedListNode[Animal]]
	cat queue.Queue[*linkedlist.DoubleLinkedListNode[Animal]]
}

func (aq *AnimalQueue) Enqueue(a Animal) {
	f := aq.q.PushLeft(a)
	switch a.(type) {
	case *Dog:
		aq.dog.Enqueue(f)
	case *Cat:
		aq.cat.Enqueue(f)
	}
}

func (aq *AnimalQueue) Dequeue() Animal {
	a := aq.q.PopRight()
	switch a.(type) {
	case *Dog:
		f := aq.dog.Peek()
		if f != nil {
			if a.(*Dog) == f.Data.(*Dog) {
				aq.dog.Dequeue()
			}
		}
	case *Cat:
		f := aq.cat.Peek()
		if f != nil {
			if a.(*Cat) == f.Data.(*Cat) {
				aq.cat.Dequeue()
			}
		}
	}
	return a
}

func (aq *AnimalQueue) DequeueDog() *Dog {
	f := aq.dog.Dequeue()
	if f == nil {
		return nil
	}
	aq.q.Delete(f)
	return f.Data.(*Dog)
}

func (aq *AnimalQueue) DequeueCat() *Cat {
	f := aq.cat.Dequeue()
	if f == nil {
		return nil
	}
	aq.q.Delete(f)
	return f.Data.(*Cat)
}
