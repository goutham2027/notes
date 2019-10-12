package main

import "fmt"

// uni directional channels - channles that only send or receive data
func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	// we created a bi-directional channel but the function can
	// only write
	sendch := make(chan int)
	go sendData(sendch)
	fmt.Println(<-sendch)
}

// senders have the ability to close the channel to notify receivers that no
// more data will be sent on the channel.

// receivers can use an additional variable while receiving data from the
// channel to check whether the channel has been closed.
// v, ok := <- ch
// ok is true if the value was received by a successful send operation to a channel
// if ok is false it means that we are reading from a closed channel.
