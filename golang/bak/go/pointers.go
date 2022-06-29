package main

import "fmt"

func main() {
	x := 0

	changeXVal(&x)

	fmt.Println("x =", x)
	fmt.Println("Memory address for x =", &x)
}

func changeXVal(x *int) {
	*x = 2
}
