package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "Hello World"

	fmt.Println(strings.Contains(sampleString, "lo"))
	fmt.Println(strings.Index(sampleString, "lo"))
	fmt.Println(strings.Count(sampleString, "l"))
	fmt.Println(strings.Replace(sampleString, "l", "x", 3))
}
