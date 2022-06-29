https://app.pluralsight.com/player?course=go&author=john-sonmez&name=go-m1-overview&clip=0&mode=live

## Chapter-1

### What is Go?
* Compiled
  + Truly compiled down to the processor level.
* Garbage Collected
  + Do not have to explicitly manage memory of programs.
  + memory is managed - allocation and freeing up memory.
  + garbage collection is very fast. Latency free.
* Concurrent
  + Goroutines allows us to handle concurrency by communicating through
    messages or passing data instead of sharing data.

### Why no Inheritance
composition over inheritance.

No method overloading - slows down compiling time.


## Chapter-2: Go Development Environment

### Packages
- Way to modularize code
- Similar to namespace
- Types and functions

### Imports
To use packages

- Source Code
- Local Packages
- Remote packages

### Creating a workspace
```
mkdir go_workspace
cd go_workspace
mkdir src
```

```
echo $GOROOT
echo $GOPATH
```
Set $GOPATH to the `go_workspace` directory.

In packages method names that start with capital letter are exported and that do not are not exported.


## Chapter-3: Variables, Types and Pointers

### Types
- bool
- string
- int, int8/16/32/64
- uint, uint8/16/32/65, uintptr
- byte (uint8)
- rune (int32) like a char
- float32, float64
- complex64, complex128

### Other Types
- Array
- Slice
- Struct
- Pointer
- Function
- Interface
- Map
- Channel

Variable assignment and declaration syntax can be used only with in
functions not in structs.
```
message := "hello world"
```

Type1
```
var message string
message = "hello world"
```

Type2
```
var message = "hello world
```

Type3
```
message := "hello world"
```

### Pointers
pointer is just a type of variable to contain a memory address of
another variable.
integer pointer can only point to integer.

when we pass a variable to a function, we end up making a copy of it.
When we pass a pointer to a function, pointer gets copied but when we
change the value the outer value also changes.

pointer assignment
```
var message string = "Hello go World"
var greeting *string = &message
*greeting = "Hi"

fmt.Println(message, *greeting)
```

### No Classes in Go
Instead of defining classes, Go provides user defined types.

### Struct
```
// exposed publicly because name of the struct starts with capital
letter.
type Salutation struct {
}
```

### Constants
iota - successive integers



