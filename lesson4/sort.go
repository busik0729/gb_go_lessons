package lesson4

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Sort() {
	expressions := make([]string, 1)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Введите последовательность целых чисел через запятую (2,4,1,0): ")
		for scanner.Scan() {
			expressions = append(expressions, scanner.Text())
			res := processStack(expressions)
			fmt.Printf("Результат: %v\n", res)
			fmt.Print("Введите последовательность целых чисел через запятую (2,4,1,0): ")
			expressions = make([]string, 1)
		}
	}
}

func processStack(e []string) []int {
	result := make([]int, 1)
	for _, v := range e {
		if strings.Trim(v, " ") == "" {
			continue
		}
		c := strings.Split(v, ",")

		parseCollect := func(i string) int {
			i = strings.Trim(i, " ")
			r, err := strconv.Atoi(i)
			if err != nil {
				panic("Ошибка: в последовательности присутствует не число!")
			}
			return r
		}

		result = collect(c, parseCollect)
	}
	return result
}

func collect(list []string, f func(string) int) []int {
	result := make([]int, len(list))
	for i, item := range list {
		result[i] = f(item)
	}
	sort.Ints(result)
	return result
}
