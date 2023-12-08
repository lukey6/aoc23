package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
	"time"
)

var exampleInput1 = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"

var cardRegex = regexp.MustCompile("^Card' '+[0-9]+:$")

func main() {
	lines := parseInputLines(loadInput())

	fmt.Println("Part 1:", day4part1(lines))
}

func day4part1(lines []string) int {
	start := time.Now()
	defer fmt.Println("Runtime:", time.Since(start))

	points := 0

	for _, line := range lines {
		numberString := cardRegex.ReplaceAllString(line, "")

		splitNumbers := strings.Split(numberString, "|")

		winning := splitNumbers[0]
		have := splitNumbers[1]

		winning = strings.TrimSpace(winning)
		have = strings.TrimSpace(have)

		winningNumbers := strings.Split(winning, " ")
		haveNumbers := strings.Split(have, " ")

		wins := 0

		for _, w := range winningNumbers {
			for _, h := range haveNumbers {
				if w == h {
					wins++
					break
				}
			}
		}

		points += int(math.Pow(2, float64(wins-1)))
	}

	return points
}

func parseInputLines(input string) []string {
	return strings.Split(input, "\n")
}

func loadInput() string {
	inputBytes, err := os.ReadFile("./input")
	if err != nil {
		panic(err.Error())
	}

	return string(inputBytes)
}
