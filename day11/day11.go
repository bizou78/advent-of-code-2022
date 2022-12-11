package day11

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

const RANGE = 40

type Monkey struct {
	Items          []int
	OperationType  string
	OperationValue int
	DivisibleBy    int
	TrueValue      int
	FalseValue     int
}

func part1(input string) (int, error) {
	var result int = 0
	const ROUNDS = 20

	lines := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	var monkeys = make([]Monkey, len(lines))

	// INIT ALL MONKEYS OBJECT WITH USEFULL PROPERTIES
	for _, line := range lines {
		var currentMonkey, divisbleBy, operationValue, trueThrowTo, falseThrowTo int
		var startItems, operation string

		fmt.Sscanf(strings.NewReplacer(", ", ",", "* old", "^ 2").Replace(line),
			`Monkey %d:
	Starting items: %s
	Operation: new = old %s %d
	Test: divisible by %d
		If true: throw to monkey %d
		If false: throw to monkey %d
`,
			&currentMonkey, &startItems, &operation, &operationValue, &divisbleBy, &trueThrowTo, &falseThrowTo)

		json.Unmarshal([]byte("["+startItems+"]"), &monkeys[currentMonkey].Items)

		monkeys[currentMonkey].OperationType = operation
		monkeys[currentMonkey].OperationValue = operationValue
		monkeys[currentMonkey].DivisibleBy = divisbleBy
		monkeys[currentMonkey].TrueValue = trueThrowTo
		monkeys[currentMonkey].FalseValue = falseThrowTo
	}

	//NOW WE NEED TO LOOP OVER ROUNDS NUMBER
	inspected := make([]int, len(monkeys))
	for i := 0; i < ROUNDS; i++ {
		for k, monkey := range monkeys {
			for _, item := range monkey.Items {
				newWorry := processOperation(monkey.OperationType, item, monkey.OperationValue) / 3
				if newWorry%monkey.DivisibleBy == 0 {
					monkeys[monkey.TrueValue].Items = append(monkeys[monkey.TrueValue].Items, newWorry)
				} else {
					monkeys[monkey.FalseValue].Items = append(monkeys[monkey.FalseValue].Items, newWorry)
				}
				inspected[k] += 1
			}
			monkeys[k].Items = nil
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspected)))
	result = inspected[0] * inspected[1]

	return result, nil
}

func part2(input string) (int, error) {
	var result int = 0
	const ROUNDS = 10000
	var PGCD int = 1
	lines := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	var monkeys = make([]Monkey, len(lines))

	// INIT ALL MONKEYS OBJECT WITH USEFULL PROPERTIES
	for _, line := range lines {
		var currentMonkey, divisbleBy, operationValue, trueThrowTo, falseThrowTo int
		var startItems, operation string

		fmt.Sscanf(strings.NewReplacer(", ", ",", "* old", "^ 2").Replace(line),
			`Monkey %d:
	Starting items: %s
	Operation: new = old %s %d
	Test: divisible by %d
		If true: throw to monkey %d
		If false: throw to monkey %d
`,
			&currentMonkey, &startItems, &operation, &operationValue, &divisbleBy, &trueThrowTo, &falseThrowTo)

		json.Unmarshal([]byte("["+startItems+"]"), &monkeys[currentMonkey].Items)

		monkeys[currentMonkey].OperationType = operation
		monkeys[currentMonkey].OperationValue = operationValue
		monkeys[currentMonkey].DivisibleBy = divisbleBy
		monkeys[currentMonkey].TrueValue = trueThrowTo
		monkeys[currentMonkey].FalseValue = falseThrowTo
		PGCD *= divisbleBy
	}
	//NOW WE NEED TO LOOP OVER ROUNDS NUMBER
	inspected := make([]int, len(monkeys))
	for i := 0; i < ROUNDS; i++ {
		for k, monkey := range monkeys {
			for _, item := range monkey.Items {
				newWorry := processOperation(monkey.OperationType, item, monkey.OperationValue) % PGCD

				if newWorry%monkey.DivisibleBy == 0 {
					monkeys[monkey.TrueValue].Items = append(monkeys[monkey.TrueValue].Items, newWorry)
				} else {
					monkeys[monkey.FalseValue].Items = append(monkeys[monkey.FalseValue].Items, newWorry)
				}
				inspected[k] += 1
			}
			monkeys[k].Items = nil
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspected)))
	result = inspected[0] * inspected[1]

	return result, nil
}

func processOperation(operation string, oldValue int, operationValue int) int {
	returnValue := 0
	switch operation {
	case "+":
		returnValue = oldValue + operationValue
		break

	case "*":
		returnValue = oldValue * operationValue
		break
	case "^":
		returnValue = oldValue * oldValue
		break
	}

	return returnValue
}
