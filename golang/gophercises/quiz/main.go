package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	var answer string
	right := 0

	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		fmt.Printf("%s \n", record[0])
		fmt.Scanln(&answer)
		if answer == record[1] {
			right += 1
		}
	}
	fmt.Printf("Total right answers: %d/%d", right, len(records))
}
