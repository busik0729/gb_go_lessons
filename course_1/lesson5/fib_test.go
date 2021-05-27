package lesson5

import (
	"testing"
)

func TestGoFib(t *testing.T) {
	res := map[int]int{
		0: 0,
		1: 1,
	}
	cases := []struct {
		Name string
		In   int
		Out  int
	}{
		{Name: "First", In: 1, Out: 1},
		{Name: "Second", In: 2, Out: 1},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			if got := fib(tt.In, res); got != tt.Out {
				t.Fail()
			}
		})
	}
}
