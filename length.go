package gova

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type LengthValidator struct {
}

func (lv *LengthValidator) Validate(v interface{}, param string) (err error) {
	length, err := strconv.ParseInt(param, 10, 32)
	if err != nil {
		return err
	}

	switch v.(type) {
	case string:
		str := v.(string)
		if len(str) == int(length) {
			return
		}
	default:
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Slice, reflect.Array, reflect.Map, reflect.Chan:
			if rv.Len() == int(length) {
				return
			}
		default:
			return errors.New(fmt.Sprintf("%v has not length", v))
		}
	}

	return errors.New("Does not mutch given length")
}
