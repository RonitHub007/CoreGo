package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	   Goroutine scheduling in Go
	   . Simplify Concurrent programming
	   . Efficient handle parallel tasks such as I/O operations , calculations and more
	   . Provides a way to perform tasks concurrently without manually managing threads
	*/

	/*
		Basics of goroutines

			1. Creating Goroutines(Use the `go` keyword to start a new goroutine)
			2. Goroutine Lifecycle
			3. Goroutine Scheduling
	*/

	/*
		Goroutines Scheduling in Go
		. Managed by the Go Runtime SCHEDULER
		. Uses M:N Scheduling Model
		. Efficent Multiplexing
	*/

	/*
		Common Pitfalls and Best Practices
		1. Avoid Gorutine Leaks
		2. Limiting Goroutine Creation
		3. Proper Error Handling
		4. Synchronization
	*/

	/* Goroutines are just functions that leave the main thread and run in the background and come back to join the main thread once the functions are finished/ready to return any values */
	/* GoRoutines do not stop the program flow and are non blocking */
	// var err error
	// fmt.Println("Beginning Program")
	// go sayHello() // Extract the function from main thread to background and when the function is finished the function backs to main thread
	// fmt.Println("Afrer Sayhello")
	// go func() {
	// 	err = doWork()
	// }()
	// go printNumbers()
	// go printLetters()

	// time.Sleep(2 * time.Second)
	// if err != nil {
	// 	fmt.Println("Error")
	// } else {
	// 	fmt.Println("Work done")
	// }

	// Channels()
	// bufferedChannel()
	// bufferedChannelAdv()
	// channelSync()
	channelSyncnew()

}

/*
Goroutine lifecycle
M => number of goroutines mapped to N = number of OS Threads
Goroutines helps achieve concurrency and go scheduler helps achieve parallalism
*/
func sayHello() {
	fmt.Println("Hello From GoRoutine")
	time.Sleep(1 * time.Second)
}
func printNumbers() {
	for i := 0; i < 5; i += 1 {
		fmt.Println(i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}
func printLetters() {
	for _, v := range "Ronit" {
		fmt.Println(string(v), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

// Error Propagation in Golang , error needs to be back at main thread
func doWork() error {
	time.Sleep(1 * time.Second)
	return fmt.Errorf("an error in doWork")
}

/*
Channels
	1. Why Use Channels ?
		=> Enable safe and effcient communication between concurrent goRoutines
		=> Help synchronize and manage the flow of data in concurrent programs
    2. basics of Channels
		make(chan type)
	3. Create goRoutine to e channels
	4. Channels in Go is blocking but GoRoutines are non-blocking
*/

func Channels() {
	// variable := make(chan type)
	greeting := make(chan string)

	greetString := "hello Ron"

	// Receive :=> channelName <- value
	go func() {
		greeting <- greetString
		greeting <- "World"
		for _, j := range "Ronit" {
			greeting <- string(j)
		}
	}()
	go func() {
		receiver := <-greeting // Receving in channels is non-blocking , it is communicating b/w different go routtins, like main and Channels
		receiver = <-greeting

		for range len("Ronit") {
			fmt.Println(<-greeting)
		}
		fmt.Println(receiver)
	}()
	time.Sleep(1 * time.Millisecond)
}

/*
Buffered Channel : Means channel with storage
allows channels to hold a limited number of values

buffered vs Unbuffered is bufferd channels allows asynchronous, helps load balancing
*/
/*
Why Use Buffered Channels ?
	1. Asynchronous Communiction
	2. Load Balancing
	3. Flow Control
Creating Buffered Channels
	1. make(chan Type, capacity)
	2. Buffer capacity
Key concept of channel buffereing
	1. Blocking behaviour
	2. Non-Blocking Operations
	3. Impact on performance
Best Practices for Using Bufffered Channels
	2. Avoid Over buffering
	Graceful shutdown
	Monitroing Buffer channing


*/
func bufferedChannel() {
	// make(chan Type, capacity)

	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Received: ", <-ch)
	}()
	ch <- 3
	fmt.Println("Value: ", <-ch)
	fmt.Println("Value: ", <-ch)
	fmt.Println("Buffered Channels")

}

func bufferedChannelAdv() {
	// Blocking on receive only if the buffer is empty
	ch := make(chan int, 2)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		ch <- 2
	}()
	fmt.Println("Value: ", <-ch)
}

/*
Channel Synchronization
 1. Ensures that the data is properly exchanged between Goroutines
 2. Coordinates the execution flow to avoid race conditions and ensures predictable behaviour
 3. Helps manage the lifecycle of goroutines and the completion of tasks
*/
func channelSync() {
	done := make(chan struct{})

	go func() {
		fmt.Println("Working...")
		time.Sleep(2 * time.Second)
		done <- struct{}{}
	}()

	<-done
	fmt.Println("Finished.....")
}

func channelSyncnew() {
	ch := make(chan int)

	go func() {
		fmt.Println("proces...")
		ch <- 9
		time.Sleep(1 * time.Second)
		fmt.Println("sENT VALUE")
	}()
	value := <-ch
	fmt.Println(value)
}

// func
