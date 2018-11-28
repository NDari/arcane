package types

import (
	"fmt"

	"github.com/pkg/errors"
)

type Str struct {
	val string
}

func (self Str) IsAtom() {}

func (self *Str) String() string {
	return fmt.Sprintf("Str: %s", self.val)
}

func strHuh(a Any) bool {
	switch a.(type) {
	case *Str:
		return true
	default:
		return false
	}
}

func AsStr(a Any) (*Str, error) {
	s, ok := a.(*Str)
	if !ok {
		return nil, errors.Errorf("could not cast %T to Str", a)
	}
	return s, nil
}
