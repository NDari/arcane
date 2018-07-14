package arcane

import (
	"fmt"
)

type F64 struct {
	val float64
}

func (f *F64) String() string {
	return fmt.Sprintf("F64: %f", n.val)
}
