/*
# 1
# Hello World
# Variables
*/
package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {
	var num_i int
	var num_f float32
	var name string

	fmt.Println("Hello World")
	fmt.Println("Default int value", num_i)
	fmt.Println("Default float value", num_f)
	fmt.Println("Default string value", name)
	num_i = 10
	num_f = 20
	fmt.Println(num_i)
	fmt.Println(num_f)

	fmt.Println(reflect.TypeOf(num_i))
	fmt.Println(reflect.TypeOf(num_f))

	fmt.Println(rand.Intn(100))
}
