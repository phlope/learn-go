package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	total, err := day1Pt1("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Day 1, Part 1: %d", total)
}

func day1Pt1(filepath string) (int, error) {
	var err error
	file, err := os.Open(filepath)
	if err != nil {
		return 0, fmt.Errorf("unable to open file %v, %v", file, err)
	}

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		var first, last int

		for i := range line {
			extracted, isNum := extractDigitPt1(line[i])
			if isNum {
				first = extracted
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			extracted, isNum := extractDigitPt1(line[i])
			if isNum {
				last = extracted
				break
			}
		}

		sum += first*10 + last //ensure first digit is a factor of ten.

	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("unable to scan file %v, %v", file, err)
	}

	return sum, err

}

func extractDigitPt1(char byte) (int, bool) {
	if char >= '0' && char <= '9' {
		return int(char - '0'), true // Subtracts the ASCII value of '0' (48) from the ASCII value of 'char' (e.g. 5 == 53)
	} else {
		return 0, false
	}
}
