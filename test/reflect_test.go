package gova_test

import (
	"gova"
	"testing"
)

type Form struct {
	Alpha string `pattern:"[A-Z][a-z]+" json:"firstname"`
}

func TestStructPointerResolutionValidation(t *testing.T) {
	// this test should not cause panic
	tmp := &Form{"alpha"}
	gova.Validate(tmp)
}
