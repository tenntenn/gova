package gova_test

import (
	"gova"
	"testing"
)

func TestEmailValid(t *testing.T) {
	v := struct {
		Email string `email:"-"`
	}{
		"hoge@hoge.com",
	}

	if err := gova.Validate(v); err != nil {
		t.Error(err)
	}
}

func TestEmailInValid(t *testing.T) {
	v := struct {
		Email string `email:"-"`
	}{
		"fuga",
	}

	if err := gova.Validate(v); err == nil {
		t.Error("email validator doesn't work!")
	}
}
