package arcane

type Sym string

func (s *Sym) String() string {
	return "Sym: " + string(*s)
}

func (s *Sym) arcaneType() {}
func (s *Sym) atomic()     {}
