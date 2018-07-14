package arcane

type Atom interface {
	atomic()
}

func isAtom(a Any) bool {
	_, ok := a.(Atom)
	return ok
}
