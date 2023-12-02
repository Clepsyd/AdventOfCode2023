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

func getEmptyValues() map[string]int {
	return map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
}

func parseSet(s string) map[string]int {
	reader := csv.NewReader(strings.NewReader(s))
	result := getEmptyValues()

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

func getMaxDiceCountByColorForGame(sets []string) map[string]int {
	result := getEmptyValues()
	for _, set := range sets {
		setResult := parseSet(set)
		for color, count := range setResult {
			if count > result[color] {
				result[color] = count
			}
		}
	}

	return result
}

func main() {
	f, err := os.Open("input.txt")
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var total1 int
	var total2 int

	for scanner.Scan() {
		line := scanner.Text()
		gameData := strings.Split(line, SPLIT_GAME_DATA)
		gameID, sets :=
			strings.Split(gameData[0], SPLIT_GAME_ID)[1],
			strings.Split(strings.TrimSpace(gameData[1]), SPLIT_SETS)
		fmt.Println(strings.Join(sets, " | "))

		gameIsFeasible := isGameFeasible(sets)
		var feasibilityDisplay string
		if gameIsFeasible {
			gameValue, err := strconv.Atoi(gameID)
			check(err)
			total1 += gameValue
			feasibilityDisplay = "✅"
		} else {
			feasibilityDisplay = "❌"
		}

        powerOfMinNumbersOfDice := 1
        minDiceRequiredByColor := getMaxDiceCountByColorForGame(sets)
        for _, numberOfDice := range minDiceRequiredByColor {
            powerOfMinNumbersOfDice *= numberOfDice
        }
        total2 += powerOfMinNumbersOfDice

        fmt.Println(
            "Feasible: ", feasibilityDisplay,
            "ID: ", gameID,
            "MinDiceByColor: ", minDiceRequiredByColor,
            "powerOfMinNumbersOfDice:", powerOfMinNumbersOfDice,
        )
        fmt.Println(
            "total1: ", total1,
            "total2: ", total2,
        )
        fmt.Println()
	}
}
