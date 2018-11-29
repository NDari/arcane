package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	a := Str{"Hello"}
	b := Str{"goodbye"}
	l := NewList(a, b)

	assert.Equal(t, l.Head(), a, "wrong head")
	assert.Equal(t, l.Tail(), NewList(b), "wrong tail")
}

func TestCons(t *testing.T) {
	a := Str{"Hello"}
	b := Str{"goodbye"}
	l := NewList(a, b)
	l2 := l.Cons(b)

	assert.Equal(t, l2.Head(), b, "wrong head")
	assert.Equal(t, l2.Tail(), l, "wrong tail")
}
