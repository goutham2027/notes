package main

import "fmt"

// type Salutation string
type Salutation struct {
	name string
	greeting string
}

const (
	PI = 3.14
	Language = "Go"
)

const (
	A = iota
	B
	C
)

func main() {
	// var s = Salutation{"Foo", "Hello"}
	// var s = Salutation{name: "Foo", greeting: "Hello"}
	// var s = Salutation{}
	// s.name = "Foo"
	// s.greeting = "Hello"
	// var message Salutation = "Hello World"
	// fmt.Println(s.name)
	// fmt.Println(s.greeting)

	fmt.Println(PI)
	fmt.Println(Language)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}