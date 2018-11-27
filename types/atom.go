package types

import (
	"fmt"

	"github.com/pkg/errors"
)

type Atom interface {
	Atomic()
}

func atomHuh(a Any) bool {
	_, ok := a.(Atom)
	return ok == true
}

func asAtom(a Any) (Atom, error) {
	b, ok := a.(Atom)
	if !ok {
		return nil, errors.Errorf("unable to cast %T to Atom", a)
	}
	return b, nil
}

func pr(a Atom) {
	fmt.Println(a)
}
