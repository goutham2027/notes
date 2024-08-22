https://medium.com/hprog99/mastering-generics-in-go-a-comprehensive-guide-4d05ec4b12b

## Generics

- Introduced in Go 1.18
- to write reusable & type safe code
- with generics (parametric polymophism), we can define functions and data structures
  that work with multiple types without the need for runtime type assertions or type casting
- Generics enables us to write code that works on multiple types without explicitly specifying
  the types upfront.
- Prior to generics, we had to use interfaces and type assertions

### Type Parameters & Type constraints

- building blocks
- type parameters: placeholders for the actual types that will be used when the
  generic function or data structure is instantiated
- type constraints: restrict the types that can be used as type arguments

```golang
package main

import "fmt"

func PrintSlice[T any](s []T) {
  for _, v := range s {
    fmt.Println(v)
  }
}

func main() {
  intSlice := []int{1, 2, 3}
  stringSlice := []string{"hello", "world"}

  PrintSlice[int](intSlice)
  PrintSlice[str](intSlice)
}
```
