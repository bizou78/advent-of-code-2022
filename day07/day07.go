package day07

import (
	"fmt"
	"path"
	"strings"
	"unicode"
)

func part1(input string) (int, error) {

	var result int
	var fileSystem = map[string]int{}
	var currentDirectory string

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {

		if strings.HasPrefix(line, "$ cd") {
			currentDirectory = path.Join(currentDirectory, strings.Fields(line)[2])
		} else if unicode.IsDigit([]rune(line)[0]) {
			var fileSize int
			var fileName string

			fmt.Sscanf(line, "%d %s", &fileSize, &fileName)

			fileSystem[path.Join(currentDirectory, fileName)] = fileSize

		}
	}

	directorySize := map[string]int{}

	for completePath, size := range fileSystem {
		for directory := path.Dir(completePath); directory != "/"; directory = path.Dir(directory) {
			directorySize[directory] += size
		}
		directorySize["/"] += size
	}

	for _, size := range directorySize {
		if size <= 100000 {
			result += size
		}
	}

	return result, nil
}

func part2(input string) (int, error) {

	var fileSystem = map[string]int{}
	var currentDirectory string

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {

		if strings.HasPrefix(line, "$ cd") {
			currentDirectory = path.Join(currentDirectory, strings.Fields(line)[2])
		} else if unicode.IsDigit([]rune(line)[0]) {
			var fileSize int
			var fileName string

			fmt.Sscanf(line, "%d %s", &fileSize, &fileName)

			fileSystem[path.Join(currentDirectory, fileName)] = fileSize

		}
	}

	directorySize := map[string]int{}

	for completePath, size := range fileSystem {
		for directory := path.Dir(completePath); directory != "/"; directory = path.Dir(directory) {
			directorySize[directory] += size
		}
		directorySize["/"] += size
	}

	result := directorySize["/"]
	for _, size := range directorySize {
		if size+70000000-directorySize["/"] >= 30000000 && size < result {
			result = size
		}
	}

	return result, nil
}
