package main

import (
	"fmt"
	"github.com/busik0729/gb_go_lessons/lesson8"
)

func main() {
	conf, err := lesson8.ParseConfFile("conf.json")

	if err != nil {
		fmt.Println(err.Error())
	}

	ok, err := conf.Validate()
	if !ok {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(conf)
}
