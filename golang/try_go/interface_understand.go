package main

import "fmt"

type Animal interface {
	Food() string
}

type Cat struct {
	name string
}

func NewCat(name string) Animal {
	return &Cat{name: name}
}

func (c *Cat) Food() string {
	return fmt.Sprintf("%q drinks a lot of Milk!", c.name)
}

func (c *Cat) Shout() string {
	return fmt.Sprintf("%q shouts meow!!", c.name)
}

func understand_interfaces() {
	cat := NewCat("Purr")
	fmt.Println(cat.Food())
	// fmt.Println(cat.Shout()) // cat.Shout undefined (type Animal has no field or method Shout)

	cat2 := Cat{name: "Binx"}
	fmt.Println(cat2.Food())
	fmt.Println(cat2.Shout())
}
