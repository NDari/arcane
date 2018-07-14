package arcane

type Seq interface {
	first() Any
	rest() *Seq
	append(Any) *Seq
	empty() bool
}
