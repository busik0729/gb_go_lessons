package lesson3

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunCalc() {
	expressions := make([]string, 1)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Введите выражение через пробел (2 + 2): ")
		for scanner.Scan() {
			expressions = append(expressions, scanner.Text())
			res, err := processStack(expressions)
			if err != nil {
				fmt.Printf("ошибка: %s\n", err)
			} else {
				fmt.Printf("Результат: %.2f\n", res)
			}
			fmt.Print("Введите выражение через пробел (2 + 2): ")
			expressions = make([]string, 1)
		}

	}
}

func parseArgs(c []string) (float64, float64, error) {
	num1, err := strconv.ParseFloat(c[0], 64)
	if err != nil {
		return 0.0, 0.0, err
	}
	num2, err := strconv.ParseFloat(c[2], 64)
	if err != nil {
		return 0.0, 0.0, err
	}
	return num1, num2, nil
}

func processStack(e []string) (float64, error) {
	result := 0.0
	for _, v := range e {
		if strings.Trim(v, " ") == "" {
			continue
		}
		c := strings.Split(v, " ")
		if len(c)-1 < 2 {
			return 0.0, errors.New("ошибка: вы ввели недостаточно аргументов или написали выражение без пробелов")
		}
		num1, num2, err := parseArgs(c)
		if err != nil {
			return 0.0, err
		}
		switch c[1] {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0.0 {
				return 0.0, errors.New("ошибка: на нуль делить нельзя.")
			}
			result = num1 / num2
		}
	}
	return result, nil
}
