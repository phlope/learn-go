package main

import (
	"bufio"
	"fmt"
	"os"
)

type turn struct {
	colour string
	value  int
}

func day2Pt1(file os.File) {
	// For each turn in a game check if presented value for colour is <= held value for colour
	// If true, pass to next, if false, mark round as impossible and move to the next round
	// If all true, record game id and add to running sum

	// var bagContent []turn

	scanner := bufio.NewScanner(&file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

	}
}
