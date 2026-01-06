package simplelist

import (
	"SPL/utils/visitor"
	"fmt"
)

// Node is the node of the simple list
type Node[T any] struct {
	next  *Node[T]
	Value T
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// head -> node_1 -> node_2 -> ... -> node_n <- tail
type List[T any] struct {
	head *Node[T] // head of the list
	tail *Node[T] // tail of the list
	len  int
}

func New[T any]() *List[T] {
	return &List[T]{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (l *List[T]) Len() int {
	return l.len
}

func (l *List[T]) FrontNode() *Node[T] {
	return l.head
}

func (l *List[T]) BackNode() *Node[T] {
	return l.tail
}

func (l *List[T]) PushFront(v T) {
	n := &Node[T]{Value: v}
	if l.len == 0 {
		l.head = n
		l.tail = n
	} else {
		n.next = l.head
		l.head = n
	}
	l.len++
}

func (l *List[T]) PushBack(v T) {
	n := &Node[T]{Value: v}
	if l.len == 0 {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}
	l.len++
}

func (l *List[T]) InsertAfter(v T, at *Node[T]) *Node[T] {
	return l.insertAfter(&Node[T]{Value: v}, at)
}

func (l *List[T]) insertAfter(n, at *Node[T]) *Node[T] {
	n.next = at.next
	at.next = n
	if n.next == nil {
		l.tail = n
	}
	l.len++
	return n
}

func (l *List[T]) Remove(pre, n *Node[T]) T {
	if n == nil {
		return *new(T)
	}
	if pre == nil {
		// remove head
		l.head = n.next
		if l.head == nil {
			l.tail = nil
		}
	} else {
		pre.next = n.next
		if pre.next == nil {
			l.tail = pre
		}
	}
	l.len--
	return n.Value
}

func (l *List[T]) MoveToFront(pre, n *Node[T]) {
	if pre == nil || n == nil || pre.next != n || l.len <= 1 {
		return
	}
	pre.next = n.next
	if pre.next == nil {
		l.tail = pre
	}
	n.next = l.head
	l.head = n
}

func (l *List[T]) MoveToBack(pre, n *Node[T]) {
	if n == nil || n.next == nil || l.len <= 1 {
		return
	}
	if pre == nil {
		l.head = n.next
	} else {
		pre.next = n.next
	}
	l.tail.next = n
	l.tail = n
	n.next = nil
}

func (l *List[T]) String() string {
	str := "["
	for n := l.head; n != nil; n = n.next {
		if str != "[" {
			str += " "
		}
		str += fmt.Sprintf("%v", n.Value)
	}
	str += "]"
	return str
}

func (l *List[T]) Traversal(visitor visitor.Visitor[T]) {
	for node := l.head; node != nil; node = node.next {
		if !visitor(node.Value) {
			break
		}
	}
}
