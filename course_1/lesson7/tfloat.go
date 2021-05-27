package lesson7

import (
	"fmt"
)

type Tfloat struct {
	m float64
	n float64
}

func (o Tfloat) Plus() string {
	return o.convert(o.m + o.n)
}

func (o Tfloat) Minus() string {
	return o.convert(o.m - o.n)
}

func (o Tfloat) Mult() string {
	return o.convert(o.m * o.n)
}

func (o Tfloat) Div() string {
	return o.convert(o.m / o.n)
}

func (o Tfloat) convert(l float64) string {
	res := fmt.Sprintf("Результат: %f", l)
	return res
}
