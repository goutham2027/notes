- A weak pointer is a reference to an object that does not increase the reference count of that object.
- Unlike a strong reference, a weak pointer does not stop the garbage collector from reclaiming the referenced object
  if no strong references exist.
- In Go 1.24 weak pointers will be part of the new `weak` package.

https://goplay.tools/snippet/hHJ_Unpb3fo

```golang
package main

import (
	"fmt"
	"runtime"
	"weak"
)

func main() {
	name := new(string)
	*name = "example"

	weakName := weak.Make(name)

	fmt.Println(*weakName.Value())

	runtime.GC()
	fmt.Println(weakName.Value())
}

```
