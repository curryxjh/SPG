package vector

import (
	"SPL/utils/iterator"
	"iter"
)

type T any

// VectorIterator is an implementation of iterator.RandomAccessIterator for Vector.
var _ iterator.RandomAccessIterator[T] = (*VectorIterator[T])(nil)

type VectorIterator[T any] struct {
	vec      *Vector[T]
	position int
}

func (iter *VectorIterator[T]) IsValid() bool {
	return iter.position >= 0 && iter.position < iter.vec.Size()
}

func (iter *VectorIterator[T]) Value() T {
	val := iter.vec.At(iter.position)
	return val
}

func (iter *VectorIterator[T]) Next() iterator.Cursor[T] {
	if iter.position < iter.vec.Size() {
		iter.position++
	}
	return iter
}

func (iter *VectorIterator[T]) Prev() iterator.ConstBidIterator[T] {
	if iter.position >= 0 {
		iter.position--
	}
	return iter
}

func (iter *VectorIterator[T]) Clone() iterator.Cursor[T] {
	return &VectorIterator[T]{vec: iter.vec, position: iter.position}
}

func (iter *VectorIterator[T]) Equal(other iterator.Cursor[T]) bool {
	otherIter, ok := other.(*VectorIterator[T])
	if !ok {
		return false
	}
	if otherIter.vec == iter.vec && otherIter.position == iter.position {
		return true
	}
	return false
}

func (iter *VectorIterator[T]) SetValue(value T) {
	iter.vec.SetAt(iter.position, value)
}

func (iter *VectorIterator[T]) IteratorAt(position int) iterator.RandomAccessIterator[T] {
	return &VectorIterator[T]{vec: iter.vec, position: position}
}

func (iter *VectorIterator[T]) Position() int {
	return iter.position
}

func (iter *VectorIterator[T]) ToSeq() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, val := range iter.vec.data {
			if !yield(val) {
				return
			}
		}
	}
}
