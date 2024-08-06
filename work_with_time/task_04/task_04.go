package main

import (
	"fmt"
	"time"
)

func main() {
	const base int64 = 1589570165

	var minutes, seconds time.Duration
	fmt.Scanf("%d мин. %d сек.", &minutes, &seconds)

	baseTime := time.Unix(base, 0).UTC()
	newTime := baseTime.Add(minutes * time.Minute).Add(seconds * time.Second)
	fmt.Printf("Result: %v\n", newTime.Format(time.UnixDate))
}
