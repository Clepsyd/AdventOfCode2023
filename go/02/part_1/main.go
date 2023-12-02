package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var colors = [...]string{"red", "green", "blue"}

const SPLIT_GAME_DATA = ":"
const SPLIT_GAME_ID = " "
const SPLIT_SETS = ";"

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

var maxValues = map[string]int{
	"red":   MAX_RED,
	"green": MAX_GREEN,
	"blue":  MAX_BLUE,
}

func parseSet(s string) map[string]int {
	reader := csv.NewReader(strings.NewReader(s))
	result := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	record, err := reader.Read()
	check(err)

	for _, item := range record {
		split := strings.Split(strings.TrimSpace(item), " ")
		count, color := strings.TrimSpace(split[0]), strings.TrimSpace(split[1])
		result[color], err = strconv.Atoi(count)
		check(err)
	}


	return result
}

func isGameFeasible(sets []string) bool {
	for _, set := range sets {
		result := parseSet(set)
		for color, count := range result {
			if count > maxValues[color] {
				return false
			}
		}
	}

	return true
}

func main() {
	f, err := os.Open("input.txt")
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var total int

	for scanner.Scan() {
		line := scanner.Text()
		gameData := strings.Split(line, SPLIT_GAME_DATA)
		fmt.Println(gameData)
		gameID, sets :=
			strings.Split(gameData[0], SPLIT_GAME_ID)[1],
			strings.Split(strings.TrimSpace(gameData[1]), SPLIT_SETS)

		gameIsFeasible := isGameFeasible(sets)
        var feasibilityDisplay string
		if gameIsFeasible {
			gameValue, err := strconv.Atoi(gameID)
			check(err)
			total += gameValue
            feasibilityDisplay = "✅"
		} else {
            feasibilityDisplay = "❌"
        }

        fmt.Println(feasibilityDisplay, gameID, "total: ", total)
	}
}
