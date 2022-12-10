package day10

import (
	"fmt"
	"strings"
)

const RANGE = 40

func part1(input string) (int, error) {
	var result int = 0
	var value int = 1
	var countCycle int = 0

	lines := strings.Split(strings.TrimSpace(input), "\n")

	tick := func() {
		countCycle++
		if (countCycle+RANGE/2)%RANGE == 0 {
			result += countCycle * value
		}
	}

	for _, line := range lines {

		var x int

		tick()

		fmt.Sscanf(line, "addx %d", &x)

		if strings.TrimSpace(line) != "noop" {
			tick()
			value += x
		}
	}

	return result, nil
}

func part2(input string) (string, error) {
	var result string = ""
	var value int = 1
	var countCycle int = 0

	lines := strings.Split(strings.TrimSpace(input), "\n")

	tick := func() {

		if countCycle%RANGE >= value-1 && countCycle%RANGE <= value+1 {
			result += "#"
		} else {
			result += "."
		}
		if countCycle%RANGE == RANGE-1 {
			result += "\n"
		}

		countCycle++
	}

	for _, line := range lines {

		var x int

		tick()

		fmt.Sscanf(line, "addx %d", &x)

		if strings.TrimSpace(line) != "noop" {
			tick()
			value += x
		}
	}

	result = strings.TrimSpace(result)

	fmt.Println("result : \n", result)

	return result, nil
}
