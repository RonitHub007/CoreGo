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
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	m     sync.Mutex
	value int
}

func (c *Counter) Update(n int, wg *sync.WaitGroup) {
	c.m.Lock()
	defer wg.Done()
	fmt.Printf("Adding %d to %d\n", n, c.value)
	c.value += n
	c.m.Unlock()
}

func main() {
	var wg sync.WaitGroup

	c := Counter{}

	wg.Add(4)

	go c.Update(10, &wg)
	go c.Update(-5, &wg)
	go c.Update(25, &wg)
	go c.Update(19, &wg)

	wg.Wait()
	fmt.Printf("Result is %d", c.value)
}
