package lesson8

import (
	"os"
	"reflect"
	"strconv"
)

func ParseConf() Conf {
	conf := Conf{}
	t := reflect.TypeOf(conf)
	f := reflect.ValueOf(&conf)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get("env")

		d := f.Elem().FieldByName(field.Name)
		e := getEnv(tag)

		if d.Kind() == reflect.String {
			d.SetString(e)
			continue
		}

		if d.Kind() == reflect.Int {
			p, _ := strconv.ParseInt(e, 10, 64)
			d.SetInt(p)
		}
	}

	return conf
}

func getEnv(name string) string {
	return os.Getenv(name)
}
