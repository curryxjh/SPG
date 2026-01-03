package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorBase(t *testing.T) {
	v := New[int](WithCapacity(10))
	assert.True(t, v.Empty())
	assert.Equal(t, 10, v.Capacity())

	v.PushBack(1)
	v.PushBack(2)

	assert.False(t, v.Empty())
	assert.Equal(t, 2, v.Size())
	assert.Equal(t, 1, v.Front())
	assert.Equal(t, 2, v.Back())
	assert.Equal(t, []int{1, 2}, v.Data())
}

func TestVectorResize(t *testing.T) {
	v := New[int](WithCapacity(10))
	v.PushBack(1)
	v.PushBack(2)
	v.ShrinkToFit()
	assert.Equal(t, 2, v.Size())
	assert.Equal(t, 2, v.Capacity())

	assert.Equal(t, 1, v.At(0))
	assert.Equal(t, 2, v.At(1))

	v.Reserve(20)
	assert.Equal(t, 20, v.Capacity())
	assert.Equal(t, 2, v.Size())
	assert.Equal(t, 2, v.At(1))

	v.Clear()
	assert.Equal(t, 0, v.Size())
	assert.True(t, v.Empty())

	for i := range 10 {
		v.PushBack(i)
	}
	assert.Equal(t, 10, v.Size())
	v.Resize(20)
	assert.Equal(t, 10, v.Size())
	v.Resize(4)
	assert.Equal(t, 4, v.Size())

	b := NewFromVector(v)
	assert.Equal(t, 4, b.Size())
	assert.Equal(t, "[0 1 2 3]", b.String())
}

func TestVectorModify(t *testing.T) {
	v := New[int]()
	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)
	// [1 2 3]
	assert.Equal(t, 3, v.PopBack())
	v.PushBack(4)
	// [1 2 4]
	assert.Equal(t, 4, v.Back())

	v.SetAt(1, 9)
	assert.Equal(t, 9, v.At(1))
	// [1 9 4]

	v.InsertAt(0, 8)
	// [8 1 9 4]
	assert.Equal(t, 8, v.At(0))
	assert.Equal(t, "[8 1 9 4]", v.String())
	v.Clear()
}

func TestVectorIterator(t *testing.T) {
	v := New[int]()
	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)
	v.PushBack(4)
	// [1 2 3 4]
	i := 0
	for iter := v.Begin(); iter.IsValid(); iter.Next() {
		assert.Equal(t, i+1, iter.Value())
		i++
	}

	i = 0
	for val := range v.Values() {
		assert.Equal(t, i+1, val)
		i++
	}

	i = 0
	for iter, val := range v.All() {
		assert.Equal(t, i, iter)
		assert.Equal(t, i+1, val)
		i++
	}

	i = 3
	for iter := v.Last(); iter.IsValid(); iter.Prev() {
		assert.Equal(t, i+1, iter.Value())
		i--
	}

	iter := v.Erase(v.Begin())
	t.Logf("v: %v", v.String())
	assert.Equal(t, 2, iter.Value())

	v.PushBack(5)
	v.PushBack(6)
	// [2 3 4 5 6]
	iter = v.EraseRange(v.Begin().Next(), v.Begin().Next().Next().Next())
	// [2 5 6]
	assert.Equal(t, 5, iter.Value())
	assert.Equal(t, "[2 5 6]", v.String())

	iter = v.Begin()
	v.Insert(iter, 7)
	// [7 2 5 6]
	assert.Equal(t, "[7 2 5 6]", v.String())
	assert.Equal(t, 7, iter.Value())

	assert.True(t, v.Begin().Equal(v.Begin().Clone()))
	assert.False(t, v.Begin().Equal(v.Last()))
}
