package types

import "fmt"

type Sym string

func (s Sym) Repr() string {
	return fmt.Sprintf(":%s", s)
}
