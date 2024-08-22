package main

import "fmt"

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}

	lastIndex := len(s.data) - 1
	value := s.data[lastIndex]
	s.data = s.data[:lastIndex]
	return value, true
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func main() {
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push('a') // adds ascii value
	intStack.Push(2)

	fmt.Println(intStack.data)
	fmt.Println(intStack.Pop())
	fmt.Println(intStack.Pop())
	fmt.Println(intStack.Pop())
	fmt.Println(intStack.Pop())

	stringStack := Stack[string]{}
	stringStack.Push("hello")
	stringStack.Push("world")

	fmt.Println(stringStack.Pop())  // Output: world, true
	fmt.Println(stringStack.Size()) // Output: 1
	fmt.Println(stringStack.Pop())
	fmt.Println(stringStack.Pop())
}
