package types

type Cell struct {
	first  Any
	second Any
}

func (c Cell) IsPair() {}
