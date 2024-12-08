package part1

import (
	"log"
	"silverark/aoc-2024/pkg/shared"
	"strings"
)

func process(input []string) int {

	validTotals := 0

	for _, line := range input {
		parts := strings.Split(line, ":")
		lineTotal := shared.Atoi(parts[0])

		numberStrings := strings.Fields(parts[1])
		var numbers []int

		for _, num := range numberStrings {
			numbers = append(numbers, shared.Atoi(num))
		}

		if len(numbers) < 2 {
			log.Fatal("A line has kless than 2 numbers")
		}

		operatorCombinations := [][]rune{{}}
		for i := 0; i < (len(numbers) - 1); i++ {
			var newResults [][]rune
			for _, combination := range operatorCombinations {
				add := append([]rune(nil), combination...)
				add = append(add, '+')
				newResults = append(newResults, add)

				multi := append([]rune(nil), combination...)
				multi = append(multi, '*')
				newResults = append(newResults, multi)
			}
			operatorCombinations = newResults
		}

		hasValid := false

		for _, ops := range operatorCombinations {

			// Get the total
			result := numbers[0]
			for i, op := range ops {
				switch op {
				case '+':
					result += numbers[i+1]
				case '*':
					result *= numbers[i+1]
				}
			}

			if result == lineTotal {
				hasValid = true
				break
			}
		}

		if hasValid {
			validTotals += lineTotal
		}

	}

	return validTotals
}
