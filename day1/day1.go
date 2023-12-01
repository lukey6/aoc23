package main

import (
	"fmt"
	"os"
	"strings"
)

var exampleInputPart1 string = "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"
var exampleInputPart2 string = "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen\n"

func main() {

	inputBytes, err := os.ReadFile("./input")
	if err != nil {
		panic("could not open input file")
	}

	inputString := string(inputBytes)

	input := strings.Split(inputString, "\n")

	sum := calculateSum(input)

	fmt.Println(sum)

}

func calculateSum(inputStrings []string) int {
	sum := 0

	for index := range inputStrings {
		input := inputStrings[index]
		currentString := replace(input)
		number := ExtractNumber(currentString)
		sum += number
	}
	return sum
}
