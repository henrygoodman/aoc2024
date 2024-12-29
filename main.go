package main

import (
	"aoc2024/days/day01"
	"aoc2024/days/day02"
	"aoc2024/days/day03"
	"aoc2024/days/day04"
	"aoc2024/days/day05"
	"aoc2024/days/day06"
	"aoc2024/days/day07"
	"aoc2024/days/day08"
	"aoc2024/days/day09"
	"aoc2024/days/day10"
	"aoc2024/days/day11"
	"aoc2024/days/day12"
	"aoc2024/days/day13"
	"aoc2024/days/day14"
	"aoc2024/days/day15"
	"aoc2024/days/day16"
	"aoc2024/days/day17"
	"aoc2024/days/day18"
	"aoc2024/days/day19"
	"aoc2024/days/day20"
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
		day01.Solve()
	case 2:
		day02.Solve()
	case 3:
		day03.Solve()
	case 4:
		day04.Solve()
	case 5:
		day05.Solve()
	case 6:
		day06.Solve()
	case 7:
		day07.Solve()
	case 8:
		day08.Solve()
	case 9:
		day09.Solve()
	case 10:
		day10.Solve()
	case 11:
		day11.Solve()
	case 12:
		day12.Solve()
	case 13:
		day13.Solve()
	case 14:
		day14.Solve()
	case 15:
		day15.Solve()
	case 16:
		day16.Solve()
	case 17:
		day17.Solve()
	case 18:
		day18.Solve()
	case 19:
		day19.Solve()
	case 20:
		day20.Solve()
	default:
		fmt.Printf("Day %d not implemented yet\n", day)
	}

	fmt.Printf("Execution Time: %v\n", time.Since(start))
}
