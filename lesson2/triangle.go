package lesson2

import (
	"fmt"
	"math"
)

func Start() {
	var a int
	var b int

	var S float64
	var d float64

	var l int

	fmt.Println("Задание №1: ")
	fmt.Println("Вычисление площади прямоугольника: ")
	fmt.Println("Введите значение стороны a:")
	fmt.Scanf("%d\n", &a)

	fmt.Println("Введите значение стороны b:")
	fmt.Scanf("%d\n", &b)

	fmt.Printf("Площадь прямоугольника равна: %d\n", a*b)

	fmt.Println("")
	fmt.Println("Задание №2: ")
	fmt.Println("Вычисление диаметра и длины окружности: ")
	fmt.Println("Введите площадь окружности:")
	fmt.Scanf("%f\n", &S)

	x := S / math.Pi
	d = math.Sqrt(x)
	fmt.Printf("Диаметр окружности: %f\n", d)
	fmt.Printf("Длина окружности: %f\n", d*math.Pi)

	fmt.Println("")
	fmt.Println("Задание №3: ")
	fmt.Println("Вывести разряды числа: ")
	fmt.Println("Введите трехзначное число:")
	fmt.Scanf("%d\n", &l)

	fmt.Printf("Колличество сотен: %d\n", l/100)
	fmt.Printf("Колличество десяток: %d\n", l/10)
	fmt.Printf("Колличество единиц: %d\n", l/1)
}
