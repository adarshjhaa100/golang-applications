package concurrency

import (
	"fmt"
	"runtime"
	// "time"
)

/*
	Race Condition: Parallel read/write on same variable during Write operation
	To check for race condition during compile time: use -race flag
	*/

func RunConcurrency(){
	fmt.Println("Hello Concurrency Package")
	runtime.GOMAXPROCS(100)
	fmt.Printf("Runtime MaxProcs: %#v\n", runtime.GOMAXPROCS(-1)) // MAx CPU threads that can be utilized. -1 does not change the setting, else n = n threads
	// time.Sleep(time.Second*100)

	// primaryGoroutine()
	// wgImpl()
	mutexImpl()
}