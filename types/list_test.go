package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	e := NewList()
	assert.Equal(t, e, empty(), "new list should be empty")

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
}
