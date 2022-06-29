package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://beautifulcode.in")
	fmt.Println(resp)
}
