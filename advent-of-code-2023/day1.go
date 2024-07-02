package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day1Pt1(file os.File) (int, error) {
	var err error
	sum := 0
	var first, last int

	scanner := bufio.NewScanner(&file)

	for scanner.Scan() {
		line := scanner.Text()

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

func day1Pt2(file os.File) (int, error) {
	var err error
	var first, last int
	sum := 0

	scanner := bufio.NewScanner(&file)

	for scanner.Scan() {
		line := scanner.Text()

		for i := range line {
			extracted, isNum := extractDigitPt2(line[i:]) // add subslice check to loop e.g. 0, 1, 2, 3 | 1, 2, 3 | 2, 3 ...
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
			return i + 1, true // using index instead of parsing string value to int
		}
	}

	return expected, isNum
}
