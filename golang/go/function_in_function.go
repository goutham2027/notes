package main

import "fmt"

// closures
func main() {

	/*
	 *  num3 := 3
	 *
	 *  doubleNum := func() int {
	 *    num3 *= 2
	 *    return num3
	 *  }
	 *
	 *  fmt.Println(doubleNum())
	 *  fmt.Println(doubleNum())
	 *
	 */

	fmt.Println(factorial(3))

}

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}
