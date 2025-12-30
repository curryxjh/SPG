package array

import (
	"fmt"
	"iter"
)

type Array[T any] struct {
	data []T
}

func New[T any](size int) *Array[T] {
	return &Array[T]{
		data: make([]T, size, size),
	}
}

func NewFromArray[T any](other *Array[T]) *Array[T] {
	arr := &Array[T]{
		data: make([]T, other.Size(), other.Size()),
	}
	copy(arr.data, other.data)
	return arr
}

func (arr *Array[T]) Fill(val T) {
	for i := range arr.data {
		arr.data[i] = val
	}
}

func (arr *Array[T]) Size() int {
	return len(arr.data)
}

func (arr *Array[T]) Set(index int, val T) {
	if index < 0 || index >= arr.Size() {
		return
	}
	arr.data[index] = val
}

func (arr *Array[T]) At(index int) T {
	if index < 0 || index >= arr.Size() {
		panic("index out of range")
	}
	return arr.data[index]
}

func (arr *Array[T]) Front() T {
	return arr.At(0)
}

func (arr *Array[T]) Back() T {
	return arr.At(arr.Size() - 1)
}

func (arr *Array[T]) Empty() bool {
	return len(arr.data) == 0
}

func (arr *Array[T]) SwapArray(other *Array[T]) {
	if arr.Size() != other.Size() {
		return
	}
	arr.data, other.data = other.data, arr.data
}

func (arr *Array[T]) Data() []T {
	return arr.data
}

func (arr *Array[T]) String() string {
	return fmt.Sprintf("%v", arr.data)
}

func (arr *Array[T]) First() *ArrayIterator[T] {
	return arr.IterAt(0)
}

func (arr *Array[T]) Last() *ArrayIterator[T] {
	return arr.IterAt(arr.Size() - 1)
}

func (arr *Array[T]) Begin() *ArrayIterator[T] {
	return arr.First()
}

func (arr *Array[T]) End() *ArrayIterator[T] {
	return arr.IterAt(arr.Size())
}

func (arr *Array[T]) IterAt(index int) *ArrayIterator[T] {
	return &ArrayIterator[T]{
		arr:      arr,
		position: index,
	}
}

func (arr *Array[T]) Values() iter.Seq[T] {
	return arr.Begin().ToSeq()
}

func (arr *Array[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, val := range arr.data {
			if !yield(i, val) {
				return
			}
		}
	}
}
