package arcane

type Str struct {
	val string
}

func (s *Str) String() string {
	return "Str: " + s.val
}
