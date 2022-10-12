package main

import (
	"fmt"
	"encoding/json"
)
// source
// https://go.dev/blog/json

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	fmt.Println("Hello world")
	m:= Message{"Alice", "Hello", 1234993499}
	// var n Message

	// converting to json
	fmt.Println("converting struct to json")
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Printf("%s\n", b)

	b = []byte(`{"Name": "Bob", "Food": "Pickle", "Something": "lafksjd"}`)
	// var objs map[string]*json.RawMessage
	var objs map[string]interface{}

	// json to struct
	err = json.Unmarshal(b, &objs)
	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	fmt.Println("Reading json to a struct")

	fmt.Printf("%v", objs)

}
