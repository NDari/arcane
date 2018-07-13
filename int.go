package arcane

type Int struct {
	val int64
}

func (i *Int) String() string {
	return "Int: " + string(i.val)
}
