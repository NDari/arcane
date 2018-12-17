package types

import (
	"fmt"
)

type Str string

func (s Str) Repr() string {
	return fmt.Sprintf("%s", s)
}
