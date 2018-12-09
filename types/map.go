package types

import "fmt"

type Map struct {
	vals []*List
}

func NewMap() *Map {
	return &Map{
		make([]*List, 0),
	}
}

func (m *Map) Repr() string {
	if m.IsEmpty() {
		return "{}"
	}
	s := fmt.Sprintf("{%s %s", m.vals[0].vals[0].Repr(), m.vals[0].vals[1].Repr())
	for i := 1; i < len(m.vals); i++ {
		s += fmt.Sprintf(", %s %s", m.vals[i].vals[0].Repr(), m.vals[i].vals[1].Repr())
	}
	s += "}"
	return s
}

func (m *Map) Get(k Key) Any {
	for i := range m.vals {
		if m.vals[i].vals[0] == k {
			return m.vals[i].vals[1]
		}
	}
	return nil
}

func (m *Map) Set(l *List) {
	if l.IsEmpty() {
		return
	}

	for i := range m.vals {
		if m.vals[i].vals[0] == l.vals[0] {
			m.vals[i].vals[1] = l.vals[1]
			return
		}
	}
	m.vals = append(m.vals, l)
	return
}

func (m *Map) IsEmpty() bool {
	return len(m.vals) == 0
}

type iterableMap struct {
	*Map

	currentIndex I64
}

func (i *iterableMap) HasNext() bool {
	return int(i.currentIndex.Val) >= len(i.vals)
}

func (i *iterableMap) Next() Any {
	i.currentIndex.Val++
	return i.vals[i.currentIndex.Val-1]
}

func (m *Map) ToIterable() Iterator {
	return &iterableMap{
		m,
		I64{0},
	}
}
