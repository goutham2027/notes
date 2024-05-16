package main

// import (
// 	"fmt"
// 	"sync"
// )

// var mu sync.Mutex

// func main() {
// 	wg := sync.WaitGroup{}
// 	var var1 []string
// 	num := 30
// 	wg.Add(num)
// 	for i := 0; i < num; i++ {
// 		go fun1(&wg, &var1, i)
// 	}

// 	wg.Wait()
// 	// fun1(&var1)
// 	fmt.Println(var1)
// 	fmt.Println(len(var1) == num)
// }

// func fun1(wg *sync.WaitGroup, var1 *[]string, i int) {
// 	// func fun1(var1 *[]string) {

// 	mu.Lock()
// 	defer wg.Done()
// 	defer mu.Unlock()
// 	msg := fmt.Sprintf("From fun %d", i)
// 	*var1 = append(*var1, msg)
// 	// fmt.Println(*var1)
// }
