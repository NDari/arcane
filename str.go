package arcane

type Str string

func (s *Str) String() string {
	return "Str: " + string(*s)
}

func (s Str) arcaneType() {}
func (s Str) atomic()     {}
