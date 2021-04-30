package lesson7

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
				fmt.Println(res)
			}
			fmt.Print("Введите выражение через пробел (2 + 2): ")
			expressions = make([]string, 1)
		}

	}
}

func parseM(m []string) (Culcer, error) {
	c, ok := strconv.ParseUint(m[0], 10, 64)
	if ok == nil {
		d, ok := strconv.ParseUint(m[2], 10, 64)
		if ok != nil {
			return nil, errors.New("Тип второго аргумента не соответствует типу первого аргумента")
		}
		return Tuint{m: c, n: d}, nil
	}

	c1, ok := strconv.ParseFloat(m[0], 32)
	if ok == nil {
		d, ok := strconv.ParseFloat(m[2], 32)
		if ok != nil {
			return nil, errors.New("Тип второго аргумента не соответствует типу первого аргумента")
		}
		return Tfloat{m: c1, n: d}, nil
	}

	c2, ok := strconv.ParseInt(m[0], 10, 64)
	if ok == nil {
		d, ok := strconv.ParseInt(m[2], 10, 64)
		if ok != nil {
			return nil, errors.New("Тип второго аргумента не соответствует типу первого аргумента")
		}
		return Tint{m: c2, n: d}, nil
	}

	return nil, errors.New("Неизвестный тип")
}

func processStack(e []string) (string, error) {
	result := ""
	moper := map[string]func(Culcer) string{
		"+": Culcer.Plus,
		"-": Culcer.Minus,
		"*": Culcer.Mult,
		"/": Culcer.Div,
	}
	for _, v := range e {
		if strings.Trim(v, " ") == "" {
			continue
		}
		c := strings.Split(v, " ")
		if len(c)-1 < 2 {
			return "", errors.New("ошибка: вы ввели недостаточно аргументов или написали выражение без пробелов")
		}
		oper, err := parseM(c)
		if err != nil {
			return "", err
		}

		if _, ok := moper[c[1]]; !ok {
			return "", errors.New("ошибка: операция не распознана")
		}
		f := moper[c[1]]
		result = f(oper)
	}
	return result, nil
}
