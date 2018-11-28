package types

import (
	"github.com/pkg/errors"
)

type Pair interface {
	IsPair()
}

func PairHuh(a Any) bool {
	switch a.(type) {
	case Pair:
		return true
	default:
		return false
	}
}

func AsPair(a Any) (Pair, error) {
	b, ok := a.(Pair)
	if !ok {
		return nil, errors.Errorf("unable to cast %T to Pair", a)
	}
	return b, nil
}

func Head(a Any) Any {
	if !PairHuh(a) {
		return Nil
	}
	switch v := a.(type) {
	case Cons:
		return v.first
	default:
		return Nil
	}
}

func Tail(a Any) Any {
	if !PairHuh(a) {
		return Nil
	}
	switch v := a.(type) {
	case Cons:
		return v.second
	default:
		return Nil
	}
}
