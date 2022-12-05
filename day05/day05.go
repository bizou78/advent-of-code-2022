package day05

import (
	"fmt"
	"regexp"
	"strings"
)

// A = 65 -> Z = 90
func part1(input string) (string, error) {

	var result string
	var keyNumber int
	stock := map[rune]string{}

	lines := strings.Split(input, "\n\n")
	var numberOfKeys string

	uniqueLineArray := strings.Split(lines[0], "\n")

	for _, uniqueLine := range uniqueLineArray {

		regexMatch, _ := regexp.MatchString(`\[|\]`, uniqueLine)
		if regexMatch {
			for i, ch := range uniqueLine {
				if ch >= 65 && ch <= 90 {
					if i == 1 {
						stock[1] += string(ch)
					} else if i == 5 {
						stock[2] += string(ch)
					} else if i == 9 {
						stock[3] += string(ch)
					} else if i == 13 {
						stock[4] += string(ch)
					} else if i == 17 {
						stock[5] += string(ch)
					} else if i == 21 {
						stock[6] += string(ch)
					} else if i == 25 {
						stock[7] += string(ch)
					} else if i == 29 {
						stock[8] += string(ch)
					} else if i == 33 {
						stock[9] += string(ch)
					}
				}
			}
		} else if uniqueLine != "" {
			numberOfKeys = strings.TrimSpace(uniqueLine)
		}
	}
	// INSTRUCTIONS
	instructionLines := strings.Split(lines[1], "\n")
	for _, instruction := range instructionLines {
		if instruction != "" {
			var quantity int
			var from, to rune

			fmt.Sscanf(instruction, "move %d from %d to %d", &quantity, &from, &to)

			for i := 0; i < quantity; i++ {
				stock[to] = stock[from][:1] + stock[to]
				stock[from] = stock[from][1:]
			}
		}
	}
	for _, ch := range numberOfKeys {

		if ch >= 49 && ch <= 57 {
			keyNumber++
		}
	}

	for i := 1; i <= keyNumber; i++ {
		iString := rune(i)
		result += string(stock[iString][0])
	}

	return result, nil
}

func part2(input string) (string, error) {

	var result string
	var keyNumber int
	stock := map[rune]string{}

	lines := strings.Split(input, "\n\n")
	var numberOfKeys string

	uniqueLineArray := strings.Split(lines[0], "\n")

	for _, uniqueLine := range uniqueLineArray {

		regexMatch, _ := regexp.MatchString(`\[|\]`, uniqueLine)
		if regexMatch {
			for i, ch := range uniqueLine {
				if ch >= 65 && ch <= 90 {
					if i == 1 {
						stock[1] += string(ch)
					} else if i == 5 {
						stock[2] += string(ch)
					} else if i == 9 {
						stock[3] += string(ch)
					} else if i == 13 {
						stock[4] += string(ch)
					} else if i == 17 {
						stock[5] += string(ch)
					} else if i == 21 {
						stock[6] += string(ch)
					} else if i == 25 {
						stock[7] += string(ch)
					} else if i == 29 {
						stock[8] += string(ch)
					} else if i == 33 {
						stock[9] += string(ch)
					}
				}
			}
		} else if uniqueLine != "" {
			numberOfKeys = strings.TrimSpace(uniqueLine)
		}
	}
	fmt.Println("Stock before : ", stock)
	// INSTRUCTIONS
	instructionLines := strings.Split(lines[1], "\n")
	for _, instruction := range instructionLines {
		if instruction != "" {
			var quantity int
			var from, to rune

			fmt.Sscanf(instruction, "move %d from %d to %d", &quantity, &from, &to)

			stock[to] = stock[from][:quantity] + stock[to]
			stock[from] = stock[from][quantity:]
		}
	}
	fmt.Println("Stock after : ", stock)
	for _, ch := range numberOfKeys {

		if ch >= 49 && ch <= 57 {
			keyNumber++
		}
	}

	for i := 1; i <= keyNumber; i++ {
		iString := rune(i)
		result += string(stock[iString][0])
	}

	return result, nil
}
