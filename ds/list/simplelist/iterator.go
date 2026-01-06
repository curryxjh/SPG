package simplelist

import (
	"SPL/utils/iterator"
	"iter"
)

type ListIterator[T any] struct {
	node *Node[T]
}

func NewIterator[T any](node *Node[T]) *ListIterator[T] {
	return &ListIterator[T]{node: node}
}

func (iter *ListIterator[T]) IsValid() bool {
	return iter.node != nil
}

func (iter *ListIterator[T]) Next() iterator.Cursor[T] {
	if iter.node != nil {
		iter.node = iter.node.next
	}
	return iter
}

func (iter *ListIterator[T]) Value() T {
	if iter.node == nil {
		panic("invalid iterator")
	}
	return iter.node.Value
}

func (iter *ListIterator[T]) Equal(other iterator.Cursor[T]) bool {
	otherIter, ok := other.(*ListIterator[T])
	if !ok {
		return false
	}
	if otherIter.node == iter.node {
		return true
	}
	return false
}

func (iter *ListIterator[T]) Clone() iterator.Cursor[T] {
	return NewIterator(iter.node)
}

func (iter *ListIterator[T]) SetValue(value T) {
	if iter.node != nil {
		iter.node.Value = value
	}
}

func (iter *ListIterator[T]) ToSeq() iter.Seq[T] {
	return func(yield func(T) bool) {
		for it := iter.node; it != nil; it = it.next {
			if !yield(it.Value) {
				return
			}
		}
	}
}
