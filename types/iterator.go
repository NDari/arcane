package types

type Iterable interface {
	ToIterable() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() Any
}
