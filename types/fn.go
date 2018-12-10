package types

import "fmt"

type Fn struct {
	Name   Str
	DocStr Str
	Body   *List
}

func (f *Fn) Repr() string {
	return fmt.Sprintf("#%s", f.Name.Repr())
}

func (f *Fn) Doc() string {
	return fmt.Sprintf("#%s\n\n%s", f.Name.Repr(), f.DocStr.Repr())
}
