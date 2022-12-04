package day04

import (
	"fmt"
	"strings"
)

func part1(input string) (int, error) {

	var total int
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {

		fmt.Println("Line : ", line)

		assignement := strings.Split(line, ",")

		splitAssignement1 := strings.Split(assignement[0], "-")
		splitAssignement2 := strings.Split(assignement[1], "-")

		fmt.Println("Assignement 1 : ", splitAssignement1)
		fmt.Println("Assignement 2 : ", splitAssignement2)

		if splitAssignement1[0] >= splitAssignement2[0] && splitAssignement1[1] <= splitAssignement2[1] {
			total++
		} else if splitAssignement1[0] <= splitAssignement2[0] && splitAssignement1[1] >= splitAssignement2[1] {
			total++
		}

	}

	return total, nil
}

// func part2(input string) (int, error) {

// 	var total int
// 	lines := strings.Split(strings.TrimSpace(input), "\n")

// 	for z := 0; z < len(lines); z += 3 { // 3 LIGNES

// 		stockLetter := [3][123]bool{}

// 		for k := 0; k < 3; k++ {
// 			for i := 0; i < len(lines[z+k]); i++ {
// 				charFirstLine := int(lines[z+k][i])
// 				stockLetter[k][charFirstLine] = true
// 			}
// 		}

// 		for j := 65; j < 128; j++ {

// 			if stockLetter[0][j] && stockLetter[1][j] && stockLetter[2][j] {

// 				if j > 90 {
// 					thislowerValue := j - 96
// 					total += thislowerValue
// 				} else {
// 					thisUpperValue := j - 38
// 					total += thisUpperValue
// 				}
// 				break
// 			}

// 		}

// 	}

// 	return total, nil
// }
