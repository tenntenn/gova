package gova_test

import (
	"gova"
	"testing"
)

func TestLengthValid(t *testing.T) {
	v := struct {
		Str string `length:"10"`
	}{
		"1234567890",
	}

	if err := gova.Validate(v); err != nil {
		t.Error(err)
	}
}

func TestLengthInValid(t *testing.T) {
	v := struct {
		Str string `length:"10"`
	}{
		"1234",
	}

	if err := gova.Validate(v); err == nil {
		t.Error("length validator doesn't work!")
	}
}
