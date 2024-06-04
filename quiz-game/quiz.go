package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var input string
var score int
var answers []string

func main() {
	filepath := flag.String("f", "./problems.csv", "filepath to the questions file")
	flag.Parse()

	questions := fileReader(*filepath)

	for i, el := range questions {
		qa := strings.Fields(strings.Join(el, " "))
		q := qa[0]
		a := qa[1]
		answers = append(answers, a)

		fmt.Println(i+1, q)
		fmt.Scan(&input)
		if input == a {
			score += 1
		}
	}

	fmt.Printf("You scored %v, out of %v", score, len(questions))

	fmt.Println("\nDo you want to see the answers? y/n")
	fmt.Scan(&input)
	if input == "y" {
		for i, ans := range answers {
			fmt.Println(i+1, ans)
		}
	}
}

func fileReader(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Unable to open file %v, %v", file, err)
	}

	reader := csv.NewReader(file)
	questions, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to read file %v, %v", reader, err)
	}

	return questions
}
