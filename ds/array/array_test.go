package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayBase(t *testing.T) {
	a := New[int](10)
	assert.Equal(t, 10, a.Size())
	assert.Equal(t, false, a.Empty())

	valA := 10
	a.Fill(valA)
	for v := range a.Values() {
		assert.Equal(t, valA, v)
	}
	for i, v := range a.All() {
		k := a.At(i)
		assert.Equal(t, k, v)
	}
	valB := 66
	b := New[int](10)
	b.Fill(valB)
	a.SwapArray(b)
	for v := range b.Values() {
		assert.Equal(t, valA, v)
	}
	for i, v := range b.All() {
		ka := a.At(i)
		assert.Equal(t, ka, valB)
		kb := b.At(i)
		assert.Equal(t, kb, v)
	}

	for i := 0; i < a.Size(); i++ {
		a.Set(i, i)
	}
	for i, v := range a.All() {
		assert.Equal(t, i, v)
	}

	i := 0
	for iter := a.First(); iter.IsValid(); iter.Next() {
		assert.Equal(t, i, iter.Value())
		i++
	}
	i = a.Size() - 1
	for iter := a.Last(); iter.IsValid(); iter.Prev() {
		assert.Equal(t, i, iter.Value())
		i--
	}
}
func TestNewFromArray(t *testing.T) {
	a := New[int](10)
	for i := 0; i < a.Size(); i++ {
		a.Set(i, i*10)
	}
	b := NewFromArray(a)
	for i := 0; i < a.Size(); i++ {
		assert.Equal(t, a.At(i), b.At(i))
	}
}
