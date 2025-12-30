package iterator

import "iter"

// Cursor is an interface for iterators.
type Cursor[T any] interface {
	IsValid() bool
	Value() T
	Next() Cursor[T]
	Clone() Cursor[T]
	Equal(other Cursor[T]) bool
}

// ConstIterator is an interface for constant iterators.
type ConstIterator[T any] interface {
	Cursor[T]
	// All returns a sequence of all elements in the container.
	ToSeq() iter.Seq[T]
}

// Iterator is an interface for mutable iterators.
type Iterator[T any] interface {
	ConstIterator[T]
	SetValue(value T)
}

// ConstKvIterator is an interface for constant key-value iterators.
type ConstKvIterator[K any, V any] interface {
	Cursor[V]
	Key() K
	ToSeq2() iter.Seq2[K, V]
}

// KvIterator is an interface for mutable key-value iterators.
type KvIterator[K any, V any] interface {
	ConstKvIterator[K, V]
	SetValue(value V)
}

// ConstBidIterator is an interface for constant bidirectional iterators.
type ConstBidIterator[T any] interface {
	ConstIterator[T]
	Prev() ConstBidIterator[T]
}

// BidIterator is an interface for mutable bidirectional iterators.
type BidIterator[T any] interface {
	ConstBidIterator[T]
	SetValue(value T)
}

// ConstKvBidIterator is an interface for constant bidirectional key-value iterators.
type ConstKvBidIterator[K any, V any] interface {
	ConstKvIterator[K, V]
	BidIterator[V]
}

// KvBidIterator is an interface for mutable bidirectional key-value iterators.
type KvBidIterator[K any, V any] interface {
	KvIterator[K, V]
	BidIterator[V]
}

// RandomAccessIterator is an interface for mutable random access iterators.
type RandomAccessIterator[T any] interface {
	BidIterator[T]
	// Returns an iterator pointing to the element at the specified position in the container.
	IteratorAt(position int) RandomAccessIterator[T]
	Position() int
}
