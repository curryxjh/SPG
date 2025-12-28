package vector

import (
	"SPL/utils/iterator"
	"fmt"
	"iter"
)

// Options is the options for vector.
type Options struct {
	Capacity int
}

// NewVector creates a new vector with the given options.
type Option func(option *Options)

// WithCapacity sets the capacity of the vector.
func WithCapacity(capacity int) Option {
	return func(option *Options) {
		option.Capacity = capacity
	}
}

// Vector is a linear structure.
type Vector[T any] struct {
	data []T
}

func New[T any](opts ...Option) *Vector[T] {
	option := Options{}
	for _, opt := range opts {
		opt(&option)
	}
	return &Vector[T]{
		data: make([]T, 0, option.Capacity),
	}
}

func NewFromVector[T any](other *Vector[T]) *Vector[T] {
	v := &Vector[T]{
		data: make([]T, other.Size(), other.Capacity()),
	}
	for i := range other.data {
		v.data[i] = other.data[i]
	}
	return v
}

func (v *Vector[T]) Size() int {
	return len(v.data)
}

func (v *Vector[T]) Capacity() int {
	return cap(v.data)
}

func (v *Vector[T]) Empty() bool {
	return len(v.data) == 0
}

func (v *Vector[T]) PushBack(value T) {
	v.data = append(v.data, value)
}

func (v *Vector[T]) PopBack() T {
	if v.Empty() {
		panic("pop from empty vector")
	}
	val := v.data[len(v.data)-1]
	v.data = v.data[:len(v.data)-1]
	return val
}

func (v *Vector[T]) SetAt(index int, value T) {
	if index < 0 || index >= v.Size() {
		return
	}
	v.data[index] = value
}

func (v *Vector[T]) InsertAt(index int, value T) {
	if index < 0 || index > v.Size() {
		return
	}
	v.data = append(v.data, value)
	for i := v.Size() - 1; i > index; i-- {
		v.data[i] = v.data[i-1]
	}
	v.data[index] = value
}

func (v *Vector[T]) EraseAt(index int) {
	v.EraseIndexRange(index, index+1)
}

// EraseIndexRange erases the range of indices [first, last).
func (v *Vector[T]) EraseIndexRange(first, last int) {
	if first > last {
		return
	}
	if first < 0 || last >= v.Size() {
		return
	}
	left := v.data[:first]
	right := v.data[last+1:]
	v.data = append(left, right...)
}

func (v *Vector[T]) Clear() {
	v.data = v.data[:0]
}

func (v *Vector[T]) At(index int) T {
	if index < 0 || index >= v.Size() {
		panic("index out of range")
	}
	return v.data[index]
}

func (v *Vector[T]) Front() T {
	return v.At(0)
}

func (v *Vector[T]) Back() T {
	return v.At(v.Size() - 1)
}

func (v *Vector[T]) Begin() *VectorIterator[T] {
	return v.First()
}

func (v *Vector[T]) End() *VectorIterator[T] {
	return v.IterAt(v.Size())
}

func (v *Vector[T]) First() *VectorIterator[T] {
	return v.IterAt(0)
}

func (v *Vector[T]) Last() *VectorIterator[T] {
	return v.IterAt(v.Size() - 1)
}

func (v *Vector[T]) IterAt(index int) *VectorIterator[T] {
	return &VectorIterator[T]{
		vec:      v,
		position: index,
	}
}

// Insert inserts a value at the given iterator. Returns the iterator after the inserted element.
func (v *Vector[T]) Insert(iter iterator.Cursor[T], value T) *VectorIterator[T] {
	index := iter.(*VectorIterator[T]).position
	v.InsertAt(index, value)
	return &VectorIterator[T]{
		vec:      v,
		position: index,
	}
}

// Erase erases the element at the given iterator. Returns the iterator after the erased element.
func (v *Vector[T]) Erase(iter iterator.Cursor[T]) *VectorIterator[T] {
	index := iter.(*VectorIterator[T]).position
	v.EraseAt(index)
	return &VectorIterator[T]{
		vec:      v,
		position: index,
	}
}

func (v *Vector[T]) EraseRange(first, last iterator.Cursor[T]) *VectorIterator[T] {
	begin := first.(*VectorIterator[T]).position
	end := last.(*VectorIterator[T]).position
	v.EraseIndexRange(begin, end)
	return &VectorIterator[T]{
		vec:      v,
		position: begin,
	}
}

func (v *Vector[T]) Reverse() {
	for i := 0; i < v.Size()/2; i++ {
		v.data[i], v.data[v.Size()-1-i] = v.data[v.Size()-1-i], v.data[i]
	}
}

// Reserve reserves the capacity of the vector. If the capacity is already larger than the given capacity, it does nothing.
func (v *Vector[T]) Reserve(capacity int) {
	if cap(v.data) >= capacity {
		return
	}
	data := make([]T, v.Size(), capacity)
	copy(data, v.data)
	v.data = data
}

func (v *Vector[T]) ShrinkToFit() {
	if len(v.data) == cap(v.data) {
		return
	}
	origin_length := v.Size()
	data := make([]T, origin_length, origin_length)
	copy(data, v.data)
	v.data = data
}

func (v *Vector[T]) Data() []T {
	return v.data
}

func (v *Vector[T]) Resize(size int) {
	if size > v.Size() {
		return
	}
	v.data = v.data[:size]
}

func (v *Vector[T]) String() string {
	return fmt.Sprintf("%v", v.data)
}

func (v *Vector[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, val := range v.data {
			if !yield(i, val) {
				return
			}
		}
	}
}

func (v *Vector[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, val := range v.data {
			if !yield(val) {
				return
			}
		}
	}
}
