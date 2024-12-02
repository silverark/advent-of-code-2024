package part1

import (
	"silverark/aoc-2024/pkg/shared"
	"strings"
)

func process(input []string) int {

	var list1 []int
	list2 := make(map[int]int)
	for _, line := range input {
		numbers := strings.Fields(line)
		list1 = append(list1, shared.Atoi(numbers[0]))
		list2Num := shared.Atoi(numbers[1])
		if _, found := list2[shared.Atoi(numbers[1])]; !found {
			list2[list2Num] = 0
		}
		list2[list2Num]++
	}

	total := 0
	for _, num := range list1 {
		total += num * list2[num]
	}

	return total
}
