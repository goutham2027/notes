https://www.youtube.com/watch?v=YS4e4q9oBaU

### Creators

- Robert Griesemar
- Rob Pike
- Ken Thompson

### Why a new language?

Python - Easy to use, but slow
Java - Fast but increasingly complex type system
C/C++ - Fast but complext type system, slow compile times

- These 3 languages were created when multi-threaded applications were rare.
- Concurrency patterns are patched in

### Go (golang)

- Strong and statically typed
  - Strongly typed: a variable type cannot be changed overtime
  - Static typed: all variables need to be defined at compile time
- Key Features
  - simplicity
  - fast compile times
  - garbage collected
  - built-in concurrency
  - compile to standalone binaries

### Resources

- golang.org
- golang.org/doc
- Effective go in golang.org/doc
- play.golang.org

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```

- Every go application is structured into packages
- main is a special package, entrypoint to the app

### Env variables

#### GOROOT

Where to find go binaries. Use this variable only if go is installed other than the default locations.

#### GOPATH

To put Go apps

```
export GOPATH=<path_to_golib> # all third party code
export PATH=$PATH:$GOPATH/bin
export GOPATH=$GOPATH:<app_code> # workspace code, looks for src code
```

3 directories

```
src
bin
pkg # for intermediate binaries
```

### Editor - VSCode

Extensions

- Go extensions by lukehoban

```go
go run <path>

go build <path_to_package>
go build github.com/goutham2027/firstapp

go install <package> # will instlall in bin directory
go install github.com/goutham2027/firstapp
```

- If you run into the following issue, set the following env `go env -w GO111MODULE=off`

```
go.mod file not found in current directory or any parent directory; see 'go help modules'
```

### Variables

#### Variable declaration

3 ways to declare variables

```go
var <name> <type>
var i int
i = 42
fmt.Println(i)
```

```go
var i int = 42
fmt.Println(i)
```

```go
i := 42
fmt.Println(i)
```

```go
fmt.Printf("%v, %T", i, i)
```

#### Redeclaration and shadowing

#### Visibility

- Lower case variables are seen within the package
- Upper case variables expose outside the package, globally visible
- block scope, when declared in function block

#### Naming conventions

- Pascal or camelCase
- Length of the variable should reflect life of the variable
- Handling acronyms

  - acronyms in uppercase eg: theURL, theHTTP

#### Type conversions

```go
var i int = 42
var j float32
j = float32(i)

// convert int to string
import "strconv"
var i int = 42
var j string
j = strconv.Itoa(i)
```

### Primitives

- Boolean type
- Numeric types
  - Integers
  - Floating point
  - Complext numbers
- Text types

Everytime we initialize a variable it has a zero value
eg: for bool zero value is `false`

```go
var b bool = true
fmt.Printf("%v, %T\n", b, b)
```

```
n := 42 // int: takes system's int
fmt.Printf("%v, %T\n", b, b)

int8
int16
int32
int64
// unsigned integers
uint8
uint16
uint32

// ints of different bytes cannot add
var a int = 10
var b int8 = 3
fmt.Println(a + int(b))

// 32 bit and 64 bit floating point
float32
float64

n := 3.14
n = 13.7e72 // exponential notation
var n float32

// complext number
complex64 (float32 + float32)
complex128 (float64 + float64)
var n complex64 = 1+ 2i

// to get real part
real(n)

// to get imaginary part
imag(n)
```

```go
s := "this is a string"
fmt.Println("%v %T\n", s, s)
fmt.Println("%v %T\n", s[2], s[2]) // returns ascii character and uint8
s[2] = "a" // will result in error, because s[2] has a byte
```

```go
s := "this is a string"
// strings are double quotes
b := []byte(s) // byte slice
fmt.Println("%v, %T\n", b, b)
```

```go
// rune - represents utf32
// single quote
// runes are type alias of uint32
r := 'a' // var r rune = 'a'
fmt.Println("%v, %T\n", r, r) // 97, int32
```

### Constants

- Naming convention
- Types constants
- Untyped constants
- Enumerated constants
- Enumerated expressions

constants need to be assigned in compile time

```go
// name constanats as variable
const myConst

// typed constants
const myConst int = 42

```

iota - a counter

```go
const (
  a = iota
  b = iota
  c = iota
)

func main() {
  fmt.printf("%v\n", a)
  fmt.printf("%v\n", b)
  fmt.printf("%v\n", c)
}

//output
0
1
2
```

```go
const (
  a = iota
  b
  c
)
const (
  a2 = iota
)

func main() {
  fmt.printf("%v\n", a)
  fmt.printf("%v\n", b)
  fmt.printf("%v\n", c)
  fmt.printf("%v\n", a2)
}

//output
0
1
2
0 // iota resets
```

- constants are immutable but can be shadowed
- Replaced by the compiler at compile time. Value must be calculable at compile time.
- Named like variables

  - PascalCase for exported constants
  - camelCase for internal constants

- Enumerated constants

  - Special symbol iota allows related constants to be created easily
  - Iota starts at 0 in each const block and increments by one.
  - Watch out of constant values that match zero values for variables

- Enumerated expressions
  - Operations that can be determined at compile time are allowed
    - Arithmetic
    - Bitwise operations
    - Bitshifting

### Arrays and Slices

#### Array

- by design contiguous in memory
- accessing elements in the array is fast

```go
grades := [3]int
// initializer syntax
grades := [3]int{97, 85, 93}

//
grades := [...]int{97, 85, 93}

var students [3]string

var identityMatrix [3][3]int{[3] int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}

var identityMatrix [3][3]int
identityMatrix[0] = [3]int{1, 0, 0}
identityMatrix[1] = [3]int{0, 1, 0}
identityMatrix[2] = [3]int{0, 0, 1}

b := &a
```

#### Slice

- Slices are backed by an array
- Creation styles

  - Slice existing array or slice
  - Literal style
  - Via make function

  ```
  // create slice with capacity and length = 10
  a := make([]int, 10)

  // slice with length 10 and cap 100
  a := make([]int, 10, 100)
  ```

- Append function to add elements to slice may cause expensive copy operation if underlying array is too small
- Copies refer to same underlying array

```go
a := []int{1, 2, 3}

// make takes two or three args
a := make(
  <type of object we want to create>,
  <length of the slice>,
  <capacity>
)
a := make([]int, 3)
```

```
a := []int{}
fmt.Println(a)
fmt.Printf("Length: %v\n", len(a))
fmt.Printf("Capacity: %v\n", cap(a))
a = append(a, 1)
fmt.Println(a)
fmt.Printf("Length: %v\n", len(a))
fmt.Printf("Capacity: %v\n", cap(a))
a = append(a, 2,3,4,5)
a = append(a, []int{2,3,4,5}...) // spread operator

a := []int{1, 2, 3, 4, 5}
b := a[1:]
fmt.Println(b)
```

```
// array declaration styles
a := [3]int{1, 2, 3}
a := [...]int{1, 2, 3}
var a [3]int

// Everytime we are moving the array around, it's going to copy the underlying data
```

### Maps and Structs

#### Maps

- Cannot use slice as a map's key
- The return order of a map is not gauranteed
- Pass by reference

```
//literal syntax
statePopulations := map[string]int {
  "California": 1123,
}

//using make function
statePopulations := make(map[string]int)
statePopulations["California"] = 123
fmt.Println(statePopulations)

delete(statePopulations, "Georgia")
fmt.Println(statePopulations["Georgia"]) // returns 0

// comma, ok syntax
population, ok := statePopulations["oho"]
// ok is false if value isn't there
```

#### Struct

- dot notation
- anonymous struct
  - used in relatively few cases
  - very shortlived

```
aDoctor := struct{name string}{name: "John"}
```

- naming conventions

  - Capital will
  - Struct name is Capital. Exposed in the package
  - fields are small

- structs are value types
- to reference existing struct

```golang
aDoctor := struct{name string}{name: "John"}
anotherDoctor := &aDoctor
```

- embedding
  - Go uses composition to achieve inheritance called embedding

```golang
type Animal struct {
  Name string
  Origin string
}

type Bird struct {
  Animal
  SpeedKPH float32
  CanFly bool
}

func main() {
  /*
  b := Bird{}
  b.Name = "Emu"
  b.Origin = "Australia"
  b.SpeedKPH = 48
  b.canFly = false
  */
  b := Bird{
    Animal: Animal{Name: "emu", Origin: "Australia"},
    SpeedKPH: 48,
    CanFly: false,
  }
  fmt.Println(b)
}
```

- Tags
  - Provides a string of text
  - Need reflect package to get the tags from the fields

```golang
type Animal struct {
  Name string `required max:"100"`
  Origin string
}

func main() {
  t := reflect.TypeOf(Animal{})
  field, _ := t.FieldByName("Name")
  fmt.Println(field.Tag)

  a := Animal{Name: "emu", Origin: "Australia"}
  fmt.Prinln(a)
}

// we need to use reflect package
```

```golang
type Doctor struct {
  number int
  actorName string
  companions []string
}
func main() {
  aDoctor := Doctor {
    number: 3,
    actorName: "Jon",
    companions: []string {
      "Liz", "Jo",
    }

  }
}

aDoctor.number
```

### If and Switch statements

#### If statement

```golang
if true {
  fmt.Println("True")
}
```

```golang
// the first statement is initializer
if pop, ok := statePopulations["ohio"]; ok {
  fmt.Println(pop)
}
fmt.Println(pop) // pop variable is scoped inside if block
```

```golang
func main() {
  number := 50
  guess := 30
  if guess < number {
    fmt.Println("Too Low")
  }

  if guess > number {
    fmt.Println("Too high")
  }

  if guess == number {
    fmt.Println("You got it")
  }
}

if <condition> {

} else {

}

if <condition> {

} else if <condition> {

} else {

}
```

#### switch

- implicit break in switch between cases
- exit switch early using break keyword

```golang
switch <ch> {
case 1:
  fmt.Println("one")
case 2:
  fmt.Println("two")
default:
  fmt.Println("default")
}

case 1,4, 10: // will test for 1 or 5 or 10

switch i:= 2+3; i {

}


i := 10
switch {
  case i <=10:

  case i >= 20:

}

i := 10
switch {
  case i <=10:
      fallthrough // will execute second case as well

  case i >= 20:

}
```

```golang
// type switch

var i interface{} = 1
switch i.(type) {
  case int:
    fmt.Println("i is an int")
  case float64:
    fmt.Println("i is a float")
  case [2]int:
    fmt.Println("i is an [2]int array")
  case string:
    fmt.Println("i is a string")
}
```

### Looping

- simple loop
  - for initializer; test; incrementer {}
  - for test {}
  - for {}
- exiting early
  - break
  - continue
  - labels
- looping through collections
  - arrays, slices, maps, strings, channels
  - for k, v = range s {}
  - for k = range s {}
  - for \_, v = range s {}

```golang
for i := 0; i < 5; i++ { // i is scoped to for loop
  fmt.Println(i)
}

for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
  fmt.Println(i, j)
}

i := 0 // i is scoped to main function
for ; i< 5; i++ {
  fmt.Println(i)
}

i := 0
for i<5 {
  fmt.Println(i)
  i++
}

// infinite for loop
for {
  break
  continue
}

   Loop:
		for i := 1; i <= 3; i++ {
			for j := 1; j <= 3; j++ {
				fmt.Println(i * j)
				if i*j >= 3 {
					break Loop
				}
			}
		}

  s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

```

### Defer, Panic, Recover

#### Defer

- Invoke a function but delay it's execution to some future point in time
- Executes the function passed to defer after the last statement of the function and before returns
- Defer executes in LIFO order
- For resources in loop, handle them individually instead of using defer. Because until you close the function all the resources opened in the loop will be opened and can potentially cause memory issues.
- Args evaluated at time defer is executed, not at the time of the function call.

```
a := "start"
defer fmt.Println(a)
a = "end"

// outputs: start
// takes the value at the time defer is called
```

#### Panic

- In Golang we do not have exceptions
- panic("<message>")
- panics happen after defer statements are executed

```
fmt.Println("start")
defer fmt.Println('this was deferred')
paic("something bad happened")
fmt.Println("end")
```

#### Recover

- use to recover from panics
- recover() returns nil if application is not panicing
- returns errors if application is panicing
- only useful in deferred functions

### Pointers

- Creating pointers

```
var a int = 42
var b *int = &a

```

- `b` is pointing to memory location of `a`
- no pointer arithmetic in golang. use `unsafe` package to allow pointer arithmetic.

- Dereferencing pointers
  - Dereference a pointer by preceding with an asterisk (\*)
  - Complext types (eg: structs) are automatically dereferenced
- The new function
  ```golang
  ms := myStruct{foo:42}
  p := &ms
  ```
- nil
- Types with internal pointers
  - all assignment operations in Go are copy operations
  - slices and maps contain internal pointers, so copies point to same underlying data
- Slice is references where as array is copy by value. Slice points to the first element of the array.

### Functions

- Basic syntax

```golang
package main

func main() {

}

```

- Parameters

```go
func main() {
  sayMessage("Hello Go!")
}

func sayMessage(msg string) {
  fmt.Println(msg)
}
```

```golang
// passing by reference
func main() {
  greeting := "Hello"
  name := "Stacey"
  sayGreeting(&greeting, &name)
  fmt.Println(name)
}

func sayGreeting(greeting, name *string) {
  fmt.Println(*greeting, *name)
  *name = "Ted"
  fmt.Println(*name)
}

// op
Hello Stacey
Ted
Ted
```

```golang
// variadic parameters
func sum(values ...int) int {
  result := 0
  for _, v := range values {
    result += v
  }
  return result
}
```

- Return values

```golang
// returning a pointer
func sum(values ...int) *int {
  return &result
}
s := sum(1, 2, 3, 4, 5)
fmt.Println("The sum is ", *s)

// named return value
func sum(values ...int) (result int) {
return
}

// multiple return values
// idiomatic go
func divide(a, b float64) (float64, error) {
  if b == 0.0 {
    return 0.0, fmt.Errorf("Cannot dividde by zero")
  }
  return a/b, nil
}
```

- Anonymous functions

```golang
func main() {
  func() {
    fmt.Println("Hello Go!")
  }() // anonymous function. paranthesis are to invoke the function
}

// to avoid reading from outer scope
// this works fine when we are running asynchronously
func main() {
  for i:=0; i<5; i++ {
    func(i int) {
      fmt.Println(i)
    }(i)
  }
}

func main() {
  f := func() {
    fmt.Println("Hello Go")
  }
  // var f func() = func() {}
  f()
}
```

- Functions as types
  we can pass functions like any other type
  ```golang
  var divide func(float64, float64) (float64, error)
  divide = func(a, b float64) (float64, error) {
    if b == 0.0 {
      return 0.0, fmt.Errorf("Cannot divide by zero")
    }
    d, err := divide(6.0, 0.0)
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Println(d)
  }
  ```
- Methods

```
func main() {
  g := greeter {
    greeting: "Hello",
    name: "go",
  }
  g.greet()
}

type greeter struct {
  greeting string
  name string
}
// (g greeting) makes a function a method
// we get copy of the struct not the pointer to the original struct
func (g greeting) greet() {
  fmt.Println(g.greeting, g.name)
}

// we can use a pointer receiver
func (g *greeting) greet() {
  fmt.Println(g.greeting, g.name)
}
```

### Interfaces
