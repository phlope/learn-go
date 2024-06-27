package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var input string
var score int

type problem struct {
	question string
	answer   string
}

func main() {
	filepath := flag.String("f", "./problems.csv", "Path to the questions csv file")
	duration := flag.Int("t", 60, "Quiz timer total seconds")
	flag.Parse()

	rows, err := fileReader(*filepath)
	if err != nil {
		log.Fatal(err)
	}

	questions, err := parseRows(rows)
	if err != nil {
		log.Fatal(err)
	}

	timer := time.NewTimer(time.Duration(*duration) * time.Second)

	for i, problem := range questions {
		fmt.Println(i+1, problem.question)

		ansCh := make(chan string)
		go func() {
			fmt.Scan(&input)
			ansCh <- input
		}()

		select {
		case <-timer.C:
			break
		case input := <-ansCh:
			if input == problem.answer {
				score += 1
			}
			if i == len(questions)-1 {
				break
			}
		}
	}
	fmt.Printf("You scored %v out of %v", score, len(questions))

}

func parseRows(rows [][]string) ([]problem, error) {
	var err error

	problems := make([]problem, len(rows))
	for i, row := range rows {
		if len(row) != 2 {
			return problems, fmt.Errorf("unexpected element content, expected question and answer, got: %v", row)
		}
		problems[i] = problem{
			question: row[0],
			answer:   row[1],
		}

	}
	return problems, err
}

func fileReader(filepath string) ([][]string, error) {
	var rows [][]string
	var err error

	file, err := os.Open(filepath)
	if err != nil {
		return rows, fmt.Errorf("unable to open file from %s, %v", filepath, err)

	}

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	rows, err = reader.ReadAll()
	if err != nil {
		return rows, fmt.Errorf("unable to read csv from %v, %v", file, err)
	}

	return rows, err
}
