package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var input string
	if scanner.Scan() {
		input = scanner.Text()
	}

	date, err := time.Parse("2006-01-02 15:04:05", input)
	if err != nil {
		fmt.Println("Error parsing date: ", err)
		return
	}

	if date.Hour() >= 13 {
		date = date.AddDate(0, 0, 1)
	}
	fmt.Println(date.Format("2006-01-02 15:04:05"))
}
