package types

import (
	"testing"
)

type testPair struct{}

func (tt testPair) IsPair() {}

func TestIsPair(t *testing.T) {
	var tt testPair

	if !PairHuh(tt) {
		t.Error("type implementing IsPair type failed 'PairHuh'")
	}
	if PairHuh(10) {
		t.Error("type not implementing IsPair passed 'PairHuh'")
	}
}

func TestAsPair(t *testing.T) {
	var tt testPair
	_, err := AsPair(tt)
	if err != nil {
		t.Errorf("type implementing IsPair() could not be converted to Pair: %v", err)
	}
	bad, err := AsPair("thing")
	if err == nil {
		t.Errorf("type %T converted to Pair %v without implementing IsPair()", "thing", bad)
	}
}
