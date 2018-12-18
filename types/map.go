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
	if m.IsEmpty() {
		return nil
	}
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

type iterableMap struct {
	self         *Map
	keyMap       map[int]string
	currentIndex int
}

func (i *iterableMap) HasNext() bool {
	return i.currentIndex > 0
}

func (i *iterableMap) Next() Any {
	i.currentIndex--
	return i.self.vals[i.keyMap[i.currentIndex]]
}

func (m *Map) ToIterable() Iterator {
	keyMap := make(map[int]string)
	idx := 0
	for k, _ := range m.vals {
		keyMap[idx] = k
		idx++
	}
	return &iterableMap{
		m,
		keyMap,
		len(keyMap),
	}
}
