package main

import "fmt"

func main() {
	/*
	 *  profile := make(map[string]string)
	 *  profile["name"] = "Goutham"
	 *  profile["age"] = "10"
	 *  printMap(profile)
	 *
	 */
	// Map with array of ints
	scores := make(map[string][]string)
	scores["Goutham"] = append(scores["Goutham"], "20")
	scores["Goutham"] = append(scores["Goutham"], "30")

	printMap(scores)

}

func printMap(m map[string][]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}
