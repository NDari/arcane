package arcane

import "fmt"

type Sym string

func (s *Sym) String() string {
	return "Sym: " + string(*s)
}

func (s *Sym) arcaneType() {}
func (s *Sym) atomic()     {}

type Fn struct {
	name Sym
	call func(*List) Any
}

func (f *Fn) arcaneType() {}
func (f *Fn) atomic()     {}

type List struct {
	head *Any
	tail *List
}

func NewList(vals ...*Any) *List {
	switch len(vals) {
	case 0:
		return nil
	case 1:
		return &List{
			head: vals[0],
			tail: nil,
		}
	default:
		l := &List{
			head: vals[len(vals)-1],
			tail: nil,
		}
		for i := len(vals) - 2; i >= 0; i-- {
			l = l.Cons(vals[i])
		}
		return l
	}
}

func (l *List) Cons(a *Any) *List {
	if a == nil {
		return l
	}
	return &List{
		head: a,
		tail: l,
	}
}

func (l *List) Car() *Any {
	return l.head
}

func (l *List) Cdr() *List {
	return l.tail
}

type ConsCell struct {
	left  *Any
	right *Any
}

func NewCell(left, right *Any) *ConsCell {
	return &ConsCell{
		left,
		right,
	}
}

func (c *ConsCell) Car() *Any {
	return c.left
}

func (c *ConsCell) Cdr() *Any {
	return c.right
}

type Str string

func (s *Str) String() string {
	return "Str: " + string(*s)
}

func (s Str) arcaneType() {}
func (s Str) atomic()     {}

type Key string

func (k *Key) String() string {
	return "Key: :" + string(*k)
}

func (k *Key) arcaneType() {}
func (k *Key) atomic()     {}

type I64 int64

func (i *I64) String() string {
	return fmt.Sprintf("I64: %d", i)
}

func (i I64) arcaneType() {}
func (i I64) atomic()     {}

type F64 float64

func (f *F64) String() string {
	return fmt.Sprintf("F64: %f", float64(*f))
}

func (f F64) arcaneType() {}
func (f F64) atomic()     {}
