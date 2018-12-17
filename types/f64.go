package types

import (
	"fmt"
)

type F64 float64

func (f F64) Repr() string {
	return fmt.Sprintf("%f", f)
}
