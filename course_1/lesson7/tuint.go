package lesson7

import (
	"fmt"
)

type Tuint struct {
	m uint64
	n uint64
}

func (o Tuint) Plus() string {
	return o.convert(o.m + o.n)
}

func (o Tuint) Minus() string {
	return o.convert(o.m - o.n)
}

func (o Tuint) Mult() string {
	return o.convert(o.m * o.n)
}

func (o Tuint) Div() string {
	return o.convert(o.m / o.n)
}

func (o Tuint) convert(l uint64) string {
	res := fmt.Sprintf("Результат: %d", l)
	return res
}
