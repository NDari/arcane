package types

import (
	"testing"
)

type testAtom struct{}

func (tt testAtom) IsAtom() {}

func TestIsAtom(t *testing.T) {
	var tt testAtom

	if !AtomHuh(tt) {
		t.Error("atomic type failed 'atomHuh'")
	}
	if AtomHuh(10) {
		t.Error("non-atomic type passed 'atomHuh'")
	}
}

func TestAsAtom(t *testing.T) {
	var tt testAtom
	_, err := AsAtom(tt)
	if err != nil {
		t.Errorf("atomic type could not be converted to an atom: %v", err)
	}
	bad, err := AsAtom("thing")
	if err == nil {
		t.Errorf("non-atomic type %T converted to atom %v without error", "thing", bad)
	}
}

func TestAtomicTypes(t *testing.T) {
	tt := &Str{"test"}
	if !AtomHuh(tt) {
		t.Error("Str is not atomic")
	}
}
