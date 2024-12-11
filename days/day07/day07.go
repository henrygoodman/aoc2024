package day07

import (
	"aoc2024/common"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	input, err := common.ReadInput(7)
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

func validSolutionExists(result int, operands []int) bool {
    var backtrack func(index, current int) bool

    backtrack = func(index, current int) bool {
        if index < 0 {
            return current == 0
        }

        if current < 0 {
            return false
        }

        if backtrack(index-1, current-operands[index]) {
            return true
        }

        if current%operands[index] == 0 {
            if backtrack(index-1, current/operands[index]) {
                return true
            }
        }

        return false
    }

    return backtrack(len(operands)-1, result)
}

func solvePart1(input []string) int {
    var sum int

    for _, line := range input {
        parts := strings.Split(line, ":")

        if len(parts) != 2 {
            fmt.Println("Line could not be parsed")
            return -1
        }

        resultStr, operandsStr := strings.TrimSpace(parts[0]), strings.Fields(parts[1])
        var operands []int

        for _, numStr := range operandsStr {
            num, err := strconv.Atoi(numStr)
            if err != nil {
                fmt.Println("Error parsing int:", numStr)
                return -1
            }
            operands = append(operands, num)
        }

        result, err := strconv.Atoi(resultStr)
        if err != nil {
            fmt.Println("Error parsing result as int:", resultStr)
            return -1
        }

        if validSolutionExists(result, operands) {
            sum += result
        }
    }
    return sum
}

func validSolutionConcatExists(result int, operands []int) bool {
    var backtrack func(index, current int) bool

    backtrack = func(index, current int) bool {
        if index < 0 {
            return current == 0
        }

        if current < 0 {
            return false
        }

        if backtrack(index-1, current-operands[index]) {
            return true
        }

        if current%operands[index] == 0 {
            if backtrack(index-1, current/operands[index]) {
                return true
            }
        }

        concat := splitConcatenation(current, operands[index])
        if concat != -1 {
            if backtrack(index-1, concat) {
                return true
            }
        }

        return false
    }

    return backtrack(len(operands)-1, result)
}

// All ChatGPT magic here
func splitConcatenation(current, operand int) int {
    strCurr := strconv.Itoa(current)
    strOperand := strconv.Itoa(operand)
    if strings.HasSuffix(strCurr, strOperand) {
        remaining := strCurr[:len(strCurr)-len(strOperand)]
        if remaining == "" {
            return 0
        }
        res, err := strconv.Atoi(remaining)
        if err == nil {
            return res
        }
    }
    return -1
}


func solvePart2(input []string) int {
    var sum int

    for _, line := range input {
        parts := strings.Split(line, ":")

        if len(parts) != 2 {
            fmt.Println("Line could not be parsed")
            return -1
        }

        resultStr, operandsStr := strings.TrimSpace(parts[0]), strings.Fields(parts[1])
        var operands []int

        for _, numStr := range operandsStr {
            num, err := strconv.Atoi(numStr)
            if err != nil {
                fmt.Println("Error parsing int:", numStr)
                return -1
            }
            operands = append(operands, num)
        }

        result, err := strconv.Atoi(resultStr)
        if err != nil {
            fmt.Println("Error parsing result as int:", resultStr)
            return -1
        }

        if validSolutionConcatExists(result, operands) {
            sum += result
        }
    }
    return sum
}
