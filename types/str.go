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

func (s Str) Atomic() {}

func strHuh(a Any) bool {
	_, ok := a.(*Str)
	return ok == true
}

func asStr(a Any) (*Str, error) {
	s, ok := a.(*Str)
	if !ok {
		return nil, errors.Errorf("could not cast %T to Str", a)
	}
	return s, nil
}
