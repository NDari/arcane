package main

import (
	"github.com/NDari/arcane/types"
)

type Namespace struct {
	*types.Map
	upper *Namespace
	name  types.Sym
}

func NewNS(name types.Sym, upper *Namespace) *Namespace {
	return &Namespace{types.NewMap(), upper, name}
}

func BaseNS() *Namespace {
	return NewNS(types.Sym("Base"), nil)
}
