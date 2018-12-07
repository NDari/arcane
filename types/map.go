package types

type Map struct {
	val map[Key]Any
}

func NewMap() *Map {
	return &Map{make(map[Key]Any)}
}

func (e *Map) Get(k Key) Any {
	v, ok := e.val[k]
	if ok {
		return v
	}
	return empty()
}

func (e *Map) Set(k Key, v Any) Any {
	v, ok := e.val[k]
	if ok {
		return v
	}
	return empty()
}
