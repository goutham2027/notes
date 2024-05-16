package main

import (
	"flag"
	"fmt"
)

func GetVersion() {

	version := flag.Bool("version", false, "get version")
	fmt.Println(*version)
}
