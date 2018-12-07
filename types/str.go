package types

import (
	"fmt"
)

type Str struct {
	Val string
}

func (s Str) Repr() string {
	return fmt.Sprintf("%s", s.Val)
}
