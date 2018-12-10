package types

import "fmt"

type Ident struct {
	Val string
}

func (i Ident) Repr() string {
	return fmt.Sprintf("%s", i.Val)
}
