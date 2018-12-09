package types

type Map struct {
	vals []*List
}

func NewMap() *Map {
	return &Map{
		make([]*List, 0),
	}
}

func (m *Map) Get(k Key) Any {
	for i := range m.vals {
		if m.vals[i].vals[0] == k {
			return m.vals[i].vals[1]
		}
	}
	return nil
}

func (m *Map) Set(k Key, v Any) {
	for i := range m.vals {
		if m.vals[i].vals[0] == k {
			m.vals[i].vals[1] = v
			return
		}
	}
	l := NewList(k, v)
	m.vals = append(m.vals, l)
	return
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
