package main

import (
	"fmt"
	"testing"
)

var testCases map[int][]string = make(map[int][]string)

func TestMain(m *testing.M) {
	testCases[163] = []string{"eightwo", "eightwone892seven1"}
	testCases[12] = []string{"oneightwo"}
	m.Run()
}

func TestRun(t *testing.T) {
	for k, v := range testCases {
		sum := calculateSum(v)
		if k != sum {
			fmt.Printf("expected: %v, real: %v, input: %s\n", k, sum, v)
			t.FailNow()
		}
	}
}
