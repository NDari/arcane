package types

import "sync"

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

func NewList(args ...Any) *List {
	switch len(args) {
	case 0:
		return empty()
	case 1:
		return &List{args[0], empty()}
	default:
		l := &List{args[len(args)-1], empty()}
		for i := len(args) - 2; i >= 0; i-- {
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
