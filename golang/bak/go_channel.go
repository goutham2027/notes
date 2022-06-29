package main

import (
	"fmt"
)

// TODO: Refactor this

func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch, 0)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func digits(number int, dchn1 chan int, called_by int) {
	fmt.Println("called_by", called_by)
	for number != 0 {
		digit := number % 10
		dchn1 <- digit
		number /= 10
	}
	close(dchn1)
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch, 1)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}
