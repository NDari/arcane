package types

import (
	"fmt"
)

type I64 struct {
	Val int64
}

func (i I64) Repr() string {
	return fmt.Sprintf("%d", i.Val)
}
