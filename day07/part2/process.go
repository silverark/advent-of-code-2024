package part1

import (
	"log"
	"silverark/aoc-2024/pkg/shared"
	"strconv"
	"strings"
)

func process(input []string) int {
	validTotals := 0

	for _, line := range input {
		parts := strings.Split(line, ":")
		lineTotal := shared.Atoi(strings.TrimSpace(parts[0]))

		numberStrings := strings.Fields(strings.TrimSpace(parts[1]))
		var numbers []int

		for _, numStr := range numberStrings {
			numbers = append(numbers, shared.Atoi(numStr))
		}

		if len(numbers) < 2 {
			log.Fatal("A line has less than 2 numbers")
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

				concat := append([]rune(nil), combination...)
				concat = append(concat, '|')
				newResults = append(newResults, concat)
			}
			operatorCombinations = newResults
		}

		hasValid := false

		// Check all operator combinations
		for _, ops := range operatorCombinations {

			result := numbers[0]
			for i, op := range ops {
				if result > lineTotal {
					break
				}
				nextNum := numbers[i+1]
				switch op {
				case '+':
					result += nextNum
				case '*':
					result *= nextNum
				case '|':
					result = concatNums(result, nextNum)
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
func concatNums(a, b int) int {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)
	combined := shared.Atoi(aStr + bStr)
	return combined
}
