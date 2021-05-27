package lesson1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GoMillionFiles generate one million files
func GoMillionFiles() {
	var i = 1
	for i < 1000000 {
		createFile(i)
		i++
	}
}

// createFile function created files on files directory
func createFile(i int) {
	file, err := os.Create(strings.Join([]string{"./files/", strconv.Itoa(i), ".txt"}, ""))
	if err != nil {
		fmt.Println(fmt.Sprintf("error to create file: %s", err.Error()))
	}

	fmt.Println("file created: ", file.Name())
}
