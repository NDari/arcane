package main

import (
	"github.com/NDari/arcane/types"
)

type Namespace struct {
	upper *Namespace
	name  types.Key
	vals  *types.Map
}

func NewNS(name types.Key, upper *Namespace) *Namespace {
	m := types.NewMap()
	return &Namespace{
		upper,
		name,
		m,
	}
}

func TopLevel() *Namespace {
	topNameSpace := types.Key{
		"arcane",
	}

	ns := NewNS(topNameSpace, nil)
	return ns
}
