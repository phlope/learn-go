package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// DAY 1
	day1File, err := openFile("input-day-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	total, err := day1Pt1(day1File)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Day 1, Part 1: %d\n", total)

	day1File, err = openFile("input-day-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	total, err = day1Pt2(day1File)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Day 1, Part 2: %d\n", total)

	// DAY 2
	day2File, err := openFile("input-day-2.txt")
	day2Pt1(day2File)
}

func openFile(filepath string) (os.File, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return *file, fmt.Errorf("unable to open file %v, %v", file, err)
	}
	return *file, err
}
