package main

import (
	"bufio"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("input.txt")
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var total int

	for scanner.Scan() {
		line := scanner.Text()
		print(line, " -- ")

		var first string
		var last string

		for i := range line {
			char := string(line[i])
			if _, err := strconv.Atoi(char); err == nil {
				first = char
				break
			}
		}
		for i := len(line) - 1; i >= 0; i = i - 1 {
			char := string(line[i])
			if _, err := strconv.Atoi(char); err == nil {
				last = char
				break
			}
		}
		doubleDigitNumber, err := strconv.Atoi(first + last)
		check(err)

		print("Number: ", doubleDigitNumber, " -- ")
		total += doubleDigitNumber
		print("Total: ", total)
		println()
	}

}
