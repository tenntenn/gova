package gova

import (
	"regexp"
)

var emailRegexp = regexp.MustCompile(`^[\w!#$%&'*+/=?^_{}\\|~-]+([\w!#$%&'*+/=?^_{}\\|~\.-]+)*@([\w][\w-]*\.)+[\w][\w-]*$`)

type EmailValidator struct {
	PatternValidator
}

func (ev *EmailValidator) Validate(v interface{}, param string) (err error) {
	return ev.validate(v, emailRegexp)
}
