package types

import "sync"

type List struct {
	head Any
	tail *List
}

var _emptyList *List
var once sync.Once

func emptyList() *List {
	once.Do(func() {
		_emptyList = &List{}
	})
	return _emptyList
}

func NewList(args ...Any) *List {
	switch len(args) {
	case 0:
		return emptyList()
	case 1:
		return &List{args[0], emptyList()}
	default:
		l := &List{args[len(args)-1], emptyList()}
		for i := len(args) - 1; i >= 0; i-- {
			l = l.Cons(args[i])
		}
		return l
	}
}

func (l *List) Cons(a Any) *List {
	if a == emptyList() {
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
	return l == emptyList()
}
