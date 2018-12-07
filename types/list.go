package types

import (
	"fmt"
	"sync"
)

type List struct {
	head Any
	tail *List
}

var _empty *List
var once sync.Once

func empty() *List {
	once.Do(func() {
		_empty = &List{}
	})
	return _empty
}

func (l *List) Repr() string {
	fmt.Printf("LIST IS: %+v\n", &l)
	if l == empty() {
		return "()"
	}
	s := fmt.Sprintf("(%s %s)", l.head.Repr(), l.tail.Repr())
	return s
}

func NewList(args ...Any) *List {
	switch len(args) {
	case 0:
		return empty()
	case 1:
		return &List{args[0], empty()}
	default:
		l := &List{args[0], empty()}
		for i := 1; i < len(args); i++ {
			l = l.Cons(args[i])
		}
		return l
	}
}

func (l *List) Cons(a Any) *List {
	if a == empty() {
		return l
	}
	return &List{a, l}
}

func (l *List) Head() Any {
	return l.head
}

func (l *List) Tail() *List {
	if l.IsEmpty() {
		return l
	}
	return l.tail
}

func (l *List) IsEmpty() bool {
	return l == empty()
}
