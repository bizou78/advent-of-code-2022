package day02

import (
	"testing"
	"github.com/bizou78/advent-of-code-2022/internal"
)

const (
	sample1 = `
A Y
B X
C Z
`

	part1Sample = 15
	part1Answer = 13009

	part2Sample = 12
	part2Answer = 10398
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{"sample 1", sample1, part1Sample},
		{"puzzle input", internal.ReadInput(t, "./testdata/input.txt"), part1Answer},
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
		{"sample 1", sample1, part2Sample},
		{"puzzle input", internal.ReadInput(t, "./testdata/input.txt"), part2Answer},
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
