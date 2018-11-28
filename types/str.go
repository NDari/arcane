package types

import (
	"fmt"

	"github.com/pkg/errors"
)

type Str struct {
	val string
}

func (s *Str) String() string {
	return fmt.Sprintf("Str: %s", s.val)
}

func (s Str) IsAtom() {}

func strHuh(a Any) bool {
	switch a.(type) {
	case *Str:
		return true
	default:
		return false
	}
}

func asStr(a Any) (*Str, error) {
	s, ok := a.(*Str)
	if !ok {
		return nil, errors.Errorf("could not cast %T to Str", a)
	}
	return s, nil
}
