package iterator

type ConstIterator[T any] interface {
	IsValid() bool
	Next() ConstIterator[T]
	Value() T
	Clone() ConstIterator[T]
	Equal(other ConstIterator[T]) bool
}

type Iterator[T any] interface {
}
