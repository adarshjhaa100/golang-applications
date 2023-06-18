package concurrency


import (
	"sync"
	"fmt"
)


/*
	Please note: Mutex lock is a blocking operation, which at times maybe 
	against the idea of concurrency and one goroutine may block other
*/
type CounterMut struct {
	mu sync.RWMutex // Lock can be held by arbitrary number of reads or 1 write
	ctr int
}



var counterWgM = 0 // Note: Global variables can't be declared using shorthand
var wgM = sync.WaitGroup{}
var mtx = sync.Mutex{} // this would apply to all global variables




func incCounterM(){
	defer mtx.Unlock() // defer until out of scope
	counterWgM++
	wgM.Done() // Mark as done. This would decrease the number of groups by 1 
}

func displayCounterM(){
	defer mtx.Unlock()
	fmt.Printf("counter: %v\n", counterWgM)
	wgM.Done()
}

func mutexImpl(){
	fmt.Println("Implement Wait Group: ")
	for i := 0; i < 10; i++ {
		wgM.Add(2) // Add 2 groups to the wait groups (Will wait for 2 routines 
			// say done)
		// Lock for write. This lock is in the caller for better control
		// mutex inside the callee could make things out of order as 
		// goroutines could be called different amount of times giving diff value 
		mtx.Lock() 
 		go incCounterM()
		mtx.Lock() 
		go displayCounterM()
	}
	// wait until groups are 0. Counter value would still be out of order
	// Use mutex or other methods to fix this.
	wgM.Wait()
}
