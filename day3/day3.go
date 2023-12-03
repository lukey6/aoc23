package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var exampleInput1 = "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."

func main() {

	lines := parseInputLines(exampleInput1)

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

		fmt.Println(line)

		for j := range line {
			fmt.Println(j)
			if isDigit(line[j]) {
				if start == -1 {
					start = j
				}
			} else if start != -1 {
				end = j - 1

				nr, err := strconv.ParseInt(line[start:end+1], 10, 64)
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

		for i := range numbers {
			markAdjacentToSymbol(numbers[i], lines)
		}
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

func markAdjacentToSymbol(nr *number, lines []string) {
	if nr == nil {
		return
	}
	localLines := make([]string, len(lines))
	copy(localLines, lines)

	for j := nr.start - 1; j <= nr.end+1; j++ {
		if j < 0 {
			continue
		}
		if j > 0 && nr.line > 0 && !isDigitOrPoint(localLines[nr.line-1][j]) {
			nr.add = true
			break
		} else if !isDigitOrPoint(localLines[nr.line][j]) {
			nr.add = true
			break
		} else if nr.line < len(localLines)-2 && j < len(localLines[j])-1 && !isDigitOrPoint(localLines[nr.line+1][j]) {
			nr.add = true
			break
		}
	}
}

func isDigitOrPoint(c uint8) bool {
	return isDigit(c) || c == '.'
}

func isDigit(c uint8) bool {
	return unicode.IsDigit(rune(c)) && c != '.'
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
