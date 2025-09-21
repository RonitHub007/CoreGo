// // package main

// // import (
// // 	"fmt"
// // 	"sync"
// // )

// // type Person struct {
// // 	Name string
// // 	Age  int
// // 	Data []byte
// // }

// // var personPool = sync.Pool{
// // 	New: func() interface{} {
// // 		return &Person{
// // 			Data: make([]byte, 1024),
// // 		}
// // 	},
// // }

// // func main() {

// // 	people := make([]*Person, 0, 10000)

// // 	for i := 0; i < 10000; i++ {
// // 		p := personPool.Get().(*Person)
// // 		p.Name = fmt.Sprintf("User%d", i)
// // 		p.Age = 20 + i
// // 		people = append(people, p)
// // 	}
// // 	for _, p := range people {
// // 		fmt.Println(p.Name, p.Age, len(p.Data))

// // 		// Reset fields if needed (to avoid old data leaking)
// // 		p.Name = ""
// // 		p.Age = 0

// // 		// Put back in pool
// // 		personPool.Put(p)
// // 	}
// // }

// // // package main

// // // import "fmt"

// // // type Person struct {
// // // 	Name string
// // // 	Age  int
// // // 	Data []byte
// // // }

// // // func newPerson(name string, age int) *Person {
// // // 	return &Person{
// // // 		Name: name,
// // // 		Age:  age,
// // // 		Data: make([]byte, 1024), // allocate 1KB buffer
// // // 	}
// // // }

// // // func main() {
// // // 	people := make([]*Person, 0, 10000)
// // // 	for i := 0; i < 10000; i++ {
// // // 		p := newPerson(fmt.Sprintf("User%d", i), 20+i)
// // // 		people = append(people, p)
// // // 	}

// // // 	for _, p := range people {
// // // 		fmt.Println(p.Name, p.Age, len(p.Data))
// // // 	}
// // // }

// // Built a super over system where over will be stopped after the wicket falls

// // package main

// // import (
// // 	"fmt"
// // 	"sync"
// // )

// // type Database struct {
// // 	Name string
// // }

// // var (
// // 	db   *Database
// // 	once sync.Once
// // )

// // func connectedDB() *Database {
// // 	once.Do(func() {
// // 		fmt.Println("Connection establishing")
// // 		db = &Database{Name: "MongoDB"}
// // 	})
// // 	return db
// // }
// // func main() {
// // 	var wg sync.WaitGroup

// // 	wg.Add(5)

// // 	for i := 0; i < 5; i++ {
// // 		go func(id int) {
// // 			defer wg.Done()
// // 			conn := connectedDB()
// // 			fmt.Printf("GoRoutine %d got DB: %s\n", id, conn.Name)
// // 		}(i)
// // 	}

// // 	wg.Wait()

// // }

// // Golang Atomic sync pacakge

// // package main

// // import (
// // 	"fmt"
// // 	"sync"
// // 	"sync/atomic"
// // 	"time"
// // )

// // func main() {
// // 	start := time.Now()
// // 	var counter int32
// // 	var wg sync.WaitGroup
// // 	for i := 0; i < 100000; i++ {
// // 		wg.Add(1)
// // 		go func() {
// // 			atomic.AddInt32(&counter, 1)
// // 			wg.Done()
// // 		}()
// // 	}
// // 	wg.Wait()
// // 	end := time.Since(start)
// // 	fmt.Println("Counter", atomic.LoadInt32(&counter))
// // 	fmt.Println("timetaken", end)
// // }

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var (
// 	cond *sync.Cond
// 	mu   sync.Mutex
// 	data int
// )

// func producer() {
// 	for i := 1; i <= 5; i++ {
// 		mu.Lock()
// 		data = i
// 		cond.Broadcast()
// 		mu.Unlock()
// 		time.Sleep(500 * time.Microsecond)
// 	}
// }

// func consumer() {
// 	for i := 1; i <= 5; i++ {
// 		mu.Lock()
// 		for data != i {
// 			cond.Wait()
// 		}
// 		mu.Unlock()
// 		fmt.Println(data)
// 		time.Sleep(500 * time.Microsecond)
// 	}
// }
// func main() {
// 	cond = sync.NewCond(&mu)
// 	go producer()
// 	go consumer()
// 	time.Sleep(500 * time.Microsecond)
// 	fmt.Println("Sync condition checking")
// }

package main

func main() {

}
