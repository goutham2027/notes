package main

import "fmt"

// Reference: https://golangbot.com/channels/
// data := <- a // read from channel a
// a <- data // write to channel a
/*
Sends and receives to a channel are blocking by default.
When a data is sent to a channel, the control is blocked in the
send statement until some other Goroutine reads from that channel.

Similarly when data is read from a channel, the read is blocked until some
Goroutine writes data to that channel.
*/

func main() {
	var a chan int
	fmt.Println("Type of a is %T", a)
	if a == nil {
		fmt.Println("Channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}
}
