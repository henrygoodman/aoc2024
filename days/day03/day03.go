package day03

import (
	"aoc2024/common"
	"fmt"
	"regexp"
	"strconv"
)

func Solve() {
	input, err := common.ReadInput(3)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	common.Time("Part 1", func() {
		fmt.Println("Part 1 Answer:", solvePart1(input))
	})

	common.Time("Part 2", func() {
		fmt.Println("Part 2 Answer:", solvePart2(input))
	})
}

func solvePart1(input []string) int {
    mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

    total := 0
    for _, line := range input {
        matches := mulRe.FindAllStringSubmatch(line, -1)
        for _, match := range matches {
            x, err1 := strconv.Atoi(match[1])
            y, err2 := strconv.Atoi(match[2])
            if err1 == nil && err2 == nil {
                total += x * y
            }
        }
    }

    return total
}


func solvePart2(input []string) int {
    combinedRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\)`)

    total := 0
    multiplicationEnabled := true

    for _, line := range input {
        matches := combinedRe.FindAllStringSubmatch(line, -1)
        for _, match := range matches {
            if match[1] != "" && match[2] != "" {
                if multiplicationEnabled {
                    x, _ := strconv.Atoi(match[1])
                    y, _ := strconv.Atoi(match[2])
                    total += x * y
                }
            } else if match[0] == "don't()" {
                multiplicationEnabled = false
            } else if match[0] == "do()" {
                multiplicationEnabled = true
            }
        }
    }

    return total
}
