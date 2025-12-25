package comparator

import "cmp"

type Ordered = cmp.Ordered

// Comparator is a function type that compares two values of type T and returns a integer.
// -1 if a < b
// 0 if a == b
// 1 if a > b
type Comparator[T any] func(a, b T) int

// OrderedTypeCmp is a comparator for ordered types.
func OrderedTypeCmp[T Ordered](a, b T) int {
	return cmp.Compare(a, b)
}

// Reverse is a function that returns a reverse comparator.
func Reverse[T any](cmp Comparator[T]) Comparator[T] {
	return func(a, b T) int {
		return -cmp(a, b)
	}
}

func BoolComparator(a, b bool) int {
	if a == b {
		return 0
	}
	if !a && b {
		return -1
	}
	return 1
}

func Complex64Comparator(a, b complex64) int {
	if a == b {
		return 0
	}
	if real(a) < real(b) {
		return -1
	}
	if real(a) == real(b) && imag(a) < imag(b) {
		return -1
	}
	return 1
}

func Complex128Comparator(a, b complex128) int {
	if a == b {
		return 0
	}
	if real(a) < real(b) {
		return -1
	}
	if real(a) == real(b) && imag(a) < imag(b) {
		return -1
	}
	return 1
}
