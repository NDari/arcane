package arcane

import (
	"fmt"
)

type I64 struct {
	val int64
}

func (i *I64) String() string {
	return fmt.Sprintf("I64: %d", i.val)
}
