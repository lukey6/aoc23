package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var maxNumbers = make(map[string]int64)

var exampleInput1 = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"

func main() {
	initialize()

	games := splitInput(openInput())

	fmt.Println(processPossibleGames(games))

}

func initialize() {
	maxNumbers["red"] = int64(12)
	maxNumbers["green"] = int64(13)
	maxNumbers["blue"] = int64(14)
}

func openInput() string {
	inputBytes, err := os.ReadFile("./input")
	if err != nil {
		panic("could not open input file")
	}

	return string(inputBytes)
}

// getss input string and returns map [gameID]gamestrings
func splitInput(input string) map[int][]string {
	games := strings.Split(input, "\n")
	gameMap := make(map[int][]string)
	gameID := 1

	for i := range games {
		games[i] = strings.Trim(games[i], " ,;\n")
		gamesString, found := strings.CutPrefix(games[i], fmt.Sprintf("Game %v:", gameID))
		if !found {
			panic("error processing input")
		}

		draws := strings.FieldsFunc(gamesString, separators)

		gameMap[gameID] = draws
		gameID++
	}

	return gameMap
}

func separators(r rune) bool {
	return r == ',' || r == ';'
}

func processPossibleGames(gamesMap map[int][]string) int {

	sum := 0
	k := 0
	for i := 0; i <= 100; i++ {
		k += i
	}

	fmt.Printf("k: %v\n", k)

	for id, draws := range gamesMap {
		possible := true
		for j := range draws {

			draws[j] = strings.Trim(draws[j], " ,;\n")
			comps := strings.Split(draws[j], " ")

			color := comps[1]
			number, err := strconv.ParseInt(comps[0], 10, 64)
			if err != nil {
				panic(err.Error())
			}

			maxNr := maxNumbers[color]

			if maxNr < number {
				possible = false
				break
			}
		}

		if possible {
			sum += id
		}
	}

	return sum
}
