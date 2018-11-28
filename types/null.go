package types

type null struct{}

var Nil null

func (self null) IsAtom() {}
func (self null) IsPair() {}

func NullHuh(a Any) bool {
	switch a.(type) {
	case null:
		return true
	default:
		return false
	}
}
