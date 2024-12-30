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

var pathMap map[string]string

func buildPathMap() {
    pathMap = make(map[string]string)

    // Prepare a list of keys and their coordinates from the directionalKeypadMapping
    positions := []struct {
        char rune
        coord Coordinate
    }{}
    for char, coord := range directionalKeypadMapping {
        positions = append(positions, struct {
            char  rune
            coord Coordinate
        }{char, coord})
    }

    // Build the path map between every pair of keys
    for _, from := range positions {
        for _, to := range positions {
            path := computePath(from.coord, to.coord)
            key := fmt.Sprintf("%c->%c", from.char, to.char)
            pathMap[key] = path
        }
    }
}

func computePath(from, to Coordinate) string {
    var path string
    dy := to.y - from.y
    dx := to.x - from.x

    // Move vertically
    if dy > 0 {
        path += repeat("v", dy)
    } else if dy < 0 {
        path += repeat("^", -dy)
    }

    // Move horizontally
    if dx > 0 {
        path += repeat(">", dx)
    } else if dx < 0 {
        path += repeat("<", -dx)
    }

    path += "A" // Press the key
    return path
}

type CacheValue struct {
	path        string
	endPosition Coordinate
}

var prefixCache = make(map[string]CacheValue)

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
    curr := 'A' // Starting key on the directional keypad
    paths := []string{""} // Initialize with an empty path
    prefixCache := make(map[string]string) // Cache for resolved prefixes

    for _, ch := range code {
        nextPaths := []string{}

        for _, path := range paths {
            // Check if this prefix is already resolved in the cache
            cachedPath, exists := prefixCache[path]
            if exists {
                nextPaths = append(nextPaths, cachedPath)
                continue
            }

            // Build the key for the pathMap lookup
            key := fmt.Sprintf("%c->%c", curr, ch)
            mappedPath, exists := pathMap[key]
            if !exists {
                fmt.Printf("Path not found for %c to %c\n", curr, ch)
                return []string{}
            }

            // Compute the new path and add it to the cache
            newPath := path + mappedPath
            prefixCache[path] = newPath // Cache the resolved path
            nextPaths = append(nextPaths, newPath)
        }

        // Update the current key to the target key
        curr = ch
        paths = nextPaths
    }

    return paths
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

// TODO: Add some kind of caching (it can persist across levels since the mapping is equivalent)
// - try depth first (character by character) instead of level by level
func solve(input []string, robots int) int {
	total := 0
	repeatCount := robots

	buildPathMap()

	for _, code := range input {

		code1Paths := processNumericalKeypad(code)

		shortestPaths := code1Paths
		for i := 0; i < repeatCount; i++ {
			nextPaths := []string{}
			minLength := int(^uint(0) >> 1)

			for _, path := range shortestPaths {
				if len(path) > minLength {
					continue
				}

				code2Paths := processDirectionalKeypad(path)

				for _, p2 := range code2Paths {
					pathLength := len(p2)

					// Update minLength and prune paths that exceed it
					if pathLength < minLength {
						minLength = pathLength
						nextPaths = []string{p2} // Start fresh with new shortest paths
					} else if pathLength == minLength {
						nextPaths = append(nextPaths, p2) // Add paths of equal length
					}
				}
			}

			shortestPaths = nextPaths
		}

		// Calculate the total based on the shortest path found
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
    buildPathMap() // Precompute paths between all states on the directional keypad

    totalComplexity := 0

    for _, code := range input {
        numericPart := extractNumericPart(code)

        // Step 1: Process the numerical keypad
        initialPaths := processNumericalKeypad(code)

        minComplexity := int(^uint(0) >> 1) // Set to maximum possible value

        // Try each initial path from the numerical keypad
        for _, initialPath := range initialPaths {
            // Frequency map for the directional keypad sequences
            frequencies := map[string]int{
                fmt.Sprintf("A->%c", initialPath[len(initialPath)-1]): 1,
            }

            // Step 2: Process the directional keypad for 25 robots
            for i := 0; i < 3; i++ {
                newFrequencies := map[string]int{}

                for path, freq := range frequencies {
                    // Extract the current position from the path (last character)
                    curr := path[len(path)-1]

                    for _, next := range "Av<>^" {
                        // Generate the key for pathMap lookup
                        key := fmt.Sprintf("%c->%c", curr, next)
                        _, exists := pathMap[key]
                        if !exists {
                            continue
                        }

                        // Add the new sequence to the frequencies
                        newKey := fmt.Sprintf("%c->%c", curr, next)
                        newFrequencies[newKey] += freq
                    }
                }

                frequencies = newFrequencies // Move to the next robot
            }

            // Calculate the complexity for this path
            complexity := 0
            for _, freq := range frequencies {
                complexity += freq
            }

            // Update the minimum complexity
            if complexity < minComplexity {
                minComplexity = complexity
            }
        }

        // Add the minimum complexity for this code
        totalComplexity += numericPart * minComplexity
    }

    return totalComplexity
}
