package day13

import (
	"aoc2024/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Solve() {
	input, err := common.ReadInput(13)
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

func parseInput(input []string) []machine {
	re := regexp.MustCompile(`(?m)Button A: X\+(\d+), Y\+(\d+)\s*Button B: X\+(\d+), Y\+(\d+)\s*Prize: X=(\d+), Y=(\d+)`)
	joinedInput := strings.Join(input, "\n")
	blocks := strings.Split(joinedInput, "\n\n")

	var machines []machine
	for _, block := range blocks {
		matches := re.FindStringSubmatch(block)
		if len(matches) == 7 {
			ax, _ := strconv.Atoi(matches[1])
			ay, _ := strconv.Atoi(matches[2])
			bx, _ := strconv.Atoi(matches[3])
			by, _ := strconv.Atoi(matches[4])
			prizeX, _ := strconv.Atoi(matches[5])
			prizeY, _ := strconv.Atoi(matches[6])
			machines = append(machines, machine{ax, ay, bx, by, prizeX, prizeY})
		}
	}
	return machines
}

type machine struct {
	ax, ay int
	bx, by int
	prizeX, prizeY int
}

func solvePart1(input []string) int {
	machines := parseInput(input)
	totalTokens := 0
	prizesWon := 0

	for _, m := range machines {
		tokens, solvable := solveLinearCombination(m.ax, m.ay, m.bx, m.by, m.prizeX, m.prizeY)
		if solvable {
			totalTokens += tokens
			prizesWon++
		}
	}

	return totalTokens
}

func solvePart2(input []string) int {
	machines := parseInput(input)
	totalTokens := 0
	prizesWon := 0

	offset := 10000000000000

	for _, m := range machines {
		prizeX := m.prizeX + offset
		prizeY := m.prizeY + offset


		tokens, solvable := solveLinearCombination(m.ax, m.ay, m.bx, m.by, prizeX, prizeY)
		if solvable {
			totalTokens += tokens
			prizesWon++
		}
	}
	return totalTokens
}


func solveLinearCombination(ax, ay, bx, by, targetX, targetY int) (int, bool) {
	det := ax*by - ay*bx
	if det == 0 {
		return 0, false
	}

	detX := targetX*by - targetY*bx
	detY := ax*targetY - ay*targetX

	x := detX / det
	y := detY / det

	if detX%det != 0 || detY%det != 0 || x < 0 || y < 0 {
		return 0, false
	}

	tokens := 3*x + 1*y
	return tokens, true
}
