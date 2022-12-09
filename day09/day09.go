package day09

import (
	"fmt"
	"image"
	"strings"
)

func part1(input string) (int, error) {
	var result int

	lines := strings.Split(strings.TrimSpace(input), "\n")

	var currentTail image.Point
	var currentHead image.Point

	tailMap := []image.Point{}

	tailMap = append(tailMap, currentTail)

	for _, line := range lines {
		var direction string
		var quantity int

		fmt.Sscanf(line, "%s %d", &direction, &quantity)

		for i := 0; i < quantity; i++ {

			switch direction {

			case "R":
				tmpCurrentHead := currentHead
				currentHead = image.Point{currentHead.X + 1, currentHead.Y}
				if currentHead.X-currentTail.X > 1 {
					currentTail = tmpCurrentHead
				}
				break

			case "L":
				tmpCurrentHead := currentHead
				currentHead = image.Point{currentHead.X - 1, currentHead.Y}
				if currentTail.X-currentHead.X > 1 {
					currentTail = tmpCurrentHead
				}
				break

			case "U":
				tmpCurrentHead := currentHead
				currentHead = image.Point{currentHead.X, currentHead.Y + 1}
				if currentHead.Y-currentTail.Y > 1 {
					currentTail = tmpCurrentHead
				}
				break

			case "D":
				tmpCurrentHead := currentHead
				currentHead = image.Point{currentHead.X, currentHead.Y - 1}
				if currentHead.Y-currentTail.Y < -1 {
					currentTail = tmpCurrentHead
				}
				break
			}

			tailMap = append(tailMap, currentTail)
		}

	}

	uniqueMap := map[image.Point]int{}

	for _, pair := range tailMap {
		uniqueMap[pair] = 1
	}

	result = len(uniqueMap)

	return result, nil
}

func part2(input string) (int, error) {
	var result int

	lines := strings.Split(strings.TrimSpace(input), "\n")

	dirs := map[rune]image.Point{'U': {0, -1}, 'R': {1, 0}, 'D': {0, 1}, 'L': {-1, 0}}
	rope := make([]image.Point, 10)

	part2 := map[image.Point]struct{}{}

	for _, s := range lines {
		var dir rune
		var steps int
		fmt.Sscanf(s, "%c %d", &dir, &steps)

		for i := 0; i < steps; i++ {
			rope[0] = rope[0].Add(dirs[dir])

			for i := 1; i < len(rope); i++ {
				if d := rope[i-1].Sub(rope[i]); abs(d.X) > 1 || abs(d.Y) > 1 {
					rope[i] = rope[i].Add(image.Point{sgn(d.X), sgn(d.Y)})
				}
			}
			part2[rope[len(rope)-1]] = struct{}{}
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
