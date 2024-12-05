package part1

import (
	"log"
	"strconv"
)

func process(input []string) int {
	total := 0
	for _, line := range input {
		for i := 0; i < len(line)-4; i++ {
			if line[i:i+4] == "mul(" {
				newpos, lineTotal := validateFunc(line, i+4)
				total += lineTotal
				i = newpos
			}
		}
	}
	return total
}

func validateFunc(input string, pos int) (newPos int, total int) {
	// get first number and check comma afterwards
	firstNum := ""
	newPos = pos
	for i := pos; i < len(input); i++ {
		newPos = i
		char := input[i]

		// End of string?dOESN'T MATTER IF IT'S A NUMBER OR A COMMA, THIS ISN'T VALID
		if i+1 == len(input) {
			return i, 0
		}
		if input[i] == ',' {
			if firstNum == "" {
				// not a valid first number
				return i, 0
			}
			break
		}
		if char >= '0' && char <= '9' {
			firstNum += string(char)
			continue
		}
		return i, 0
	}

	// Should have our first number now. Now get the second

	// position is current at the comma
	newPos++
	secondNum := ""

	for i := newPos; i < len(input); i++ {
		newPos = i
		char := input[i]

		if input[i] == ')' {
			if secondNum == "" {
				// not a valid first number
				return i, 0
			}
			break
		}

		// End of string?
		if i+1 == len(input) {
			return i, 0
		}

		if char >= '0' && char <= '9' {
			secondNum += string(char)
			continue
		}

		return i, 0
	}

	// Multiply the numbers together.
	num1, err := strconv.Atoi(firstNum)
	if err != nil {
		log.Fatalf("Error converting first number to int. Number: %v: %v", firstNum, err)
	}
	num2, err := strconv.Atoi(secondNum)
	if err != nil {
		log.Fatalf("Error converting second number to int. Number: %v: %v", secondNum, err)
	}
	return newPos, num1 * num2

}
