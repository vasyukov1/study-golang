package task_02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solution() {

	fileReader, err := os.Open("Stepik/work_with_files/task_02/task.data")
	if err != nil {
		log.Fatal("Ошибка при открытии файла", err)
	}
	defer fileReader.Close()

	reader := bufio.NewReader(fileReader)
	position := 1

	for {
		el, err := reader.ReadString(';')
		if err != nil {
			log.Fatal("Ошибка при чтении файла", err)
		}

		el = strings.Trim(el, ";\n")
		number, err := strconv.Atoi(el)
		if err != nil {
			fmt.Println("Данный элемент - не число")
		} else {
			if number == 0 {
				fmt.Printf("Result: %v\n", position)
				return
			}
			position++
		}
	}
}
