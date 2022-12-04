package day03

import (
	"strings"
)

// a = 97 -> z = 122 ||| a = 1 -> z = 26 ||| diff = 97 - 1 = 96
// A = 65 -> Z = 90 ||| A = 27 -> Z = 52 ||| diff = 65 - 27 = 38

func part1(input string) (int, error) {

	var total int
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		stockLetter := map[int]bool{}
		fullLength := len(line) - 1
		midString := len(line) / 2

		for i := 0; i < midString; i++ {
			charFirstLine := int(line[i])
			stockLetter[charFirstLine] = true
		}

		for j := fullLength; j >= midString; j-- {
			charInt := int(line[j])
			if stockLetter[charInt] == true {

				if charInt > 90 {
					thislowerValue := charInt - 96
					total += thislowerValue
				} else {
					thisUpperValue := charInt - 38
					total += thisUpperValue
				}
				break
			}
		}

	}

	return total, nil
}

func part2(input string) (int, error) {

	var total int
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for z := 0; z < len(lines); z += 3 { // 3 LIGNES

		stockLetter := [3][123]bool{}

		for k := 0; k < 3; k++ {
			for i := 0; i < len(lines[z+k]); i++ {
				charFirstLine := int(lines[z+k][i])
				stockLetter[k][charFirstLine] = true
			}
		}

		for j := 65; j < 128; j++ {

			if stockLetter[0][j] && stockLetter[1][j] && stockLetter[2][j] {

				if j > 90 {
					thislowerValue := j - 96
					total += thislowerValue
				} else {
					thisUpperValue := j - 38
					total += thisUpperValue
				}
				break
			}

		}

	}

	return total, nil
}
