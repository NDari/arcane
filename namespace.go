package main

import (
	"github.com/NDari/arcane/types"
)

type Namespace struct {
	upper *Namespace
	name  types.Sym
	vals  *types.Map
}

func NewNS(name types.Sym, upper *Namespace) *Namespace {
	m := types.NewMap()
	return &Namespace{
		upper,
		name,
		m,
	}
}

func TopLevel() *Namespace {
	topNameSpace := types.Sym{
		Val: "arcane",
	}

	ns := NewNS(topNameSpace, nil)
	return ns
}
