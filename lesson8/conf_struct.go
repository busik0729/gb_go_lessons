package lesson8

import (
	"errors"
	"reflect"
	"strings"
)

type Conf struct {
	Host     string `env:"HOST" validate:"require"`
	Port     int    `env:"PORT" validate:"require"`
	Protocol string `env:"PROTOCOL" validateFn:"ValidateProtocol"`
}

func (c Conf) Validate() (bool, error) {
	t := reflect.TypeOf(c)
	f := reflect.ValueOf(&c)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get("validate")
		tagValidFn := field.Tag.Get("validateFn")
		d := f.Elem().FieldByName(field.Name)
		if tag == "require" {
			if d.Kind() == reflect.String && len(d.String()) < 1 {
				return false, errors.New(strings.Join([]string{field.Name, "is require"}, " "))
			}
			if d.Kind() == reflect.Int && d.Int() <= 0 {
				return false, errors.New(strings.Join([]string{field.Name, "is require"}, " "))
			}
		}

		if len(tagValidFn) != 0 {
			res := f.MethodByName(tagValidFn).Call([]reflect.Value{})

			if res[0].Interface() != nil {
				return false, res[0].Interface().(error)
			}
		}
	}

	return true, nil
}

func (c Conf) ValidateProtocol() error {
	var allowProtocol = []string{"http", "https"}

	if len(c.Protocol) < 2 {
		return errors.New("Config error: Protocol cannot be less 2 symbols")
	}

	inArray := 0
	for i := range allowProtocol {
		if ok := allowProtocol[i] == c.Protocol; ok {
			inArray++
		}
	}

	if inArray == 0 {
		return errors.New("Config error: Unknown protocol")
	}

	return nil
}
