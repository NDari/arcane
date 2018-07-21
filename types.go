package main

import "fmt"

// type Unit struct{}

// func (u *Unit) ArcaneType() {}

type Sym struct {
	val string
}

func (s *Sym) String() string {
	return string(s.val)
}

func (s *Sym) Atomic() {}

type Fn struct {
	ns   *Sym
	name *Sym
	call func(*Namespace, ...*Any) (*Any, error)
}

func (f *Fn) String() string {
	return "#" + f.ns.String() + "/" + f.name.String()

}

func (f *Fn) Atomic() {}

type List struct {
	head *Any
	tail *List
}

type Vec struct {
	vals []*Any
}

type HashMap struct {
	name *Sym
	vals map[*Key]*Any
}

type Str struct {
	val string
}

func (s *Str) String() string {
	return "\"" + string(s.val) + "\""
}

func (s *Str) Atomic() {}

type Key struct {
	val string
}

func (k *Key) String() string {
	return "Key: :" + string(k.val)
}

func (k *Key) Atomic() {}

type I64 struct {
	val int64
}

func (i *I64) String() string {
	return fmt.Sprintf("I64: %d", i.val)
}

func (i *I64) Atomic() {}

type F64 struct {
	val float64
}

func (f *F64) String() string {
	return fmt.Sprintf("F64: %f", f.val)
}

func (f *F64) Atomic() {}
