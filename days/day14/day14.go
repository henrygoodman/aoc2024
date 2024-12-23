package day14

import (
	"aoc2024/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Solve() {
	input, err := common.ReadInput(14)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	robots := parseInput(input);

	common.Time("Part 1", func() {
		fmt.Println("Part 1 Answer:", solvePart1(robots, Coordinate{101, 103}))
	})

	common.Time("Part 2", func() {
		fmt.Println("Part 2 Answer:", solvePart2(robots, Coordinate{101, 103}))
	})
}

type Coordinate struct {
	x, y int
}

type Velocity struct {
	vx, vy int
}

type Robot struct {
	start Coordinate
	velocity Velocity
}

func parseInput(input []string) []Robot {
	re := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
	joinedInput := strings.Join(input, "\n")
	lines := strings.Split(joinedInput, "\n")
	
	var robots []Robot
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) == 5 {
			startX, _ := strconv.Atoi(matches[1])
			startY, _ := strconv.Atoi(matches[2])
			velocityX, _ := strconv.Atoi(matches[3])
			velocityY, _ := strconv.Atoi(matches[4])
			
			robot := Robot{
				start: Coordinate{x: startX, y: startY},
				velocity: Velocity{vx: velocityX, vy: velocityY},
			}
			robots = append(robots, robot)
		}
	}
	return robots
}


func findRobotPosition(robot Robot, gridSize Coordinate, iterations int) Coordinate {
	finalXPos := robot.start.x + robot.velocity.vx * iterations
	finalYPos := robot.start.y + robot.velocity.vy * iterations
	
	// Wrap the final positions using the board size
	finalXPos = finalXPos % gridSize.x
	finalYPos = finalYPos % gridSize.y

	// If any are negative, subtract from the gridSize (for wrap)
	if finalXPos < 0 {
		finalXPos = gridSize.x + finalXPos
	}

	if finalYPos < 0 {
		finalYPos = gridSize.y + finalYPos
	}

	return Coordinate{finalXPos, finalYPos}
}

func getQuadrant(position Coordinate, gridSize Coordinate) int {
	if position.x < gridSize.x / 2 && position.y < gridSize.y / 2 {
		return 0
	}
	if position.x > gridSize.x / 2 && position.y < gridSize.y / 2 {
		return 1
	}
	if position.x < gridSize.x / 2 && position.y > gridSize.y / 2 {
		return 2
	}
	if position.x > gridSize.x / 2 && position.y > gridSize.y / 2 {
		return 3
	}
	return -1
} 

func solvePart1(input []Robot, gridSize Coordinate) int {
	iterations := 100
	quadrantCount := make([]int, 4)

	for _, robot := range input {
		pos := findRobotPosition(robot, gridSize, iterations)
		quadrant := getQuadrant(pos, gridSize)
		if quadrant != -1 {
			quadrantCount[quadrant] += 1
		}
	}

	ret := 1

	for _, i := range quadrantCount {
		ret *= i
	}

	return ret
}

func solvePart2(input []Robot, gridSize Coordinate) int {
	maxIterations := 101 * 103 // Maximum number of iterations to check
	lowestRating := SafetyRating{iteration: -1, rating: int(^uint(0) >> 1)}

	for iterations := 0; iterations < maxIterations; iterations++ {
		quadrantCount := make([]int, 4)

		for _, robot := range input {
			pos := findRobotPosition(robot, gridSize, iterations)
			quadrant := getQuadrant(pos, gridSize)
			if quadrant != -1 {
				quadrantCount[quadrant]++
			}
		}

		rating := 1
		for _, count := range quadrantCount {
			if count > 0 {
				rating *= count
			}
		}

		if rating < lowestRating.rating {
			lowestRating = SafetyRating{iteration: iterations, rating: rating}
		}
	}

	return lowestRating.iteration
}


type SafetyRating struct {
	iteration int
	rating    int
}

