package main

import "fmt"

// type Unit struct{}

// func (u *Unit) ArcaneType() {}

type Sym struct {
	val string
}

func (s *Sym) String() string {
	return fmt.Sprintf("Sym: %s", s.val)
}

func (s *Sym) Atomic() {}

type Fn struct {
	ns   *Sym
	name *Sym
	call func(*Namespace, ...*Any) (*Any, error)
}

func (f *Fn) String() string {
	return fmt.Sprintf("#%s/%s", f.ns.String(), f.name.String())
}

func (f *Fn) Atomic() {}

type List struct {
	head *Any
	tail *List
}

func (l *List) String() string {
	if l.head == nil {
		return "()"
	}
	s := fmt.Sprintf("(%s", *l.head)
	for t := l.tail; t != nil; t = t.tail {
		if t.head == nil {
			break
		}
		s = fmt.Sprintf("%s, %s", s, *t.head)
	}
	s += ")"
	return s
}

type Vec struct {
	vals []*Any
}

func (v *Vec) String() string {
	if v == nil || len(v.vals) == 0 {
		return "[]"
	}
	s := fmt.Sprintf("[%s", *v.vals[0])
	for i := 1; i < len(v.vals); i++ {
		s = fmt.Sprintf("%s, %s", s, *v.vals[i])
	}
	s += "]"
	return s
}

type HashMap struct {
	name *Sym
	vals map[*Key]*Any
}

type Str struct {
	val string
}

func (s *Str) String() string {
	return fmt.Sprintf("Str: \"%s\"", s.val)
}

func (s *Str) Atomic() {}

type Key struct {
	val string
}

func (k *Key) String() string {
	return fmt.Sprintf("Key: :%s", k.val)
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
