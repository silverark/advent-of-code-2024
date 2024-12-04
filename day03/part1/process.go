package part1

import (
	"log"
	"strconv"
)

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

		if line.getChar(i) == "m" {
			if line.getChar(i+1) == "u" {
				if line.getChar(i+2) == "l" {
					if line.getChar(i+3) == "(" {
						newpos, total := line.validateFunc(i + 4)

						if total > 0 {
							lineTotal += total
							log.Printf("Total: %v\t String: %v ", total, line[i:newpos+1])
						}

						i = newpos
					}
				}
			}
		}
	}

	return lineTotal

}

type inputString string

func (input inputString) getChar(pos int) string {
	return string(input[pos])
}

func (input inputString) validateFunc(pos int) (newPos int, total int) {
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

		if input.getChar(i) == "," {
			if firstNum == "" {
				// not a valid first number
				return i, 0
			}
			break
		}

		if char >= '0' && char <= '9' {
			if firstNum == "" && '0' == char {
				log.Println("Number starts with 0. Is this OK? At position: ", i)
			}
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
			if secondNum == "" && '0' == char {
				log.Println("secondNum starts with 0. Is this OK? At position: ", i)
			}
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