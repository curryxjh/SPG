package bidlist

import (
	"errors"
)

type T any

var ErrorOutofRange = errors.New("out of range")

type Node[T any] struct {
	Value T
	prev  *Node[T]
	next  *Node[T]
	list  *List[T]
}

func (n *Node[T]) Next() *Node[T] {
	if n.list == nil {
		return nil
	}
	if n.next == n.list.head {
		return nil
	}
	return n.next
}

func (n *Node[T]) Prev() *Node[T] {
	if n.list == nil {
		return nil
	}
	return n.prev
}

type List[T any] struct {
	head *Node[T]
	len  int
}

func New[T any]() List[T] {
	list := List[T]{
		head: &Node[T]{},
		len:  0,
	}
	list.head.next = list.head
	list.head.prev = list.head
	return list
}

func (l *List[T]) Len() int {
	return l.len
}

func (l *List[T]) Size() int {
	return l.len
}

func (l *List[T]) Empty() bool {
	return l.len == 0
}

func (l *List[T]) FrontNode() *Node[T] {
	return l.head.next
}

func (l *List[T]) BackNode() *Node[T] {
	if l.head.next == nil {
		return nil
	}
	return l.head.prev
}

func (l *List[T]) Front() T {
	if l.len == 0 {
		panic(ErrorOutofRange)
	}
	return l.head.next.Value
}

func (l *List[T]) Back() T {
	if l.len == 0 {
		panic(ErrorOutofRange)
	}
	return l.head.prev.Value
}

func (l *List[T]) PushBack(v T) {
	l.pushBack(v)
}

func (l *List[T]) PushFront(v T) {
	l.pushFront(v)
}

func (l *List[T]) pushFront(v T) *Node[T] {
	node := &Node[T]{
		Value: v,
		list:  l,
	}
	return l.insertAfter(node, l.head)
}

func (l *List[T]) pushBack(v T) *Node[T] {
	node := &Node[T]{Value: v, list: l}
	return l.insertAfter(node, l.head.prev)
}

func (l *List[T]) InsertAfter(v T, mark *Node[T]) *Node[T] {
	if mark.list != l {
		return nil
	}
	return l.insertAfter(&Node[T]{Value: v, list: l}, mark)
}

func (l *List[T]) insertAfter(n, at *Node[T]) *Node[T] {
	n.next = at.next
	n.prev = at
	at.next.prev = n
	at.next = n
	l.len++
	return n
}
