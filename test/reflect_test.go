package gova_test

import (
	"gova"
	"testing"
)

type Form struct {
	Alpha string `pattern:"[A-Z][a-z]+" json:"firstname"`
	
}

func TestStructPointerResolutionValidation (){
  tmp := &Form{"alpha"}
	validationErr := gova.Validate(tmp)
}
