package types

import (
	"fmt"
	"sync"
)

type List struct {
	vals []Any
}

var _empty *List
var once sync.Once

func empty() *List {
	once.Do(func() {
		_empty = &List{
			vals: make([]Any, 0),
		}
	})
	return _empty
}

func (l *List) Repr() string {
	if l.IsEmpty() {
		return "[]"
	}
	s := fmt.Sprintf("[%s", l.vals[0].Repr())
	for i := 1; i < len(l.vals); i++ {
		s += fmt.Sprintf(", %s", l.vals[i].Repr())
	}
	s += "]"
	return s
}

func NewList(args ...Any) *List {
	switch len(args) {
	case 0:
		return empty()
	default:
		return &List{args}
	}
}

func (l *List) IsEmpty() bool {
	return len(l.vals) == 0
}

func (l *List) Append(a ...Any) {
	if len(a) == 0 {
		return
	}
	l.vals = append(l.vals, a...)
	return
}

type IterableList struct {
	*List

	currentIndex I64
}

func (i *IterableList) HasNext() bool {
	return int(i.currentIndex.Val) >= len(i.vals)
}

func (i *IterableList) Next() Any {
	i.currentIndex.Val++
	return i.vals[i.currentIndex.Val-1]
}

func (l *List) ToIterable() Iterator {
	return &IterableList{
		l,
		I64{0},
	}
}
