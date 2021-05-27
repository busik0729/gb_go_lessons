package lesson1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GoMillionFiles() {
	var i = 1
	for i < 1000000 {
		createFile(i)
		i++
	}
}

func createFile(i int) {
	file, err := os.Create(strings.Join([]string{"./files/", strconv.Itoa(i), ".txt"}, ""))
	if err != nil {
		fmt.Println(fmt.Sprintf("error to create file: %s", err.Error()))
	}

	fmt.Println("file created: ", file.Name())
}
