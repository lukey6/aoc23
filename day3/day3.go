package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var exampleInput = "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..\n"

func main() {

	lines := parseInputLines(exampleInput)

	fmt.Println("Part 1:", calculatePart1(lines))
	fmt.Println("Part 2:", calculatePart2(lines))
}

type gear struct {
	line            int
	column          int
	numberOfNumbers int
	number1         number
	number2         number
}

func calculatePart2(lines []string) int {

	numbers := findNumbers(lines)

	grs := gears(numbers, lines)

	return multiplyAndAddGears(grs)
}

func gears(nrs []*number, lines []string) []*gear {
	grs := make([]*gear, 0)

	for i := range nrs {
		nr := nrs[i]

		gr, ok := adjacentGear(nr, lines)
		if ok {
			if !gearExists(gr, grs) {
				gr.number1 = *nr
				grs = append(grs, gr)
			} else {
				if gr.numberOfNumbers == 1 {
					gr.number2 = *nr
				}
				gr.numberOfNumbers++
			}
		}
	}
	return grs
}

func multiplyAndAddGears(grs []*gear) int {
	sum := 0
	for i := range grs {
		gr := grs[i]
		if gr != nil {
			if gr.numberOfNumbers == 2 {
				sum += gr.number1.number * gr.number2.number
			}
		}
	}
	return sum
}

// returns adjacent gear and true if gear is adjacent to nr or nil and false if not
func adjacentGear(nr *number, lines []string) (*gear, bool) {
	if nr == nil {
		return nil, false
	}

	// left same line
	if nr.start > 0 && isGear(lines[nr.line][nr.start-1]) {
		return &gear{
			line:   nr.line,
			column: nr.start - 1,
		}, true
	}

	// right same line
	if nr.end < len(lines[nr.line])-1 && isGear(lines[nr.line][nr.end+1]) {
		return &gear{
			line:   nr.line,
			column: nr.start - 1,
		}, true
	}

	// line below
	for k := min(len(lines), nr.start+1); k < min(len(lines), nr.end+1); k++ {
		if isGear(lines[nr.line][k]) {
			return &gear{
				line:   nr.line,
				column: nr.start - 1,
			}, true
		}
	}

	// line above
	for k := max(0, nr.start-1); k < min(len(lines[nr.line]), nr.end+1); k++ {
		if isGear(lines[nr.line][k]) {
			return &gear{
				line:   nr.line,
				column: nr.start - 1,
			}, true
		}
	}

	//TODO...

	return nil, false
}

func gearExists(gr *gear, grs []*gear) bool {
	for i := range grs {
		if gr != nil && gr.line == grs[i].line && gr.column == grs[i].column {
			return true
		}
	}
	return false
}

func findNumbers(lines []string) []*number {
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

	return numbers
}

type number struct {
	line   int
	start  int
	end    int
	number int
	add    bool
}

func calculatePart1(lines []string) int {

	numbers := findNumbers(lines)

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

func isGear(c uint8) bool {
	return c == uint8('*')
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
