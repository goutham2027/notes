package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

func main() {
	done := make(chan bool)
	go hello(done)
	// this line is blocking until some Goroutine writes
	// data to the done channel
	<-done // this is legal
	fmt.Println("main fucntion")
}
