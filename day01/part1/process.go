package part1

import (
	"silverark/aoc-2024/pkg/shared"
	"slices"
	"strings"
)

func process(input []string) int {

	var list1 []int
	var list2 []int

	for _, line := range input {
		numbers := strings.Fields(line)
		list1 = append(list1, shared.Atoi(numbers[0]))
		list2 = append(list2, shared.Atoi(numbers[1]))
	}

	// Sort the lists
	slices.Sort(list1)
	slices.Sort(list2)

	distance := 0
	for i := 0; i < len(list1); i++ {
		diff := list1[i] - list2[i]
		if diff < 0 {
			diff = -diff
		}
		distance += diff
	}

	return distance
}
