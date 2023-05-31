package main

import (
	"fmt"
	qt "rsc.io/quote"
	// "github.com/tcsc/langtutor/varsfncs"
	"github.com/tcsc/langtutor/flowcontrol"
	// "time"
)


func main() {
	fmt.Println("Hello, World!")
	fmt.Println(qt.Go()) 
	// fmt.Println(time.Now().UnixNano())

	// Variables ,functions, types
	// varsfncs.SimpleFunc()
	// fmt.Println(varsfncs.MultiReturn())
	// varsfncs.VarTypes()

	// Flow Control
	// flowcontrol.ForLoop()
	// flowcontrol.UsingSwitchInsteadIf()
	flowcontrol.CallSqrt()

}