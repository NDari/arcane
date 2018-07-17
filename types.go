package arcane

import "fmt"

type Unit struct{}

func (u *Unit) ArcaneType() {}

type Sym struct {
	val string
}

type Bool struct {
	val bool
}

func (b *Bool) ArcaneType() {}

func (s *Sym) String() string {
	return "Sym: " + string(s.val)
}

func (s Sym) ArcaneType() {}

func (s *Sym) Atomic() {}

type Fn struct {
	op func(*List, *Namespace) (Any, error)
}

func (f Fn) ArcaneType() {}

func (f *Fn) Atomic() {}

type List struct {
	head Any
	tail *List
}

func (l *List) ArcaneType() {}

func NewList(vals ...Any) *List {
	switch len(vals) {
	case 0:
		u := &Unit{}
		return &List{
			head: u,
			tail: nil,
		}
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

func (l *List) Cons(a Any) *List {
	if a == nil {
		return l
	}
	return &List{
		head: a,
		tail: l,
	}
}

func (l *List) Car() Any {
	return l.head
}

func (l *List) Cdr() *List {
	return l.tail
}

type Cell struct {
	left  Any
	right Any
}

func NewCell(left, right Any) *Cell {
	return &Cell{
		left,
		right,
	}
}

func (c *Cell) Car() Any {
	return c.left
}

func (c *Cell) Cdr() Any {
	return c.right
}

type Str struct {
	val string
}

func (s *Str) String() string {
	return "Str: " + string(s.val)
}

func (s *Str) ArcaneType() {}
func (s *Str) Atomic()     {}

type Key struct {
	val string
}

func (k *Key) String() string {
	return "Key: :" + string(k.val)
}

func (k *Key) ArcaneType() {}
func (k *Key) Atomic()     {}

type I64 struct {
	val int64
}

func (i *I64) String() string {
	return fmt.Sprintf("I64: %d", int64(i.val))
}

func (i *I64) ArcaneType() {}
func (i *I64) Atomic()     {}

type F64 struct {
	val float64
}

func (f *F64) String() string {
	return fmt.Sprintf("F64: %f", float64(f.val))
}

func (f *F64) ArcaneType() {}
func (f *F64) Atomic()     {}
