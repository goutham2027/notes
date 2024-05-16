package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Printf("do something: mykey's value is %s\n", ctx.Value("myey"))
}

// func main() {
// 	ctx := context.Background()
// 	ctx = context.WithValue(ctx, "myKey", "myValue")
// 	doSomething(ctx)
// }
