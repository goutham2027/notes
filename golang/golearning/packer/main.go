package main

import (
	"fmt"
	str "strings" // Package Alias

	"github.com/goutham2027/packer/numbers"
	"github.com/goutham2027/packer/strings"
	"github.com/goutham2027/packer/strings/greeting"
)

func main() {
	fmt.Println(numbers.IsPrime(19))
	fmt.Println(greeting.WelcomeText)
	fmt.Println(strings.Reverse("Goutham"))
	fmt.Println(str.Count("Go is Awesome. I love Go", "Go"))
}
