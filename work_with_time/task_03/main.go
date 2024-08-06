package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	if scanner.Scan() {
		input = scanner.Text()
	}

	dates := strings.Split(input, ",")
	datesCount := len(dates)
	if datesCount != 2 {
		fmt.Printf("Incorrect data provided: len = %d", datesCount)
		return
	}
	firstDate, err := time.Parse("02.01.2006 15:04:05", dates[0])
	if err != nil {
		fmt.Println("Incorrect first date:", err)
		return
	}
	secondDate, err := time.Parse("02.01.2006 15:04:05", strings.TrimRight(dates[1], "\n"))
	if err != nil {
		fmt.Println("Incorrect second date:", err)
		return
	}
	if firstDate.After(secondDate) {
		firstDate, secondDate = secondDate, firstDate
	}
	res := time.Since(firstDate).Round(time.Second) - time.Since(secondDate).Round(time.Second)

	fmt.Printf("Result: %v\n", res)
}
