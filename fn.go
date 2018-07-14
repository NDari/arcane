package arcane

type Fn struct {
	name Sym
	args Vec
	call func(Vec) Vec
}

func (f *Fn) arcaneType() {}
func (f *Fn) atomic()     {}
