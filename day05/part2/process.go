package part1

import (
	"silverark/aoc-2024/pkg/shared"
	"slices"
	"strings"
)

func process(input string) int {
	middleTotal := 0
	data := strings.Split(input, "\n\n")
	ruleData := data[0]
	updateData := data[1]
	rules := make(map[int][]int)
	for _, r := range strings.Split(ruleData, "\n") {
		items := strings.Split(r, "|")
		a, b := shared.Atoi(items[0]), shared.Atoi(items[1])
		rules[a] = append(rules[a], b)
	}

	cmp := func(a, b int) int {
		// Make sure the second number doesn't have the first in its list
		if slices.Contains(rules[b], a) {
			return 1
		}
		return -1
	}

	for _, update := range strings.Split(updateData, "\n") {
		updateStrings := strings.Split(update, ",")
		var updatePages []int
		for _, u := range updateStrings {
			updatePages = append(updatePages, shared.Atoi(u))
		}
		// Only get the incorrectly ordered ones and order them
		if !slices.IsSortedFunc(updatePages, cmp) {
			// Sort the slice
			slices.SortFunc(updatePages, cmp)
			middleTotal += updatePages[len(updatePages)/2]
		}
	}
	return middleTotal
}
