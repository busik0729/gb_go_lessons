package lesson7

import (
	"fmt"
)

type Tint struct {
	m int64
	n int64
}

func (o Tint) Plus() string {
	return o.convert(o.m + o.n)
}

func (o Tint) Minus() string {
	return o.convert(o.m - o.n)
}

func (o Tint) Mult() string {
	return o.convert(o.m * o.n)
}

func (o Tint) Div() string {
	return o.convert(o.m / o.n)
}

func (o Tint) convert(l int64) string {
	res := fmt.Sprintf("Результат: %d", l)
	return res
}
