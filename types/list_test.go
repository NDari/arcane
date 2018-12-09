package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	e := NewList()
	assert.Equal(t, 0, len(e.vals), "new list should be empty")

	a := Str{"Hello"}
	b := Str{"goodbye"}
	l := NewList(a, b)

	assert.Equal(t, l.vals[0], a, "wrong head")
	assert.Equal(t, l.vals[1], b, "wrong tail")
}

func TestAppend(t *testing.T) {
	a := Str{"Hello"}
	b := Str{"goodbye"}
	l := NewList(a)
	l.Append(b)

	assert.Equal(t, l.vals[0], a, "wrong head")
	assert.Equal(t, l.vals[1], b, "wrong tail")

	l2 := NewList(b)
	l3 := NewList(l, l2)
	assert.Equal(t, l3.vals[0], l, "wrong head")
	assert.Equal(t, l3.vals[1], l2, "wrong tail")
}

func TestIsEmpty(t *testing.T) {
	e := NewList()
	assert.True(t, e.IsEmpty(), "new list should be empty")

	a := Str{"Item"}
	e.Append(a)
	assert.True(t, !e.IsEmpty(), "list should not be empty")
}
