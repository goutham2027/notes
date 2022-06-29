package main

import "fmt"

func main() {
	listNums := []float64{1, 2, 3, 4, 5}

	fmt.Println("Sum :", addThemUp(listNums))

	num1, num2 := next2Values(5)
	fmt.Println(num1, num2)
}

func addThemUp(numbers []float64) float64 {
	sum := 0.0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

func next2Values(number int) (int, int) {
	return number + 1, number + 2
}
