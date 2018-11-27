package types

import "testing"

func TestIsStr(t *testing.T) {
	tt := &Str{"test"}
	if !strHuh(tt) {
		t.Error("Str type failed 'strHuh'")
	}
	if strHuh(10) {
		t.Error("non-Str type passed 'strHuh'")
	}
}

func TestAsStr(t *testing.T) {
	tt := &Str{"test"}
	_, err := asStr(tt)
	if err != nil {
		t.Error("Str type could not be converted to Str")
	}
	bad, err := asStr(12.0)
	if err == nil {
		t.Errorf("non-str type %T converted to Str %v without error", 12.0, bad)
	}
}
