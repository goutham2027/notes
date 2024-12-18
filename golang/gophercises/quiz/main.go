package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var answer string
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

	ch := make(chan bool)

	go func(ch chan bool) {
		for idx, record := range records {
			fmt.Printf("Problem #%d: %s = ", idx+1, record[0])
			fmt.Scan(&answer)
			if answer == record[1] {
				right += 1
			}
		}
		ch <- timer.Stop()
	}(ch)

	select {
	case <-ch:
		fmt.Println("Exam over before time is up")
	case <-timer.C:
		fmt.Println("Time is up")
	}

	fmt.Printf("Total right answers: %d/%d", right, len(records))
}
