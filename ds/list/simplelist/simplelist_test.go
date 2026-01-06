package simplelist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListBase(t *testing.T) {
	l := New[int]()
	assert.Equal(t, 0, l.Len())
	l.PushBack(1)
	assert.Equal(t, 1, l.Len())
	assert.Equal(t, 1, l.FrontNode().Value)
	assert.Equal(t, 1, l.BackNode().Value)
	l.PushFront(2)

	assert.Equal(t, 2, l.Len())
	assert.Equal(t, 2, l.FrontNode().Value)
	assert.Equal(t, "[2 1]", l.String())
	l.PushBack(3)
	l.PushBack(4)
	assert.Equal(t, "[2 1 3 4]", l.String())

	l.MoveToFront(l.FrontNode(), l.FrontNode().Next())
	assert.Equal(t, "[1 2 3 4]", l.String())
	l.MoveToBack(l.FrontNode(), l.FrontNode().Next())
	assert.Equal(t, "[1 3 4 2]", l.String())

	ret := make([]int, 0)
	l.Traversal(func(val int) bool {
		ret = append(ret, val)
		return true
	})
	assert.Equal(t, "[1 3 4 2]", fmt.Sprintf("%v", ret))
}

func TestListInsertAfter(t *testing.T) {
	l := New[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.InsertAfter(3, l.FrontNode())
	assert.Equal(t, "[1 3 2]", l.String())

	l.InsertAfter(4, l.BackNode())
	assert.Equal(t, "[1 3 2 4]", l.String())
	l.InsertAfter(5, l.FrontNode())
	assert.Equal(t, "[1 5 3 2 4]", l.String())
}

func TestListRemove(t *testing.T) {
	l := New[int]()
	for i := 1; i <= 5; i++ {
		l.PushBack(i)
	}
	assert.Equal(t, "[1 2 3 4 5]", l.String())
	l.Remove(nil, l.FrontNode())
	assert.Equal(t, "[2 3 4 5]", l.String())
	l.Remove(l.FrontNode(), l.FrontNode().Next())
	assert.Equal(t, "[2 4 5]", l.String())
}

func TestListIterator(t *testing.T) {
	l := New[int]()
	for i := 1; i <= 5; i++ {
		l.PushBack(i)
	}
	assert.Equal(t, "[1 2 3 4 5]", l.String())
	i := 1
	for iter := NewIterator(l.FrontNode()); iter.IsValid(); iter.Next() {
		assert.Equal(t, i, iter.Value())
		iter.SetValue(i * 2)
		i++
	}
	iter := NewIterator(l.FrontNode())
	assert.Equal(t, 2, iter.Value())
	assert.True(t, iter.Equal(iter.Clone()))
}
