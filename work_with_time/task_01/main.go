package main

import (
	"fmt"
	"time"
)

func main() {
	var input string
	fmt.Scan(&input)
	parsedTime, err := time.Parse(time.RFC3339, input)
	if err != nil {
		fmt.Println("Ошибка при парсинге времени: ", err)
	}
	fmt.Println(parsedTime.Format(time.UnixDate))
}
