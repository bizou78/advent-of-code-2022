package day02

import (
	"strings"
)

var (
	mapShape = map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}

	mapValue = map[string]int{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
	}

	winningMap = map[string]string{
		"rock":     "scissors",
		"paper":    "rock",
		"scissors": "paper",
	}

	losingMap = map[string]string{
		"rock":     "paper",
		"paper":    "scissors",
		"scissors": "rock",
	}
)

func part1(input string) (int, error) {
	rounds := strings.Split(strings.TrimSpace(input), "\n")

	myPoint := 0

	for _, round := range rounds {
		elements := strings.Split(round, " ")

		ennemyShape := mapShape[elements[0]]

		myShape := mapShape[elements[1]]
		myValue := mapValue[myShape]
		myPoint += myValue

		if ennemyShape == myShape {
			myPoint += 3
		} else if winningMap[myShape] == ennemyShape {
			myPoint += 6
		}

	}
	return myPoint, nil
}

func part2(input string) (int, error) {
	rounds := strings.Split(strings.TrimSpace(input), "\n")
	var myShape string

	myPoint := 0

	for _, round := range rounds {
		elements := strings.Split(round, " ")

		ennemyShape := mapShape[elements[0]]

		switch elements[1] {
		case "X":
			myShape = winningMap[ennemyShape]
		case "Y":
			myShape = ennemyShape
			myPoint += 3
		case "Z":
			myShape = losingMap[ennemyShape]
			myPoint += 6
		}

		myPoint += mapValue[myShape]

	}

	return myPoint, nil
}
