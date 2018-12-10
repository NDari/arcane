package types

import (
	"fmt"
)

type List struct {
	vals []Any
}

func (l *List) Repr() string {
	if l.IsEmpty() {
		return "()"
	}
	s := fmt.Sprintf("(%s", l.vals[0].Repr())
	for i := 1; i < len(l.vals); i++ {
		s += fmt.Sprintf(", %s", l.vals[i].Repr())
	}
	s += ")"
	return s
}

func NewList(args ...Any) *List {
	switch len(args) {
	case 0:
		return &List{
			make([]Any, 0),
		}
	default:
		return &List{args}
	}
}

func (l *List) Append(a ...Any) {
	if len(a) == 0 {
		return
	}
	l.vals = append(l.vals, a...)
	return
}

func (l *List) IsEmpty() bool {
	return len(l.vals) == 0
}

type iterableList struct {
	*List

	currentIndex I64
}

func (i *iterableList) HasNext() bool {
	return int(i.currentIndex.Val) < len(i.vals)
}

func (i *iterableList) Next() Any {
	i.currentIndex.Val++
	return i.vals[i.currentIndex.Val-1]
}

func (l *List) ToIterable() Iterator {
	return &iterableList{
		l,
		I64{0},
	}
}
