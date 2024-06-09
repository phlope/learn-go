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

	rows := fileReader(*filepath)
	questions := parseRows(rows)

	timer := time.NewTimer(time.Duration(*duration) * time.Second)

	for i, problem := range questions {
		fmt.Println(i+1, problem.question)

		//Using concurrency guidance from solution here
		//Moving scan to a new go routine sending the response back via new channel ansCh
		//Now each case reacts on a response from either channel and the timer end is no longer blocked on user response
		ansCh := make(chan string)
		go func() {
			fmt.Scan(&input)
			ansCh <- input
		}()

		select {
		case <-timer.C:
			fmt.Printf("You scored %v out of %v", score, len(questions))
			return
		case input := <-ansCh:
			if input == problem.answer {
				score += 1
			}
		}

	}

}

func parseRows(rows [][]string) []problem {
	problems := make([]problem, len(rows))
	for i, row := range rows {
		problems[i] = problem{
			question: row[0],
			answer:   row[1],
		}
	}
	return problems
}

func fileReader(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Unable to open file %v, %v", file, err)
	}

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to read file %v, %v", reader, err)
	}

	return rows
}
