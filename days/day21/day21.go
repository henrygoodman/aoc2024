package day21

import (
	"aoc2024/common"
	"fmt"
	"strconv"
	"unicode"
)

func Solve() {
	input, err := common.ReadInput(21)
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

type Coordinate struct {
	y, x int
}

/*

	+---+---+---+
	| 7 | 8 | 9 |
	+---+---+---+
	| 4 | 5 | 6 |
	+---+---+---+
	| 1 | 2 | 3 |
	+---+---+---+
		| 0 | A |
		+---+---+

		+---+---+
		| ^ | A |
	+---+---+---+
	| < | v | > |
	+---+---+---+

*/

var keypadMapping = map[rune]Coordinate{
	'0': {3,1},
	'1': {2,0},
	'2': {2,1},
	'3': {2,2},
	'4': {1,0},
	'5': {1,1},
	'6': {1,2},
	'7': {0,0},
	'8': {0,1},
	'9': {0,2},
	'A': {3,2},
}

var directionalKeypadMapping = map[rune]Coordinate{
	'A': {0, 2},
	'^': {0, 1},
	'<': {1, 0},
	'v': {1, 1},
	'>': {1, 2},
}

// 029A -> <A^A>^^AvvvA -> v<<A>>^A<A>AvA<^AA>A<vAAA>^A -> <vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A
// - at each stage, we need to find the shortest path that produces the output
// if we have 3 candidate shortest at stage 1, we may need to try them all stage 2 to get shortest for stage 2
// to get from 1 digit to the next, there are 2 ways (y then x or x then y), IDK if it matters or we just choose 1
// - the only difference in the paths is between the A's. The A's dont change. If its a straight line, it doesnt change.
// - intuitively you want to stay on one button as much as possible
// - is the choice of button important (horizontal/vertical)
// - YES you cant go into the gap

func processNumericalKeypad(code string) []string {
	start := Coordinate{3, 2}
	curr := start
	paths := map[string]bool{"": true}

	gap := Coordinate{3, 0}

	for _, digit := range code {
		targetPos := keypadMapping[digit]
		newPaths := map[string]bool{}

		dy := targetPos.y - curr.y
		dx := targetPos.x - curr.x
		if isValidPath(curr, targetPos, gap, true) {
			for path := range paths {
				newPath := path
				if dy > 0 {
					newPath += repeat("v", dy)
				} else if dy < 0 {
					newPath += repeat("^", -dy)
				}
				if dx > 0 {
					newPath += repeat(">", dx)
				} else if dx < 0 {
					newPath += repeat("<", -dx)
				}
				newPath += "A"
				newPaths[newPath] = true
			}
		}

		if isValidPath(curr, targetPos, gap, false) {
			for path := range paths {
				newPath := path
				if dx > 0 {
					newPath += repeat(">", dx)
				} else if dx < 0 {
					newPath += repeat("<", -dx)
				}
				if dy > 0 {
					newPath += repeat("v", dy)
				} else if dy < 0 {
					newPath += repeat("^", -dy)
				}
				newPath += "A"
				newPaths[newPath] = true
			}
		}

		paths = newPaths
		curr = targetPos
	}

	uniquePaths := make([]string, 0, len(paths))
	for path := range paths {
		uniquePaths = append(uniquePaths, path)
	}

	return uniquePaths
}

func processDirectionalKeypad(code string) []string {
	start := Coordinate{0, 2}
	curr := start
	paths := map[string]bool{"": true}
	gap := Coordinate{0, 0}

	for _, ch := range code {
		targetPos := directionalKeypadMapping[ch]
		newPaths := map[string]bool{}

		dy := targetPos.y - curr.y
		dx := targetPos.x - curr.x

		if isValidPath(curr, targetPos, gap, true) {
			for path := range paths {
				newPath := path
				if dy > 0 {
					newPath += repeat("v", dy)
				} else if dy < 0 {
					newPath += repeat("^", -dy)
				}
				if dx > 0 {
					newPath += repeat(">", dx)
				} else if dx < 0 {
					newPath += repeat("<", -dx)
				}
				newPath += "A"
				newPaths[newPath] = true
			}
		}

		if isValidPath(curr, targetPos, gap, false) {
			for path := range paths {
				newPath := path
				if dx > 0 {
					newPath += repeat(">", dx)
				} else if dx < 0 {
					newPath += repeat("<", -dx)
				}
				if dy > 0 {
					newPath += repeat("v", dy)
				} else if dy < 0 {
					newPath += repeat("^", -dy)
				}
				newPath += "A"
				newPaths[newPath] = true
			}
		}

		paths = newPaths
		curr = targetPos
	}

	uniquePaths := make([]string, 0, len(paths))
	for path := range paths {
		uniquePaths = append(uniquePaths, path)
	}

	return uniquePaths
}

func isValidPath(curr, target, gap Coordinate, verticalFirst bool) bool {
	if verticalFirst {
		if curr.x == gap.x && ((curr.y < gap.y && target.y >= gap.y) || (curr.y > gap.y && target.y <= gap.y)) {
			return false
		}
	} else {
		if curr.y == gap.y && ((curr.x < gap.x && target.x >= gap.x) || (curr.x > gap.x && target.x <= gap.x)) {
			return false
		}
	}
	return true
}

func repeat(char string, n int) string {
	ret := ""
	for i := 0; i < n; i++ {
		ret += char
	}
	return ret
}

func extractNumericPart(code string) int {
	numeric := ""
	for _, ch := range code {
		if unicode.IsDigit(ch) {
			numeric += string(ch)
		}
	}
	if numeric == "" {
		return 0
	}
	result, err := strconv.Atoi(numeric)
	if err != nil {
		fmt.Println("Error converting numeric part:", err)
		return 0
	}
	return result
}

func solve(input []string, robots int) int {
	total := 0
	repeatCount := 2

	for _, code := range input {

		code1Paths := processNumericalKeypad(code)

		shortestPaths := code1Paths
		for i := 0; i < repeatCount; i++ {
			nextPaths := []string{}
			minLength := int(^uint(0) >> 1)
			for _, path := range shortestPaths {
				code2Paths := processDirectionalKeypad(path)
				for _, p2 := range code2Paths {
					pathLength := len(p2)
					if pathLength < minLength {
						minLength = pathLength
						nextPaths = []string{p2}
					} else if pathLength == minLength {
						nextPaths = append(nextPaths, p2)
					}
				}
			}

			shortestPaths = nextPaths
		}

		numericPart := extractNumericPart(code)

		if len(shortestPaths) > 0 {
			total += numericPart * len(shortestPaths[0])
		}
	}

	return total
}

func solvePart1(input []string) int {
	return solve(input, 2)
}

func solvePart2(input []string) int {
	return solve(input, 2)
}