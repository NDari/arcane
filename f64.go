package arcane

import (
	"fmt"
)

type F64 float64

func (f *F64) String() string {
	return fmt.Sprintf("F64: %f", float64(*f))
}

func (f *F64) arcaneType() {}
func (f *F64) atomic()     {}
