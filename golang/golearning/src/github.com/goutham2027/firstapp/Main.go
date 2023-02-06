package main

import (
	"sync"
)

var wg = sync.WaitGroup{}

/* Example-1
func main() {
	ch := make(chan int)
	wg.Add(2)
	go func() {
		// receiving the data from the channel
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	go func() {
		// writing to channel
		ch <- 42
		wg.Done()
	}()

	wg.Wait()
}
End of example-1 */

/* Example - 2
func main() {
	ch := make(chan int)
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() {
			ch <- 42
			wg.Done()
		}()
	}
	wg.Wait()
}

End of Example-2 */

/* Example-3
func main() {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()
}

End of example-3 */

// example-4
// Buffered channels
// func main() {
// 	ch := make(chan int, 50)
// 	wg.Add(2)
// 	go func(ch <-chan int) {
// 		for i := range ch {
// 			fmt.Println(i)
// 		}
// 		wg.Done()
// 	}(ch)
// 	go func(ch chan<- int) {
// 		ch <- 42
// 		ch <- 27
// 		close(ch)
// 		wg.Done()
// 	}(ch)
// 	wg.Wait()
// }

// Example - 5
// ok syntax channels
// func main() {
// 	ch := make(chan int, 50)
// 	wg.Add(2)
// 	go func(ch <-chan int) {
// 		for {
// 			if i, ok := <-ch; ok {
// 				fmt.Println(i)
// 			} else {
// 				break
// 			}
// 		}
// 		wg.Done()
// 	}(ch)

// 	go func(ch chan<- int) {
// 		ch <- 42
// 		ch <- 27
// 		close(ch)
// 		wg.Done()
// 	}(ch)

// 	wg.Wait()
// }

// Example - 6
// select statements
// signal only channel
// var doneCh = make(chan struct{}) // doesn't take any memory
// while using
// doneCh <- struct{}{} // struct with empty fields and initializing
