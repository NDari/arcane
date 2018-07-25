package main

import "fmt"

type Sym struct {
	val string
}

func (s *Sym) String() string {
	return fmt.Sprintf("%s : Sym", s.val)
}

func (s *Sym) Atomic() {}

type Fn struct {
	ns   *Sym
	name *Sym
	call func(*Namespace, *Any) (*Any, error)
}

func (f *Fn) String() string {
	return fmt.Sprintf("#(%s/%s)", f.ns.val, f.name.val)
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
	s += ") : List"
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
	s += "] : Vec"
	return s
}

type HashMap struct {
	vals map[*Key]*Any
}

func (h *HashMap) String() string {
	if h == nil || len(h.vals) == 0 {
		return "{}"
	}
	s := "{"
	for k, v := range h.vals {
		if len(s) > 1 {
			s += ","
		}
		s += fmt.Sprintf("%v => %v", k, *v)
	}
	s += "} : HashMap"
	return s
}

type Str struct {
	val string
}

func (s *Str) String() string {
	return fmt.Sprintf("\"%s\" : Str", s.val)
}

func (s *Str) Atomic() {}

type Key struct {
	val string
}

func (k *Key) String() string {
	return fmt.Sprintf("%s : Key", k.val)
}

func (k *Key) Atomic() {}

type I64 struct {
	val int64
}

func (i *I64) String() string {
	return fmt.Sprintf("%d : I64", i.val)
}

func (i *I64) Atomic() {}

type F64 struct {
	val float64
}

func (f *F64) String() string {
	return fmt.Sprintf("%f : F64", f.val)
}

func (f *F64) Atomic() {}
