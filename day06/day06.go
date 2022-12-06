package day06

func part1(input string) (int, error) {

	var result int
	messageLength := 4

	for i := 0; i < len(input); i++ {
		stock := []string{}
		for j := i; j < i+messageLength; j++ {
			stringInput := string(input[j])
			stock = append(stock, stringInput)
		}
		uniqueStock := unique(stock)
		if len(uniqueStock) == messageLength {
			result = i + messageLength
			break
		}
	}

	return result, nil
}

func part2(input string) (int, error) {

	var result int
	messageLength := 14

	for i := 0; i < len(input); i++ {
		stock := []string{}
		for j := i; j < i+messageLength; j++ {
			stringInput := string(input[j])
			stock = append(stock, stringInput)
		}
		uniqueStock := unique(stock)
		if len(uniqueStock) == messageLength {
			result = i + messageLength
			break
		}
	}

	return result, nil
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
