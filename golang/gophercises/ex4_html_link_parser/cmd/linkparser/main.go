package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"linkparser"
)

func main() {
	htmlFile := flag.String("htmlFile", "", "path to HTML file")
	flag.Parse()

	f, err := os.Open(*htmlFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	parsedHtml := linkparser.ParseHtml(f)
	fmt.Println(parsedHtml)

	// data, err := io.ReadAll(f)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
