package types

import (
	"github.com/pkg/errors"
)

type Atom interface {
	IsAtom()
}

func AtomHuh(a Any) bool {
	switch a.(type) {
	case Atom:
		return true
	default:
		return false
	}
}

func AsAtom(a Any) (Atom, error) {
	b, ok := a.(Atom)
	if !ok {
		return nil, errors.Errorf("unable to cast %T to Atom", a)
	}
	return b, nil
}
