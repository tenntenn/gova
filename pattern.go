package gova

import (
	"errors"
	"fmt"
	"regexp"
)

type PatternValidator struct {
}

func (pv *PatternValidator) validate(v interface{}, param interface{}) (err error) {
	var reg *regexp.Regexp
	switch param.(type) {
	case string:
		str, _ := param.(string)
		reg, err = regexp.Compile(str)
		if err != nil {
			return
		}
	case *regexp.Regexp:
		reg, _ = param.(*regexp.Regexp)
	default:
		return errors.New("param must be string or *regexp.Regexp")
	}

	switch v.(type) {
	case string:
		str := v.(string)
		if reg.MatchString(str) {
			return
		}
	case []byte:
		bytes := v.([]byte)
		if reg.Match(bytes) {
			return
		}
	case fmt.Stringer:
		stringer := v.(fmt.Stringer)
		str := stringer.String()
		if reg.MatchString(str) {
			return
		}
	default:
		return errors.New(fmt.Sprintf("%v cannot validate by PatternValidator", v))
	}

	return errors.New("Does not mutch given pattern")
}

func (pv *PatternValidator) Validate(v interface{}, param string) (err error) {
	return pv.validate(v, param)
}
