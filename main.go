package main

import (
	"fmt"
	"sync"
)

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
	var memeoryAccess sync.Mutex
	var data int
	go func() {
		memeoryAccess.Lock()
		data += 1
		memeoryAccess.Unlock()
	}()
	memeoryAccess.Lock()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
	memeoryAccess.Unlock()
	//------------------------------------------
}
