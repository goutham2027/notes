package main

import "fmt"

	presAge := make(map[string]int)
func main() {

	presAge["TheodreRoosevelt"] = 42

	fmt.Println(presAge["TheodreRoosevelt"])

	presAge["John F. Kennedy"] = 43

	fmt.Println(len(presAge))

	delete(presAge, "John F. Kennedy")

	fmt.Println(len(presAge))
}
