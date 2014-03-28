package gova

import (
	"errors"
	"reflect"
	"sync"
)

type ValidatorMap struct {
	m  map[string]Validator
	mu sync.RWMutex
}

func (vm *ValidatorMap) Add(key string, validator Validator) error {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	if _, ok := vm.m[key]; ok {
		return errors.New(key + " is already added.")
	}
	vm.m[key] = validator

	return nil
}

func (vm *ValidatorMap) Remove(key string) {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	delete(vm.m, key)
}

var Validators *ValidatorMap = &ValidatorMap{m: make(map[string]Validator)}

type Validator interface {
	Validate(v interface{}, param string) error
}

func Validate(v interface{}) error {

	val, ok := v.(reflect.Value)
	if !ok {
		val = reflect.ValueOf(v)
	}
	t := val.Type()

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		t = val.Type()
	}

	if val.Kind() != reflect.Struct {
		return errors.New("v must be struct or struct pointer.")
	}

	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		if t.Field(i).Name == "_" {
			continue
		}

		tag := t.Field(i).Tag
		if string(tag) != "" {
			Validators.mu.RLock()
			for key, validator := range Validators.m {
				if param := tag.Get(key); param != "" {
					if err := validator.Validate(f.Interface(), param); err != nil {
						return err
					}
				}
			}
			Validators.mu.RUnlock()
		}

		if f.Kind() == reflect.Struct {
			if err := Validate(f); err != nil {
				return err
			}
		}

	}

	return nil
}
