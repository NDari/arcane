package types

import "fmt"

type Map struct {
	vals map[string]Any
}

func NewMap() *Map {
	return &Map{
		make(map[string]Any),
	}
}

func (m *Map) Repr() string {
	if m.IsEmpty() {
		return "{}"
	}
	s := "{"
	first := true
	for k, v := range m.vals {
		if first {
			s += fmt.Sprintf("%s: %s", k, v.Repr())
			first = false
		} else {
			s += fmt.Sprintf(", %s: %s", k, v.Repr())
		}
	}
	s += "}"
	return s
}

func (m *Map) Get(s string) Any {
	v, ok := m.vals[s]
	if ok {
		return v
	}
	return nil
}

func (m *Map) Set(k string, v Any) {
	m.vals[k] = v
	return
}

func (m *Map) IsEmpty() bool {
	return len(m.vals) == 0
}

// type iterableMap struct {
// 	*Map

// 	currentIndex I64
// }

// func (i *iterableMap) HasNext() bool {
// 	return int(i.currentIndex.Val) < len(i.vals)
// }

// func (i *iterableMap) Next() Any {
// 	i.currentIndex.Val++
// 	return i.vals[i.currentIndex.Val-1]
// }

// func (m *Map) ToIterable() Iterator {
// 	return &iterableMap{
// 		m,
// 		I64{0},
// 	}
// }
