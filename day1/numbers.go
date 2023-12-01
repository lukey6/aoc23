package main

import (
	"strings"
)

var digitNames = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

// replaces spelled out digits with the digit strings
func replace(input string) string {
	for i, dn := range digitNames {
		if strings.Contains(input, dn) {
			switch dn {
			case "one":
				input = strings.Replace(input, digitNames[i], "one1one", -1)
			case "two":
				input = strings.Replace(input, digitNames[i], "two2two", -1)
			case "three":
				input = strings.Replace(input, digitNames[i], "three3three", -1)
			case "four":
				input = strings.Replace(input, digitNames[i], "four4four", -1)
			case "five":
				input = strings.Replace(input, digitNames[i], "five5four", -1)
			case "six":
				input = strings.Replace(input, digitNames[i], "six6six", -1)
			case "seven":
				input = strings.Replace(input, digitNames[i], "seven7seven", -1)
			case "eight":
				input = strings.Replace(input, digitNames[i], "eight8eight", -1)
			case "nine":
				input = strings.Replace(input, digitNames[i], "nine9nine", -1)
			}
		}
	}
	return input
}
