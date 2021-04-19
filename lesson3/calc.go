package lesson3

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operations struct {
	m int
	n int
}

func (o Operations) Plus() int {
	return o.m + o.n
}

func (o Operations) Minus() int {
	return o.m - o.n
}

func (o Operations) Multiple() int {
	return o.m * o.n
}

func (o Operations) Div() int {
	return o.m / o.n
}

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
				fmt.Printf("Результат: %d\n", res)
			}
			fmt.Print("Введите выражение через пробел (2 + 2): ")
			expressions = make([]string, 1)
		}

	}
}

func parseArgs(c []string) (Operations, error) {
	num1, err := strconv.Atoi(c[0])
	if err != nil {
		return Operations{}, err
	}
	num2, err := strconv.Atoi(c[2])
	if err != nil {
		return Operations{}, err
	}
	return Operations{m: num1, n: num2}, err
}

func processStack(e []string) (int, error) {
	result := 0
	moper := map[string]func(Operations) int{
		"+": Operations.Plus,
		"-": Operations.Minus,
		"*": Operations.Multiple,
		"/": Operations.Div,
	}
	for _, v := range e {
		if strings.Trim(v, " ") == "" {
			continue
		}
		c := strings.Split(v, " ")
		if len(c)-1 < 2 {
			return 0, errors.New("ошибка: вы ввели недостаточно аргументов или написали выражение без пробелов")
		}
		oper, err := parseArgs(c)
		if err != nil {
			return 0, err
		}

		if _, ok := moper[c[1]]; !ok {
			return 0, errors.New("ошибка: операция не распознана")
		}
		f := moper[c[1]]
		result = f(oper)
	}
	return result, nil
}
