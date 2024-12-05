package part1

import (
	"regexp"
	"silverark/aoc-2024/pkg/shared"
)

func process(input []string) int {
	total := 0
	var mulEnabled = true
	for _, line := range input {
		r := regexp.MustCompile("mul\\(([0-9]+),([0-9]+)\\)|don't\\(\\)|do\\(\\)")
		matches := r.FindAllStringSubmatch(line, -1)
		for i := 0; i < len(matches); i++ {
			if matches[i][0] == "do()" {
				mulEnabled = true
				continue
			}
			if matches[i][0] == "don't()" {
				mulEnabled = false
				continue
			}
			if mulEnabled {
				total += shared.Atoi(matches[i][1]) * shared.Atoi(matches[i][2])
			}
		}
	}
	return total
}
