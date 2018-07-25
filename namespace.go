package main

import "fmt"

type Namespace struct {
	upper *Namespace
	name  *Sym
	vals  map[Sym]*Any
}

func (n *Namespace) String() string {
	var s string
	if n.name.val == "arcane" {
		s = fmt.Sprintf("Namespace:<%s>, Upper:<>\nSymbol table {\n", n.name.val)
	} else {
		s = fmt.Sprintf("Namespace:<%s>, Upper:<%s>\nSymbol table {\n", n.name.val, n.upper.name.val)
	}

	for k, v := range n.vals {
		s += fmt.Sprintf("\t%s => %s\n", k, *v)
	}
	s += "}"
	return s
}

func NewNS(name *Sym, upper *Namespace) *Namespace {
	m := make(map[Sym]*Any)
	return &Namespace{
		upper,
		name,
		m,
	}
}

func TopLevel() *Namespace {
	m := make(map[Sym]*Any)
	topNameSpace := Sym{
		"arcane",
	}

	ns := &Namespace{
		nil,
		&topNameSpace,
		m,
	}

	eval := Sym{
		"eval",
	}

	var evalFn Any = &Fn{
		ns:   &topNameSpace,
		name: &eval,
		call: Eval,
	}

	ns.vals[eval] = &evalFn

	cons := Sym{
		"cons",
	}

	var consFn Any = &Fn{
		ns:   &topNameSpace,
		name: &cons,
		call: Cons,
	}

	ns.vals[cons] = &consFn

	return ns
}
