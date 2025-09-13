package main

import "fmt"

func speak(arg string, ch chan string) {
	ch <- arg
}
func main() {
	// 1. Concurrency Pattern 1 Race Condition

	// var data int
	// go func() {
	// 	data++
	// }()
	// time.Sleep(1 * time.Second)
	// if data == 0 {
	// 	fmt.Printf("the value is %v.\n", data)
	// }

	//----------------------------------------

	//2. Memory access Synchronization
	// var memeoryAccess sync.Mutex
	// var data int
	// go func() {
	// 	memeoryAccess.Lock()
	// 	data += 1
	// 	memeoryAccess.Unlock()
	// }()
	// memeoryAccess.Lock()
	// if data == 0 {
	// 	fmt.Printf("the value is %v.\n", data)
	// } else {
	// 	fmt.Printf("the value is %v.\n", data)
	// }
	// memeoryAccess.Unlock()
	//-----------------------------------------

	// 3. Channel

	// var ch chan T
	// chan means channel and T is the tyope of data we are going to send
	// var ch chan string
	// ch2 := make(chan string)
	// fmt.Println(ch, ch2)

	// Declaring and passing and receving data through channel(Unbuffered)
	// ch := make(chan string)
	// go speak("Hello World", ch)

	// data := <-ch
	// fmt.Println(data)

	// Buffered Channel

	// ch := make(chan int, 3)

	// fmt.Println("Buffer initiation")
	// ch <- 1
	// ch <- 2
	// ch <- 3
	// fmt.Println("Buffer receievs starts")

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// Directional Channel
	ch := make(chan bool)

	go sayTrue(true, ch)

	// ch2 := make(chan bool)

	// ch2 <- <-ch
	fmt.Println(<-ch)

}

func sayTrue(word bool, ch chan<- bool) {
	ch <- word
}
