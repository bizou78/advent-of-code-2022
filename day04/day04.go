package day04

import (
	"fmt"
	"strings"
)

func part1(input string) (int, error) {

	var total int
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {

		var s1, e1, s2, e2 int
		fmt.Sscanf(line, "%d-%d,%d-%d", &s1, &e1, &s2, &e2)

		if s1 <= s2 && e1 >= e2 || s1 >= s2 && e1 <= e2 {
			total++
		}

	}

	return total, nil
}

func part2(input string) (int, error) {

	var total int
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {

		var s1, e1, s2, e2 int
		fmt.Sscanf(line, "%d-%d,%d-%d", &s1, &e1, &s2, &e2)

		if s1 <= e2 && e1 >= s2 {
			total++
		}

	}

	return total, nil
}
