package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

// implementation-1
// func main() {
// 	var answer string
// 	var csvFilePath string
// 	var duration int64
// 	right := 0

// 	fmt.Println("Welcome to the quiz game")

// 	flag.StringVar(&csvFilePath, "csv-file", "problems.csv", "problems csv file path")
// 	flag.Int64Var(&duration, "duration", 3, "Duration of the quiz")
// 	flag.Parse()

// 	timer := time.NewTimer(time.Duration(duration) * time.Second)

// 	f, err := os.Open(csvFilePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	r := csv.NewReader(f)
// 	records, err := r.ReadAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	problems := parseLines(records)

// 	ch := make(chan bool)

// 	go func(ch chan bool) {
// 		for idx, p := range problems {
// 			fmt.Printf("Problem #%d: %s = ", idx+1, p.q)
// 			fmt.Scan(&answer)
// 			if answer == p.a {
// 				right += 1
// 			}
// 		}
// 		ch <- timer.Stop()
// 	}(ch)

// 	select {
// 	case <-ch:
// 		fmt.Println("Exam over before time is up")
// 	case <-timer.C:
// 		fmt.Println("Time is up")
// 	}

// 	fmt.Printf("Total right answers: %d/%d", right, len(problems))
// }

// implementation-2
func main() {
	var csvFilePath string
	var duration int64
	right := 0

	fmt.Println("Welcome to the quiz game")

	flag.StringVar(&csvFilePath, "csv-file", "problems.csv", "problems csv file path")
	flag.Int64Var(&duration, "duration", 3, "Duration of the quiz")
	flag.Parse()

	timer := time.NewTimer(time.Duration(duration) * time.Second)

	f, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	problems := parseLines(records)

	answerCh := make(chan string)

	for idx, p := range problems {
		fmt.Printf("Problem #%d: %s = ", idx+1, p.q)
		go func() {
			var answer string
			fmt.Scan(&answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime is up")
			return
		case answer := <-answerCh:
			if answer == p.a {
				right++
			}
		}
	}
	fmt.Printf("Total right answers: %d/%d \n", right, len(problems))
}

func parseLines(lines [][]string) []problem {
	// When we know the size of the slice, we can declare it at the start
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return problems
}
