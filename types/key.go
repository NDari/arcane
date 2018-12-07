package types

import "fmt"

type Key struct {
	Val string
}

func (k Key) Repr() string {
	return fmt.Sprintf(":%s", k.Val)
}
