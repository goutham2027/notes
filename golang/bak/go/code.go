/* Go programming in an hours */
// https://www.youtube.com/watch?v=CF9S4QZuV30
package main

import "fmt"

// comments

/*
	multi line
	comments
*/

func main() {
	/*
	 *  fmt.Println("Hello World")
	 *
	 *  var age int = 40
	 *
	 *  var favnum float64 = 1.61334
	 *
	 *  randNum := 1
	 *
	 *  fmt.Println(randNum)
	 *
	 *  fmt.Println(age, favnum)
	 */

	/*
	 *const pi float64 = 3.14159
	 */
	/*
	 *
	 * var (
	 *   varA = 2
	 *   varB = 3
	 * )
	 */

	/*
	 * var myName string = "Goutham"
	 *
	 * fmt.Println(myName + " " + " Pilla")
	 */

	/*
	 *var isOver40 bool = true
	 */

	/*
	 *fmt.Printf("%f \n", pi)
	 */
	// Data type
	/*
	 *fmt.Printf("%T \n", isOver40)
	 */
	/* to print boolean */
	/*
	 *fmt.Printf("%T \n", isOver40)
	 */

	// logical operators
	// &&, ||, ! (not)

	/*
	 *  i := 1
	 *
	 *  for i <= 10 {
	 *    fmt.Println(i)
	 *    i++
	 *  }
	 */

	// Relational operators
	// == != <= >=
	/*
		  *for j := 0; j < 5; j++ {
			*  fmt.Println(j)
		  *}
	*/

	/*
	 *yourAge := 18
	 */

	/*
	 *if yourAge >=16 {
	 *  fmt.Println("You can drive")
	 *} else if yourAge >=18 {
	 *  fmt.Println("You can Vote")
	 *} else  {
	 *  fmt.Println("have fun")
	 *}
	 */

	/*
	*switch yourAge {
	*case 16: fmt.Println("Go drive")
	*case 18: fmt.Println("Go vote")
	*default: fmt.Println("Have fun")
	*}
	 */

	/*
	 *  var favNums2[5] float64
	 *
	 *  favNums2[0] = 12.1
	 *  favNums2[1] = 2.1
	 *  favNums2[2] = 1.1
	 *  favNums2[3] = 12.0
	 *  favNums2[4] = 12.112
	 *
	 *  fmt.Printf("%f", favNums2[3])
	 */

	/*
	 * favNums3 := [5] float64 {1, 2, 3, 4, 5}
	 *
	 * for i, value := range favNums3 {
	 *   fmt.Println(value, i)
	 * }
	 */

	// slices
	/*
	 * numSlice := []int {5, 4, 3, 2, 1}
	 * numSlice2 := numSlice[3:5]
	 * // remaining slices are just like python
	 *
	 * fmt.Println("numSlice2[0] =", numSlice2[0])
	 */

	numSlice := []int{5, 4, 3, 2, 1}

	/* creates array of size 10, with first values as zero */
	numSlice3 := make([]int, 5, 10)

	copy(numSlice3, numSlice)

	fmt.Println(numSlice3[0])

	numSlice3 = append(numSlice3, 0, -1)

	fmt.Println(numSlice3[6])
}
