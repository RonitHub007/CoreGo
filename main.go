// package main

// import (
// 	"fmt"
// )

// func produceNumbers(mainCh chan int) {
// 	for i := 0; i < 20; i++ {
// 		mainCh <- i
// 	}
// 	close(mainCh)
// }

// func squareNumbers(workerCh chan int, mainCh chan int) {
// 	// length := len(mainCh)

// 	for i := 0; i < 20; i++ {
// 		data := <-mainCh
// 		workerCh <- data * data
// 	}
// 	close(workerCh)
// }
// func main() {
// 	ch := make(chan int)
// 	ch2 := make(chan int)
// 	go produceNumbers(ch)

// 	go squareNumbers(ch2, ch)

// 	// time.Sleep(20 * time.Second)
// 	for i := 0; i < 20; i++ {
// 		fmt.Println(<-ch2)
// 	}
// }

package main

import (
	"fmt"
	"time"
)

func main() {

	one := make(chan string)
	two := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		two <- "Two"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		one <- "One"
	}()

	select {
	case result := <-one:
		fmt.Println("Recived: ", result)
	case result := <-two:
		fmt.Println("Recived: ", result)
	}
	close(one)
	close(two)
}
