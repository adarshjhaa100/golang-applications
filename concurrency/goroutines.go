package concurrency

import (
	"fmt"
	"time"
)

func primaryGoroutine(){

	msgs := "Hello"

	// Create a goroutine
	go func(){
		fmt.Println(msgs) // this is a race condition
	}()
	
	msgs = "World"
	time.Sleep(time.Second * 1) // Just for testing. 
	// Dont use in production for this purpose

}

