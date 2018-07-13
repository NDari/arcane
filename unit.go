package arcane

type Unit struct{}

func (u *Unit) String() string {
	return "unit"
}

var c []int = make([]int, 3)
