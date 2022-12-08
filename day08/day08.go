package day08

import (
	"strings"
)

func part1(input string) (int, error) {

	var result int

	lines := strings.Split(strings.TrimSpace(input), "\n")

	stock1 := []map[int]bool{}
	stock2 := []map[int]bool{}
	stock3 := []map[int]bool{}
	stock4 := []map[int]bool{}

	result = len(lines)*2 + len(lines[0])*2 - 4

	// UP TEST = true, true, false | false, false, false | false, false, false
	for i := 1; i < len(lines)-1; i++ {
		stockLine := map[int]bool{}
		for j := 1; j < len(lines[i])-1; j++ {
			for k := i - 1; k >= 0; k-- {
				if val, ok := stockLine[j]; ok {
					if lines[i][j] > lines[k][j] && val != false {
						stockLine[j] = true
					} else {
						stockLine[j] = false
					}
				} else {
					if lines[i][j] > lines[k][j] {
						stockLine[j] = true
					} else {
						stockLine[j] = false
						break
					}
				}
			}
		}
		stock1 = append(stock1, stockLine)
	}

	// DOWN TEST = false, false, false | false, false, false | false, true, false
	for i := 1; i < len(lines)-1; i++ {
		stockLine := map[int]bool{}
		for j := 1; j < len(lines[i])-1; j++ {
			for k := i + 1; k < len(lines); k++ {
				if val, ok := stockLine[j]; ok {
					if lines[i][j] > lines[k][j] && val != false {
						stockLine[j] = true
					} else {
						stockLine[j] = false
					}
				} else {
					if lines[i][j] > lines[k][j] {
						stockLine[j] = true
					} else {
						stockLine[j] = false
					}
				}
			}
		}
		stock2 = append(stock2, stockLine)
	}

	// LEFT TEST = false, true, false | false, false, false | false, true, false
	for i := 1; i < len(lines)-1; i++ {
		stockLine := map[int]bool{}
		for j := 1; j < len(lines[i])-1; j++ {
			for k := j - 1; k >= 0; k-- {
				if val, ok := stockLine[j]; ok {
					if lines[i][j] > lines[i][k] && val != false {
						stockLine[j] = true
					} else {
						stockLine[j] = false
					}
				} else {
					if lines[i][j] > lines[i][k] {
						stockLine[j] = true
					} else {
						stockLine[j] = false
					}
				}
			}
		}
		stock3 = append(stock3, stockLine)
	}
	// RIGHT TEST = false, true, false | true, false, true | false, false, false
	for i := 1; i < len(lines)-1; i++ {
		stockLine := map[int]bool{}
		for j := 1; j < len(lines[i])-1; j++ {
			for k := j + 1; k < len(lines[i]); k++ {
				if val, ok := stockLine[j]; ok {
					if lines[i][j] > lines[i][k] && val != false {
						stockLine[j] = true
					} else {
						stockLine[j] = false
					}
				} else {
					if lines[i][j] > lines[i][k] {
						stockLine[j] = true
					} else {
						stockLine[j] = false
					}
				}
			}
		}
		stock4 = append(stock4, stockLine)
	}

	for i := 0; i < len(stock1); i++ {
		for j := 1; j < len(stock1[i])+1; j++ {
			if stock1[i][j] == true || stock2[i][j] == true || stock3[i][j] == true || stock4[i][j] == true {
				result += 1
			}
		}
	}

	return result, nil
}

func part2(input string) (int, error) {

	var result int
	var countTrees int
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var viewUp, viewDown, viewLeft, viewRight int
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			// VIEW UP
			viewUp = 1
			for k := i - 1; k >= 0; k-- {
				if lines[i][j] > lines[k][j] {
					viewUp += 1
				} else {
					break
				}
			}
			// VIEW DOWN
			viewDown = 1
			for k := i + 1; k < len(lines)-1; k++ {
				if lines[i][j] > lines[k][j] {
					viewDown += 1
				} else {
					break
				}
			}
			// VIEW LEFT
			viewLeft = 1
			for k := j - 1; k > 0; k-- {
				if lines[i][j] > lines[i][k] {
					viewLeft += 1
				} else {
					break
				}
			}
			//VIEW RIGHT
			viewRight = 1
			for k := j + 1; k < len(lines[i])-1; k++ {
				if lines[i][j] > lines[i][k] {
					viewRight += 1
				} else {
					break
				}

			}
			countTrees = viewUp * viewDown * viewLeft * viewRight
			if countTrees > result {
				result = countTrees
			}
		}
	}

	return result, nil
}
