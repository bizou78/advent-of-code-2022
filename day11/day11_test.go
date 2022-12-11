package day11

import (
	"testing"

	"github.com/bizou78/advent-of-code-2022/internal"
)

const (
	sample1 = `
Monkey 0:
	Starting items: 79, 98
  	Operation: new = old * 19
  	Test: divisible by 23
		If true: throw to monkey 2
		If false: throw to monkey 3

Monkey 1:
  	Starting items: 54, 65, 75, 74
  	Operation: new = old + 6
  	Test: divisible by 19
		If true: throw to monkey 2
		If false: throw to monkey 0

Monkey 2:
  	Starting items: 79, 60, 97
  	Operation: new = old * old
  	Test: divisible by 13
		If true: throw to monkey 1
		If false: throw to monkey 3

Monkey 3:
  	Starting items: 74
  	Operation: new = old + 3
  	Test: divisible by 17
		If true: throw to monkey 0
		If false: throw to monkey 1
`
	input1 = `
Monkey 0:
	Starting items: 63, 84, 80, 83, 84, 53, 88, 72
	Operation: new = old * 11
	Test: divisible by 13
		If true: throw to monkey 4
		If false: throw to monkey 7

Monkey 1:
	Starting items: 67, 56, 92, 88, 84
	Operation: new = old + 4
	Test: divisible by 11
		If true: throw to monkey 5
		If false: throw to monkey 3

Monkey 2:
	Starting items: 52
	Operation: new = old * old
	Test: divisible by 2
		If true: throw to monkey 3
		If false: throw to monkey 1

Monkey 3:
	Starting items: 59, 53, 60, 92, 69, 72
	Operation: new = old + 2
	Test: divisible by 5
		If true: throw to monkey 5
		If false: throw to monkey 6

Monkey 4:
	Starting items: 61, 52, 55, 61
	Operation: new = old + 3
	Test: divisible by 7
		If true: throw to monkey 7
		If false: throw to monkey 2

Monkey 5:
	Starting items: 79, 53
	Operation: new = old + 1
	Test: divisible by 3
		If true: throw to monkey 0
		If false: throw to monkey 6

Monkey 6:
	Starting items: 59, 86, 67, 95, 92, 77, 91
	Operation: new = old + 5
	Test: divisible by 19
		If true: throw to monkey 4
		If false: throw to monkey 0

Monkey 7:
	Starting items: 58, 83, 89
	Operation: new = old * 19
	Test: divisible by 17
		If true: throw to monkey 2
		If false: throw to monkey 1
`
	part1Sample = 10605
	part1Answer = 117640

	part2Sample = 2713310158
	part2Answer = 30616425600
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{"sample 1", sample1, part1Sample},
		{"puzzle input", input1, part1Answer},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := part1(tc.input)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if got != tc.want {
				t.Errorf("Got: %v, want: %v", got, tc.want)
			}
		})
	}

}

func TestPart2(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{"sample 2", sample1, part2Sample},
		{"puzzle input", input1, part2Answer},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := part2(tc.input)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if got != tc.want {
				t.Errorf("Got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	input := internal.ReadInput(b, "./testdata/input.txt")
	parts := []struct {
		name   string
		fn     func(input string) (int, error)
		answer int
	}{
		{"part1", part1, part1Answer},
		{"part2", part2, part2Answer},
	}

	for _, part := range parts {
		b.Run(part.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				got, err := part.fn(input)
				if err != nil {
					b.Errorf("Error: %v", err)
				}
				if got != part.answer {
					b.Errorf("Got: %v, want: %v", got, part.answer)
				}
			}
		})
	}
}
