package lesson5

import (
	"fmt"
)

func fib(n int, m map[int]int) int {

	if n == 0 {
		return m[0]
	}

	if n == 1 {
		return m[1]
	}

	if _, ok := m[n]; !ok {
		m[n] = fib(n-1, m) + fib(n-2, m)
	}

	return m[n]
}

func GoFib() {
	var a int
	res := map[int]int{
		0: 0,
		1: 1,
	}

	fmt.Println("Вычисление числа Фибоначчи: ")
	fmt.Println("Введите число:")
	fmt.Scanf("%d\n", &a)

	fmt.Println("Результат: ")
	fmt.Println(fib(a, res))
}
