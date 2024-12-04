package part1

import (
	"log"
	"strconv"
)

var mulEnabled = true

func process(input []string) int {

	total := 0
	for _, line := range input {
		total += processLine(inputString(line))
	}

	return total
}

func processLine(line inputString) int {
	lineTotal := 0
	for i := 0; i < len(line); i++ {
		if line.checkAhead(i, "do()") {
			mulEnabled = true
			i += 4
		}
		if line.checkAhead(i, "don't()") {
			mulEnabled = false
			i += 7
		}
		if line.checkAhead(i, "mul(") {
			newpos, total := line.validateFunc(i + 4)
			if mulEnabled == true {
				lineTotal += total
			}
			i = newpos
		}
	}
	return lineTotal
}

type inputString string

func (input inputString) getChar(pos int) string {
	return string(input[pos])
}

func (input inputString) checkAhead(pos int, check string) bool {
	if len(input) < pos+len(check) {
		return false
	}
	for i := 0; i < len(check); i++ {
		if input.getChar(pos+i) != string(check[i]) {
			return false
		}
	}
	return true
}

func (input inputString) validateFunc(pos int) (newPos int, total int) {
	// get first number and check comma afterwards
	firstNum := ""
	newPos = pos
	for i := pos; i < len(input); i++ {
		newPos = i
		char := input[i]

		// End of string?
		if i+1 == len(input) {
			return i, 0
		}

		if input.getChar(i) == "," {
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
	newPos++
	secondNum := ""
	for i := newPos; i < len(input); i++ {
		newPos = i
		char := input[i]

		if input.getChar(i) == ")" {
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
