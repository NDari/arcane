package arcane

type Unit struct{}

func (u *Unit) String() string {
	return "Unit"
}

func (u *Unit) arcaneType() {}
func (u *Unit) atomic()     {}
