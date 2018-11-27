package types

import (
	"testing"
)

type testAtom struct{}

func (tt testAtom) Atomic() {}

func TestIsAtom(t *testing.T) {
	var tt testAtom

	if !atomHuh(tt) {
		t.Error("atomic type failed 'atomHuh'")
	}
	if atomHuh(10) {
		t.Error("non-atomic type passed 'atomHuh'")
	}
}

func TestAsAtom(t *testing.T) {
	var tt testAtom
	pr(tt)
	_, err := asAtom(tt)
	if err != nil {
		t.Errorf("atomic type could not be converted to an atom: %v", err)
	}
	bad, err := asAtom("thing")
	if err == nil {
		t.Errorf("non-atomic type %T converted to atom %v without error", "thing", bad)
	}
}
