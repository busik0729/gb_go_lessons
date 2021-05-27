package main

import (
	"fmt"
	"github.com/busik0729/gb_go_lessons/course_2/lesson1"
	"runtime/debug"
	"time"
)

type ErrorWithTrace struct {
	text  string
	trace string
	time  string
}

func New(text string) error {
	return &ErrorWithTrace{
		text:  text,
		trace: string(debug.Stack()),
		time:  time.Now().String(),
	}
}

func (e *ErrorWithTrace) Error() string {
	return fmt.Sprintf("error: %s\ntrace:\n%s \ntime: %s", e.text, e.trace, e.time)
}

func main() {
	var err error
	defer func() {
		if v := recover(); v != nil {
			err = New(fmt.Sprintf("%s", v))

			fmt.Println("Panic was catched: ", err.Error())
		}
	}()

	// lesson1.GoPanic()
	lesson1.GoMillionFiles()
}
