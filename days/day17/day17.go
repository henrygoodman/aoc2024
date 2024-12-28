package day17

import (
	"aoc2024/common"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Solve() {
	input, err := common.ReadInput(17)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	machines := parseInput(input)

	if len(machines) != 1 {
		fmt.Println("Unexpected number of machines")
		return
	}

	common.Time("Part 1", func() {
		fmt.Println("Part 1 Answer:", solvePart1(machines[0]))
	})

	common.Time("Part 2", func() {
		fmt.Println("Part 2 Answer:", solvePart2(machines[0]))
	})
}

type Machine struct {
	registerA int;
	registerB int;
	registerC int;
	program []int;
}


func parseInput(input []string) []Machine {
	joinedInput := strings.Join(input, "\n")
	re := regexp.MustCompile(`Register A:\s*(\d+)\s*Register B:\s*(\d+)\s*Register C:\s*(\d+)\s*Program:\s*([\d,]+)`)

	matches := re.FindAllStringSubmatch(joinedInput, -1)

	var machines []Machine
	for _, match := range matches {
		if len(match) == 5 {
			registerA, _ := strconv.Atoi(match[1])
			registerB, _ := strconv.Atoi(match[2])
			registerC, _ := strconv.Atoi(match[3])
			programStr := match[4]
			programParts := strings.Split(programStr, ",")
			var program []int
			for _, part := range programParts {
				value, _ := strconv.Atoi(part)
				program = append(program, value)
			}
			machines = append(machines, Machine{
				registerA: registerA,
				registerB: registerB,
				registerC: registerC,
				program:   program,
			})
		}
	}
	return machines
}

func getComboOperand(machine *Machine, operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return machine.registerA
	case 5:
		return machine.registerB
	case 6:
		return machine.registerC
	default:
		fmt.Println("Unexpected operand", operand)
		return -1
	}
}

func adv(machine *Machine, operand int) {
	numerator := machine.registerA
	denominator := math.Pow(2, float64(getComboOperand(machine, operand)))
	machine.registerA = int(float64(numerator) / denominator)
}

func bxl(machine *Machine, operand int) {
	machine.registerB ^= operand
}

func bst(machine *Machine, operand int) {
	machine.registerB = getComboOperand(machine, operand) % 8
}

func jnz(machine *Machine, operand int, ip int) int {
	if machine.registerA == 0 { return ip }
	return operand - 2

}

func bxc(machine *Machine, _ int) {
	machine.registerB ^= machine.registerC
}

func out(machine *Machine, operand int) int {
	return getComboOperand(machine, operand) % 8
}

func bdv(machine *Machine, operand int) {
	numerator := machine.registerA
	denominator := math.Pow(2, float64(getComboOperand(machine, operand)))
	machine.registerB = int(float64(numerator) / denominator)
}

func cdv(machine *Machine, operand int) {
	numerator := machine.registerA
	denominator := math.Pow(2, float64(getComboOperand(machine, operand)))
	machine.registerC = int(float64(numerator) / denominator)
}

func solvePart1(machine Machine) int {
	fmt.Println(solveMachineOutput(machine))
	return 0
}

// Observations:
// - we only ever output register B
// - TBH even though we output to register B, I don't think we need to care specifically about register B, it depends on the final octal in regA
// (that is to say, registersB and C do matter, but they are kinda transitively dependant on registerA, so fiddling with A will also fiddle B/C)
// - the program runs in a loop restarting at end (always same instruction parity)
// - if we have the instruction (0 3) it is equivalent to bitshift registerA 3 (% 8)
// - at each program loop, the 3 least significant bits of A determine the output digit, and then A is bitshifted 3
// - we can try outputting each digit, then we we find a match, shift 3 and try again. Backtrack to handle deadends
// (since multiple A values could potentially lead to the same digit)

func solvePart2(machine Machine) int {
	var backtrack func(machine Machine, A int, compareIndex int) int

	backtrack = func(machine Machine, A int, compareIndex int) int {
		program := machine.program
		expectedOutput := program[len(program)-compareIndex:]

		// Try all 8 possible values for the 3 least significant bits
		for n := 0; n < 8; n++ {
			// Compute a candidate value for registerA
			candidateA := (A << 3) | n

			newMachine := machine
			newMachine.registerA = candidateA
			output := solveMachineOutput(newMachine)

			// Check if the output matches the expected pattern
			if len(output) >= compareIndex {
				outputSegment := output[len(output)-compareIndex:]

				if fmt.Sprint(outputSegment) == fmt.Sprint(expectedOutput) {
					if compareIndex == len(program) {
						return candidateA
					}

					// Otherwise, continue searching deeper
					result := backtrack(newMachine, candidateA, compareIndex+1)
					if result != -1 {
						return result
					}
				}
			}
		}

		return -1
	}

	return backtrack(machine, 0, 1)
}

func solveMachineOutput(machine Machine) []int {
	ip := 0
	var output []int

	for ip < len(machine.program)-1 {
		opcode := machine.program[ip]
		operand := machine.program[ip+1]

		switch opcode {
		case 0:
			adv(&machine, operand)
		case 1:
			bxl(&machine, operand)
		case 2:
			bst(&machine, operand)
		case 3:
			ip = jnz(&machine, operand, ip)
		case 4:
			bxc(&machine, operand)
		case 5:
			value := out(&machine, operand)
			output = append(output, value)
		case 6:
			bdv(&machine, operand)
		case 7:
			cdv(&machine, operand)
		}

		ip += 2
	}

	return output
}


