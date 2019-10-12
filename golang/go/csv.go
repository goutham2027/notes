package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	csvString := "1,2,3,4,5,6"

	fmt.Println(strings.Split(csvString, ","))

	listOfLetters := []string{"c", "a", "b"}

	sort.Strings(listOfLetters)

	fmt.Println("Letters: ", listOfLetters)

	listOfNums := strings.Join([]string{"3", "2", "1"}, ", ")

	fmt.Println("List of Nums: ", listOfNums)

}
