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
	var err error
	fmt.Println("Beginning Program")
	go sayHello() // Extract the function from main thread to background and when the function is finished the function backs to main thread
	fmt.Println("Afrer Sayhello")
	go func() {
		err = doWork()
	}()
	go printNumbers()
	go printLetters()

	time.Sleep(2 * time.Second)
	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println("Work done")
	}

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
