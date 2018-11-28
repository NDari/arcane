package types

type Cons struct {
	first  Any
	second Any
}

func (self Cons) IsPair() {}
