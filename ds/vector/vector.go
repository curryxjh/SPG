package vector

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

func (v *Vector[T]) EraseAt(index int) {}

func (v *Vector[T]) EraseIndexRange(first, last int) {}

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

func (v *Vector[T]) Reverse() {
	for i := 0; i < v.Size()/2; i++ {
		v.data[i], v.data[v.Size()-1-i] = v.data[v.Size()-1-i], v.data[i]
	}
}

func (v *Vector[T]) Data() []T {
	return v.data
}

func (v *Vector[T]) Resize(size int) {

}
