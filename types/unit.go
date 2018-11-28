package types

type Unit struct{}

func (self Unit) IsAtom() {}
func (self Unit) IsPair() {}

func UnitHuh(a Any) bool {
	switch a.(type) {
	case Unit:
		return true
	default:
		return false
	}
}
