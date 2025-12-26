package vector

import "SPL/utils/iterator"

type T any

var _ iterator.RandomAccessIterator[T] = (*VectorIterator[T])(nil)

type VectorIterator[T any] struct {
	vec      *Vector[T]
	position int
}

func (v VectorIterator[T]) IsValid() bool {
	//TODO implement me
	panic("implement me")
}

func (v VectorIterator[T]) Value() T {
	//TODO implement me
	panic("implement me")
}

func (v VectorIterator[T]) Next() iterator.Cursor[T] {
	//TODO implement me
	panic("implement me")
}

func (v VectorIterator[T]) Clone() iterator.Cursor[T] {
	//TODO implement me
	panic("implement me")
}

func (v VectorIterator[T]) Equal(other iterator.Cursor[T]) bool {
	//TODO implement me
	panic("implement me")
}

func (v VectorIterator[T]) SetValue(value T) {
	//TODO implement me
	panic("implement me")
}

func (v VectorIterator[T]) IteratorAt(position int) iterator.RandomAccessIterator[T] {
	//TODO implement me
	panic("implement me")
}

func (v VectorIterator[T]) Position() int {
	//TODO implement me
	panic("implement me")
}
