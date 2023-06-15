package methodsinterface

import "fmt"


type display interface {
	Show() string
}

// type *V implements Show() method
func (v *Vehicle) Show() string{
	return v.display()
}


func describe(v any){
	fmt.Printf("(%v, %T)\n", v, v)
}


func ImplInterface(){
	var v display // nil type and value

	// Interface can be thought as a tuple of concrete value and type : (val, type)
	describe(v)

	// v = Vehicle{} // this is wrong as method Show() has pointer receiver. Need to pass address
	v = &Vehicle{}

	// nil underlying value
	describe(v)

	// type assertions
	var v1 interface{} = 12
	v11, ok := v1.(string) // this would be false

	fmt.Println(v11, ok)






}