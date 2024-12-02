package main

import (
	"aoc2024/days"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day>")
		return
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil || day < 1 || day > 25 {
		fmt.Println("Please provide a valid day between 1 and 25")
		return
	}

	start := time.Now()

	switch day {
	case 1:
		days.Day01()
	default:
		fmt.Printf("Day %d not implemented yet\n", day)
	}

	fmt.Printf("Execution Time: %v\n", time.Since(start))
}
