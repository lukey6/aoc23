package main

import (
	"strings"
)

var digitNames = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "zero"}

// replaces spelled out digits with the digit strings
func replace(input string) string {
	for i, dn := range digitNames {
		if strings.Contains(input, dn) {
			switch dn {
			case "one":
				input = strings.Replace(input, digitNames[i], "one1ne", -1)
			case "two":
				input = strings.Replace(input, digitNames[i], "two2", -1)
			case "three":
				input = strings.Replace(input, digitNames[i], "three3", -1)
			case "four":
				input = strings.Replace(input, digitNames[i], "four4", -1)
			case "five":
				input = strings.Replace(input, digitNames[i], "five5", -1)
			case "six":
				input = strings.Replace(input, digitNames[i], "six6", -1)
			case "seven":
				input = strings.Replace(input, digitNames[i], "seven7", -1)
			case "eight":
				input = strings.Replace(input, digitNames[i], "eight8", -1)
			case "nine":
				input = strings.Replace(input, digitNames[i], "nine9", -1)
			}
		}
	}
	return input
}
