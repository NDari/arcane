package types

import "fmt"

type Sym struct {
	Val string
}

func (s Sym) Repr() string {
	return fmt.Sprintf(":%s", s.Val)
}
