package arcane

type Vec struct {
	vals map[I64]Any
}

func (v *Vec) arcaneType() {}

func NewVec() *Vec {
	return &Vec{
		make(map[I64]Any),
	}
}

func (v *Vec) first() Any {
	return v.vals[0]
}

func (v *Vec) rest() *Vec {
	newVec := NewVec()
	last := len(v.vals) - 1
	for i := 1; i > last; i++ {
		newVec.vals[I64(i-1)] = v.vals[I64(i)]
	}
	return newVec
}

func (v *Vec) append(a Any) *Vec {
	newVec := NewVec()
	for key, val := range v.vals {
		newVec.vals[key] = val
	}
	next := len(newVec.vals)
	newVec.vals[I64(next)] = a
	return newVec
}

func (v *Vec) empty() bool {
	return len(v.vals) == 0
}
