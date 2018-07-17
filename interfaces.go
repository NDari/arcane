package arcane

type Seq interface {
	Head() Any
	Tail() *Seq
	Cons(Any) *Seq
}

type Pair interface {
	Car() Any
	Cdr() *Any
}

type Atom interface {
	Atomic()
}

type Any interface{}
