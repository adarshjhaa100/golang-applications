package concurrency

import (
	// "syn"
	"fmt"
	"sync"
)

/*
	Waitgroups are used to synchronize goroutines
*/


var counterWg = 0 // Note: Global variables can't be declared using shorthand
var wg = sync.WaitGroup{}

func incCounter(){
	counterWg++
	wg.Done() // Mark as done. This would decrease the number of groups by 1 
}

func displayCounter(){
	fmt.Printf("counter: %v\n", counterWg)
	wg.Done()
}

func wgImpl(){
	fmt.Println("Implement Wait Group: ")

	
	
	for i := 0; i < 10; i++ {
		wg.Add(2) // Add 2 groups to the wait groups (Will wait for 2 routines 
			// say done)
		go incCounter()
		go displayCounter()
	}
	// wait until groups are 0. Counter value would still be out of order
	// Use mutex or other methods to fix this.
	wg.Wait()
}

