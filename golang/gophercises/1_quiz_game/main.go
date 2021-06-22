package main

import (
  "encoding/csv"
  "fmt"
  "os"
  "log"
  "io"
  "strconv"
  //"bufio"
)

func main() {
  // fileName defaults to problems.csv but should allow user to input
  fileName := "problems.csv"
  //var questionsAnswers map[string]string
  csvfile, err := os.Open(fileName)
  numOfRightAnswers := 0
  var captureAnswer int64

  //questionsAnswers = make(map[string]string)

  if err != nil {
    log.Fatalln(err)
  }

  r := csv.NewReader(csvfile)

  for {
    fmt.Println("Capture answer")
    fmt.Println(captureAnswer)
    // Read each record from csv
    record, err := r.Read()
    if err == io.EOF {
      break
    }

    if err != nil {
      log.Fatal(err)
    }

    question := record[0]
    answer, _ := strconv.ParseInt(record[1], 10, 64)

    fmt.Println(question)
    fmt.Print("Enter the answer: ")
    n, _ := fmt.Scanf("%d\n", &captureAnswer)

    if n == 0 {
      fmt.Println("you didn't enter")
    }

    if captureAnswer == answer {
      numOfRightAnswers += 1
    }
    captureAnswer = 0

    //fmt.Printf("Question: %s Answer %s\n", question, answer)
    //questionsAnswers[record[0]] = record[1]
    //fmt.Println(questionsAnswers)

  }

  fmt.Printf("Number of right answers %d", numOfRightAnswers)

  //records, _ := r.ReadAll()
  //fmt.Print(records)
}
