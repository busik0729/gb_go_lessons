package lesson8

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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

func ParseConfFile(path string) (Conf, error) {
	conf := Conf{}

	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return conf, errors.New("Config error: Cannot read config file")
	}

	errUnmarshalConf := json.Unmarshal(dat, &conf)
	if errUnmarshalConf != nil {
		return conf, errors.New("Config error: Error unmarshal config file")
	}

	return conf, nil
}
