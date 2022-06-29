package main

import "fmt"

/*
	recover continues execution even if there is any error
*/
func main() {
	/*
	 *fmt.Println(safeDiv(3, 0))
	 *fmt.Println(safeDiv(3, 2))
	 */

	demPanic()
}

func safeDiv(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()

	solution := num1 / num2

	return solution
}

func demPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("PANIC")
}
