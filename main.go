// package main

// import (
// 	"fmt"
// 	"sync"
// )

// // func work() {

// // 	fmt.Println("Working....")
// // }

// //#First way to call
// // func main() {
// // 	// As we have mentioned there are three functions a waitgroup have , lets implement and
// // 	// See how and what they are capable of doing

// // 	var wg sync.WaitGroup

// // 	wg.Add(1)
// // 	go func() {
// // 		defer wg.Done()
// // 		work()
// // 	}()
// // 	wg.Wait()

// // }

// // #2nd way of doing it
// /*
// func work(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Working.....")
// }
// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	go work(&wg)
// 	wg.Wait()
// }
// */

// // Mutex

// // We will make a counter to update our value

// type Counter struct {
// 	value int
// }

// func (c *Counter) updateCounter(n int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Adding %d to %d\n", n, c.value)
// 	c.value += n
// }
// func main() {
// 	var count Counter
// 	var wg sync.WaitGroup

// 	wg.Add(3)

// 	go count.updateCounter(24, &wg)
// 	go count.updateCounter(24, &wg)
// 	go count.updateCounter(24, &wg)

// 	wg.Wait()
// 	fmt.Println(count.value)
// }

// // In the above example c.Value stays 0 alwas evemn though we are the final output as 72
//

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type Counter struct {
// 	mu    sync.RWMutex
// 	value int
// }

// func (c *Counter) UpdateCounter(value int, Wg *sync.WaitGroup) {
// 	defer Wg.Done()
// 	c.mu.Lock()
// 	fmt.Printf("Adding %d to %d\n", value, c.value)
// 	c.value += value
// 	c.mu.Unlock()
// }

// func (c *Counter) GetValue(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	c.mu.RLock()
// 	fmt.Println("Get value:", c.value)
// 	c.mu.RUnlock()
// 	// time.Sleep(400 * time.Millisecond)
// }
// func main() {

// 	var Wg sync.WaitGroup
// 	c := Counter{}
// 	// First run writers
// 	Wg.Add(4)
// 	go c.UpdateCounter(10, &Wg)
// 	go c.UpdateCounter(20, &Wg)
// 	go c.GetValue(&Wg)
// 	go c.GetValue(&Wg)
// 	Wg.Wait()

// 	Wg.Wait()
// }

/*
sync.Cond => It allows one or more goroutines to wait until another goroutine signals them to continue.
Useful when goroutines are depepndent on some shared state change
*/

/*
Pool : It relates with memory .

sync.Pool  will make it memory effecient
.Get() : Get resource
.Put() : Put resource
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

type SomeObject struct {
	Data []byte
}

func createObject() *SomeObject {
	return &SomeObject{
		Data: make([]byte, 1024*1024),
	}
}

func main() {
	var objects []*SomeObject

	objectPool := sync.Pool{
		New: func() interface{} {
			return createObject()
		},
	}
	for i := 0; i < 1000; i++ {
		obj := objectPool.Get().(*SomeObject)
		// obj := createObject()
		objects = append(objects, obj)
		objectPool.Put(obj)
	}
	time.Sleep(5 * time.Second)
	// for _, obj := range objects {
	// }

	fmt.Println("Done")
}
