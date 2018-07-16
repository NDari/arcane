package arcane

type Empty struct{}

func (e *Empty) String() string {
	return "Unit"
}

func (e *Empty) arcaneType() {}
func (e *Empty) atomic()     {}
