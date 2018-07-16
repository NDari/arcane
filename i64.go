package arcane

import (
	"fmt"
)

type I64 int64

func (i *I64) String() string {
	return fmt.Sprintf("I64: %d", i)
}

func (i I64) arcaneType() {}
func (i I64) atomic()     {}
