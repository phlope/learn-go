package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	total, err := day1Pt1("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Day 1, Part 1: %d", total)

	total, err = day1Pt2("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nDay 1, Part 2: %d", total)
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

/* Duplicating the day 1 function to keep a history of steps taken for pt1 and 2
   To clean up as this progesses
*/

func day1Pt2(filepath string) (int, error) {
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
			extracted, isNum := extractDigitPt2(line[i:])
			if isNum {
				first = extracted
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			extracted, isNum := extractDigitPt2(line[i:])
			if isNum {
				last = extracted
				break
			}
		}

		sum += first*10 + last

	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("unable to scan file %v, %v", file, err)
	}

	return sum, err

}

func extractDigitPt2(line string) (int, bool) {
	var expected int
	var isNum bool

	numbers := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	expected, isNum = extractDigitPt1(line[0])
	if isNum {
		return expected, true
	}

	for i, num := range numbers {
		if strings.HasPrefix(line, num) {
			return i + 1, true
		}
	}

	return expected, isNum
}
