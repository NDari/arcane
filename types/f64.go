package types

import (
	"fmt"
)

type F64 struct {
	Val float64
}

func (f F64) Repr() string {
	return fmt.Sprintf("F64: %f", f.Val)
}
