package day09

import (
	"fmt"
	"image"
	"strings"
)

func part1(input string) (int, error) {
	var result int

	lines := strings.Split(strings.TrimSpace(input), "\n")

	rope := make([]image.Point, 2)
	part1 := map[image.Point]struct{}{}
	dirs := map[rune]image.Point{'U': {0, -1}, 'R': {1, 0}, 'D': {0, 1}, 'L': {-1, 0}}

	for _, line := range lines {
		var direction rune
		var quantity int

		fmt.Sscanf(line, "%c %d", &direction, &quantity)

		move := dirs[direction]

		for i := 0; i < quantity; i++ {
			rope[0] = rope[0].Add(move)
			for i := 1; i < len(rope); i++ {
				diff := rope[i-1].Sub(rope[i])
				if abs(diff.X) > 1 || abs(diff.Y) > 1 {
					rope[i] = rope[i].Add(image.Point{sgn(diff.X), sgn(diff.Y)})
				}
				part1[rope[1]] = struct{}{}
			}
		}
	}

	result = len(part1)

	return result, nil
}

func part2(input string) (int, error) {
	var result int

	lines := strings.Split(strings.TrimSpace(input), "\n")

	rope := make([]image.Point, 10)
	part2 := map[image.Point]struct{}{}
	dirs := map[rune]image.Point{'U': {0, -1}, 'R': {1, 0}, 'D': {0, 1}, 'L': {-1, 0}}

	for _, line := range lines {
		var direction rune
		var quantity int

		fmt.Sscanf(line, "%c %d", &direction, &quantity)

		move := dirs[direction]

		for i := 0; i < quantity; i++ {
			rope[0] = rope[0].Add(move)
			for i := 1; i < len(rope); i++ {
				diff := rope[i-1].Sub(rope[i])
				if abs(diff.X) > 1 || abs(diff.Y) > 1 {
					rope[i] = rope[i].Add(image.Point{sgn(diff.X), sgn(diff.Y)})
				}
				part2[rope[len(rope)-1]] = struct{}{}
			}
		}
	}

	result = len(part2)

	return result, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sgn(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}
