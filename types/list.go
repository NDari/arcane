package types

import (
	"fmt"
)

type List struct {
	vals  map[int]Any
	index int
}

func (l *List) Repr() string {
	if l.IsEmpty() {
		return "[]"
	}
	s := fmt.Sprintf("[%s", l.vals[0].Repr())
	for i := 1; i < l.index; i++ {
		s += fmt.Sprintf(", %s", l.vals[i].Repr())
	}
	s += "]"
	return s
}

func NewList(args ...Any) *List {
	switch len(args) {
	case 0:
		return &List{
			make(map[int]Any),
			0,
		}
	default:
		l := NewList()
		for i := range args {
			l.vals[i] = args[i]
		}
		l.index = len(args)
		return l
	}
}

func (l *List) Append(a ...Any) {
	if len(a) == 0 {
		return
	}
	for i := range a {
		l.vals[l.index] = a[i]
		l.index++
	}
	return
}

func (l *List) IsEmpty() bool {
	return len(l.vals) == 0
}

type iterableList struct {
	*List

	currentIndex int
}

func (i *iterableList) HasNext() bool {
	return i.currentIndex > 0
}

func (i *iterableList) Next() Any {
	i.currentIndex--
	return i.vals[i.currentIndex]
}

func (l *List) ToIterable() Iterator {
	return &iterableList{
		l,
		l.index,
	}
}
