package gova_test

import (
	"gova"
	"testing"
)

func TestPatternValid(t *testing.T) {
	v := struct {
		Num string `pattern:"[0-9]+"`
	}{
		"11111",
	}

	if err := gova.Validate(v); err != nil {
		t.Error(err)
	}
}

func TestPatternInValid(t *testing.T) {
	v := struct {
		Num string `pattern:"[0-9]+"`
	}{
		"fuga",
	}

	if err := gova.Validate(v); err == nil {
		t.Error("pattern validator doesn't work!")
	}
}
