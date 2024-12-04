package part1

import (
	"log"
	"silverark/aoc-2024/pkg/shared"
	"strings"
)

func process(input []string) int {

	safeCount := 0
	for _, line := range input {
		if isSafe(line) {
			log.Printf("Safe line: %v", line)
			safeCount++
		}
	}
	return safeCount
}

func isSafe(line string) bool {
	numbers := strings.Fields(line)

	if len(numbers) < 0 {
		return false
	}

	lastNum := shared.Atoi(numbers[0])
	incOrDec := ""

	for i := 1; i < len(numbers); i++ {
		currentNum := shared.Atoi(numbers[i])
		// Same number return false
		if lastNum == currentNum {
			return false
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
				return false
			}
			if diff > 3 {
				return false
			}
		}
		if incOrDec == "dec" {
			if diff > 0 {
				return false
			}
			if diff < -3 {
				return false
			}
		}

		lastNum = currentNum
	}

	return true
}
