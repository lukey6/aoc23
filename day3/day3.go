package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var exampleInput1 = "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..\n"

func main() {

	lines := parseInputLines(loadInput())

	fmt.Println("Part 1:", calculatePart1(lines))
}

type number struct {
	line   int
	start  int
	end    int
	number int
	add    bool
}

func calculatePart1(lines []string) int {
	numbers := make([]*number, 1)

	for i := range lines {
		line := lines[i]

		start := -1
		end := -1

		for j := range line {
			if isDigit(line[j]) {
				// if line[j] is digit and start not marked, mark start
				if start == -1 {
					start = j
				}

				if j >= len(line)-1 || !isDigit(line[j+1]) {
					// if line[j] is last in line or the next one is no digit, mark end
					end = j
				}
			}

			if start != -1 && end != -1 {
				// if line[j] is no digit and start and end are marked save number, unmark start and end
				var nr int64
				var err error

				if j >= len(line) {
					// if end+1 is out of bounds take string until end of line
					nr, err = strconv.ParseInt(line[start:], 10, 64)
				} else {
					nr, err = strconv.ParseInt(line[start:end+1], 10, 64)
				}

				if err != nil {
					panic(err.Error())
				}

				numbers = append(numbers, &number{
					line:   i,
					start:  start,
					end:    end,
					number: int(nr),
					add:    false,
				})
				start = -1
				end = -1
			}
		}
	}

	for i := range numbers {
		markAdjacentToSymbol(numbers[i], lines)
	}

	return addMarked(numbers)
}

func addMarked(nrs []*number) int {
	sum := 0

	for i := range nrs {
		if nrs[i] != nil && nrs[i].add {
			sum += nrs[i].number
		}
	}
	return sum
}

// TODO result too low => missing nrs
func markAdjacentToSymbol(nr *number, lines []string) {
	if nr == nil {
		return
	}

	for j := nr.start - 1; j <= nr.end+1; j++ {
		if j < 0 {
			continue
		}
		if j >= 0 && nr.line > 0 && j < len(lines[nr.line-1]) && !isDigitOrPoint(lines[nr.line-1][j]) {
			nr.add = true
			break
		} else if j < len(lines[nr.line]) && !isDigitOrPoint(lines[nr.line][j]) {
			nr.add = true
			break
		} else if nr.line < len(lines)-1 && j < len(lines[nr.line+1])-1 && !isDigitOrPoint(lines[nr.line+1][j]) {
			nr.add = true
			break
		}
	}
}

func isDigitOrPoint(c uint8) bool {
	return isDigit(c) || c == '.'
}

func isDigit(c uint8) bool {
	return strings.ContainsRune("0123456789", rune(c))
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
