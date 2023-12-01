package main

import (
	"fmt"
	"os"
	"strings"
)

var exampleInput string = "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"

func main() {

	inputBytes, err := os.ReadFile("./input")
	if err != nil {
		panic("could not open input file")
	}

	inputString := string(inputBytes)

	input := strings.Split(inputString, "\n")

	sum := 0

	for index := range input {
		number := ExtractNumber(input[index])

		sum += number
	}

	fmt.Println(sum)

}
