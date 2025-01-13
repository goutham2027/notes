package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := flag.String("file", "stories.json", "json file with story")
	flag.Parse()
	fmt.Printf("Using the story in %s. \n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	arcs, err := cyoa.JsonStory(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", arcs)
	// fmt.Println(arcs["debate"].ArcSuggestions[0].ArcName)
}
