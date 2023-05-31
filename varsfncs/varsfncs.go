package varsfncs

import "fmt"

func SimpleFunc() {
	fmt.Println("This is a single function")
}

func MultiReturn() (int, string){
	return 1, "hello"
}

func VarTypes() {

	var (
		bl bool = true //true/false
		st string = "hello" //string
		it int32 = 456546 // 32 bit integer
		unit uint64 = 2323 // 64 bit unsigned int
		flt32 float32 = 3.4534e12 - 11111	
		flt64 float64 = float64(54-678) //explicit type conversion
	)

	fmt.Println(bl, st, it, unit, flt32, flt64)
}

