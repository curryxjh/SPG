package array

import (
	"SPL/utils/iterator"
	"iter"
)

type T any

var _ iterator.RandomAccessIterator[T] = (*ArrayIterator[T])(nil)

type ArrayIterator[T any] struct {
	arr      *Array[T]
	position int
}

func (iter *ArrayIterator[T]) IsValid() bool {
	return iter.position >= 0 && iter.position < iter.arr.Size()
}

func (iter *ArrayIterator[T]) Value() T {
	val := iter.arr.At(iter.position)
	return val
}

func (iter *ArrayIterator[T]) Next() iterator.Cursor[T] {
	if iter.position < iter.arr.Size() {
		iter.position++
	}
	return iter
}

func (iter *ArrayIterator[T]) Clone() iterator.Cursor[T] {
	return &ArrayIterator[T]{arr: iter.arr, position: iter.position}
}

func (iter *ArrayIterator[T]) Equal(other iterator.Cursor[T]) bool {
	otherIter, ok := other.(*ArrayIterator[T])
	if !ok {
		return false
	}
	if otherIter.arr == iter.arr && otherIter.position == iter.position {
		return true
	}
	return false
}

func (iter *ArrayIterator[T]) ToSeq() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, val := range iter.arr.data {
			if !yield(val) {
				return
			}
		}
	}
}

func (iter *ArrayIterator[T]) Prev() iterator.ConstBidIterator[T] {
	if iter.position > 0 {
		iter.position--
	}
	return iter
}

func (iter *ArrayIterator[T]) SetValue(value T) {
	iter.arr.Set(iter.position, value)
}

func (iter *ArrayIterator[T]) IteratorAt(position int) iterator.RandomAccessIterator[T] {
	return &ArrayIterator[T]{arr: iter.arr, position: position}
}

func (iter *ArrayIterator[T]) Position() int {
	return iter.position
}
