package main

import (
	"strconv"
	"strings"
)

const digitsString = "0123456789"

// Takes input string and returns number whose digits are first and last encountered digit in that order
func ExtractNumber(input string) int {
	length := len(input)

	var first int
	var last int

	for i := 0; i < length; i++ {
		char := rune(input[i])
		if strings.ContainsRune(digitsString, char) {
			first = i
			break
		}
	}

	for i := length - 1; i >= 0; i-- {
		char := rune(input[i])
		if strings.ContainsRune(digitsString, char) {
			last = i
			break
		}
	}

	num, err := strconv.ParseInt(string([]byte{input[first], input[last]}), 10, 32)
	if err != nil {
		//TODO
	}
	return int(num)
}
