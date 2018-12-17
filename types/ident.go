package types

import "fmt"

type Ident string

func (i Ident) Repr() string {
	return fmt.Sprintf("%s", i)
}
