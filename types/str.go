package types

import (
	"fmt"
)

type Str struct {
	val string
}

func (self *Str) String() string {
	return fmt.Sprintf("Str: %s", self.val)
}
