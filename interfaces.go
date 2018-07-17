package arcane

type Seq interface {
	Head() *Any
	Tail() *Seq
	Cons(Any) *Seq
}

type Pair interface {
	Car() *Any
	Cdr() *Any
}

type Atom interface {
	atomic()
}

func isAtom(a Any) bool {
	_, ok := a.(Atom)
	return ok
}

type Any interface {
	arcaneType()
}
