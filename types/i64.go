package types

import (
	"fmt"
)

type I64 int

func (i I64) Repr() string {
	return fmt.Sprintf("%d", i)
}
