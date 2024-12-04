package part1

import (
	"silverark/aoc-2024/pkg/shared"
	"strings"
)

func process(input []string) int {
	safeCount := 0
	for _, line := range input {
		if processLine(line) {
			safeCount++
		}
	}
	return safeCount
}

func processLine(line string) bool {
	safe, badLevel := isSafe(line)
	if safe {
		return true
	}

	// Try block of 3
	for i := badLevel; i > badLevel-3; i-- {
		if i < 0 {
			break
		}
		safe, _ = dampenerTest(line, i)
		if safe {
			return true
		}
	}
	return false
}

func isSafe(line string) (bool, int) {
	numbers := strings.Fields(line)

	if len(numbers) < 0 {
		return false, 0
	}

	lastNum := shared.Atoi(numbers[0])
	incOrDec := ""

	for i := 1; i < len(numbers); i++ {
		currentNum := shared.Atoi(numbers[i])
		// Same number return false
		if lastNum == currentNum {
			return false, i
		}

		diff := currentNum - lastNum

		// Set increment or decrement if not set
		if incOrDec == "" {
			incOrDec = "inc"
			if diff < 0 {
				incOrDec = "dec"
			}
		}

		// Check inc or dec
		if incOrDec == "inc" {
			if diff < 0 {
				return false, i
			}
			if diff > 3 {
				return false, i
			}
		}
		if incOrDec == "dec" {
			if diff > 0 {
				return false, i
			}
			if diff < -3 {
				return false, i
			}
		}
		lastNum = currentNum
	}

	return true, 0
}

func dampenerTest(line string, badLevel int) (bool, int) {
	numbers := strings.Fields(line)
	newNumbers := append(numbers[:badLevel], numbers[badLevel+1:]...)
	newLine := strings.Join(newNumbers, " ")
	return isSafe(newLine)
}
