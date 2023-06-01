package flowcontrol

import (
	"fmt"
	"time"
)


func deferTest(a *int){
	*a = 1
	defer func(){
		*a = 0		
	}()
	
	for i := 0; i < 10; i++ {
		*a++
		fmt.Println(*a)
	}
}

// All defers in B will be called, but not normal code after the panicked part
func deferB() {
	fmt.Println("deferB start")
	// recover from panic, defer should be called before a statement that's possible to cause a panic
	defer func(){
		if r := recover(); r != nil {
			fmt.Println("Recovered panicking in deferB", r)
		}
	}()

	panicInMiddle()
	fmt.Println("deferB end")
}

// A will function normally (Since defer in B has handled the panic)
func deferA() {
	fmt.Println("defer A Start")
	defer func(){
		fmt.Println("defer A1")
	}()
	defer func(){
		fmt.Println("defer A2")
		if r := recover(); r != nil {
			fmt.Println("Recovered panicking in deferA", r)
		}

	}()
	defer func(){
		fmt.Println("defer A3")
	}()
	
	deferB()

	fmt.Println("Defer A Stack Called...")
}

func panicInMiddle() {
	defer func(){
		fmt.Println("defer B2")
		if r := recover(); r != nil {
			fmt.Println("Recovered panicking in deferC", r)
		}

	}()
	for i := 10; i<= 100; i+=10 {
		fmt.Println("i=",i)
		if(i>50){
			panic(fmt.Sprintf("panicked for %v\n", i))
		}
	}
}





func RunAllDefers(){
	a := 50
	go deferTest(&a)

	time.Sleep(8000000000)

	for {
		fmt.Println(a)
	}
}

// A defer function in the call stack will recover from panic
// No other non deferred code will be executed for that current function where recovery was made.
// Parent function will run normally as expected
func DeferAndRecover() {
	deferA()
	// deferB() 
}