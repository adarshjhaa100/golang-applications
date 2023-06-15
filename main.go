package main

import (
	"fmt"

	qt "rsc.io/quote"

	// "github.com/tcsc/langtutor/varsfncs"
	// "github.com/tcsc/langtutor/flowcontrol"
	// "github.com/tcsc/langtutor/ioexamples"
	// "time"
	// "github.com/tcsc/langtutor/datastructs"
	// "github.com/tcsc/langtutor/aoc"
	methodsinterface "github.com/tcsc/langtutor/methodsInterface"
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
	// flowcontrol.CallSqrt()
	// flowcontrol.RunAllDefers()
	// flowcontrol.DeferAndRecover()

	// IO
	// ioexamples.CallFileIOFunctions()
	// datastructs.DataStructsPrimitives()

	// aoc.SolveQuestion2021("2021_1")
	// aoc.SolveQuestion2021("2021_2")
	// methodsinterface.CallMethodVehicle()
	methodsinterface.ImplInterface()  
}