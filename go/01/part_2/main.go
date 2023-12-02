package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var stringsToDigits = map[string]string{
	"zero":  "0", // technically not in the provided input but you never know with those elves...
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var threeLetterDigits = []string{
	"one",
	"two",
	"six",
}

var fourLetterDigits = []string{
	"zero",
	"four",
	"five",
	"nine",
}
var fiveLetterDigits = []string{
	"three",
	"seven",
	"eight",
}

var wordsByLetterCount = [3][]string{
	threeLetterDigits,
	fourLetterDigits,
	fiveLetterDigits,
}

const SMALLEST_WORD_LENGTH = 3

func getFirst(line string) (string, error) {
	for i := range line {
		char := string(line[i])
		if _, err := strconv.Atoi(char); err == nil {
			return char, nil
		}

		maxSliceLength := i + 1
		getSliceStart := func(i int, wL int) int { return i + 1 - wL }
		getSliceEnd := func(i int, wL int) int { return i + 1 }

		if found, value := findDigitWord(
			line, i, maxSliceLength, getSliceStart, getSliceEnd,
		); found {
			return value, nil
		}
	}

	return "", errors.New("Failed to find digit")
}

func getLast(
	line string,
) (string, error) {
	lineLength := len(line)
	for i := lineLength - 1; i >= 0; i = i - 1 {
		char := string(line[i])
		if _, err := strconv.Atoi(char); err == nil {
			return char, nil
		}

		maxSliceLength := lineLength - i
		getSliceStart := func(i int, wL int) int { return i }
		getSliceEnd := func(i int, wL int) int { return i + wL }

		if found, value := findDigitWord(
			line, i, maxSliceLength, getSliceStart, getSliceEnd,
		); found {
			return value, nil
		}
	}
	return "", errors.New("Failed to find digit")
}

func findDigitWord(
	line string,
	i int,
	maxSliceLength int,
	getStartSliceIndex func(i int, wordLength int) int,
	getEndSliceIndex func(i int, wordLength int) int,
) (bool, string) {
	// shortest word is 3 letters long, not worth checking until len(line) (- 1) - 3
	for j, xLetterWords := range wordsByLetterCount {
		var wordLength int = j + SMALLEST_WORD_LENGTH
		if wordLength > maxSliceLength {
			break
		}
		sliceStart := getStartSliceIndex(i, wordLength)
		sliceEnd := getEndSliceIndex(i, wordLength)
		ss := line[sliceStart:sliceEnd]
		for _, digitString := range xLetterWords {
			if ss == digitString {
				return true, stringsToDigits[ss]
			}
		}
	}

	return false, ""
}

func main() {
	f, err := os.Open("input.txt")
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var total int

	for scanner.Scan() {
		line := scanner.Text()
		first, err := getFirst(line)
		check(err)
		last, err := getLast(line)
		check(err)

		doubleDigitNumber, err := strconv.Atoi(first + last)
		check(err)

		println("Line:", line)
		println("Number:", doubleDigitNumber)
		total += doubleDigitNumber
		println("Total:", total)
		println()
	}

}
