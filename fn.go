package arcane

type Fn struct {
	name Sym
	args *List
	call func(*List) *List
}

func (f *Fn) arcaneType() {}
func (f *Fn) atomic()     {}
