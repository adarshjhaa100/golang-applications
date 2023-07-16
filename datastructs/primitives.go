package datastructs

import (
	"fmt"
)

type Vehicle struct {
	name   string
	price  float32
	typeOf string // "car"|"bike"|"truck" etc.
}

func pointers() {
	// No pointer arithmetic in go
	a, b := 12, 13

	ptrInt := &a

	fmt.Printf("value(a): %#v, ptr: %#v, value(b): %#v, ptr: %#v\n", a, ptrInt, b, &b)
}

func structImpl() {

	v1 := Vehicle{}                            //Vehicle is struct literal with all default values
	v2 := Vehicle{"WagonR", 32545, "car"}      // values are in the order of fields
	v3 := Vehicle{price: 32545, typeOf: "car"} // When assigning less fields, need to mention the attribute

	fmt.Printf("v1: %#v, v2: %#v, v3: %#v\n", v1, v2, v3)

	// Assigning to a new struct copies the data
	v4 := v1
	v4.price = 1234
	fmt.Printf("post modification of v4: v4 = %#v, v1 = %#v\n", v4, v1)

	// pointer to a struct: Often times a complex data structure takes a significant amount of memory space.
	// Passing to a variable, assigns a copy in memory. So instead we need to pass memory address
	ptrVehicle := &v2

	fmt.Printf("ptr to v2: %#v, adderss: %p\n", ptrVehicle, ptrVehicle)
	fmt.Printf("v2 price using ptr: %#v\n", ptrVehicle.price) // For Structs: can dereference w/o explicitly using (*)

}

func arrays() {
	// Arrays hold fixed number of elements in memory and declared using [n]T
	a := [5]int{1, 2, 3, 4} // last element would be 0
	fmt.Printf("Array a: %#v\n", a)
}

// NOTE: Array/Slices are passed as reference by default
func sliceIntro() {
	arr := [5]int{3, 4, 5, 6, 7}

	// Slices are a preview into a part of underlying array. Slice header is of the format: (*ptr, len, cap)
	sl1 := arr[:] // here ptr is at first element of array with len = 5 and cap = 5
	printSlice[int](sl1)

	// modifying an element in slice actually modifies the element in the underlying arr
	sl1[3] = 60
	fmt.Printf("Array arr post modifying element at 3: %#v\n", arr)

	// re-slicing
	sl1 = sl1[:2] // ptr: &arr[0], len = 2, cap = 5
	printSlice[int](sl1)

	// extend the slice. Can extend upto capacity
	sl1 = sl1[:4] // ptr: &arr[0], len = 2, cap = 5
	printSlice[int](sl1)

	// drop first element
	sl1 = sl1[1:] // This would move ptr to arr[1]m reducing cap and len. ptr: &arr[0], len = 1, cap = 4
	printSlice[int](sl1)

	// Create slice using make
	sl2 := make([]int, 5, 10) // The underlying array would be created. ptr = 0th, len=5, cap=5
	printSlice[int](sl2)

	// nil slice. Slice created w/o assignment is nil with 0 length and capacity
	var sl3 []int
	printSlice[int](sl3)

}

func sliceResizing(){
	// On reaching the full capacity,slices can be resized by copying the
	// underlying elements to a new array of 2*(cap + 1) size
	sl1 := []int{1,2,3,4,5}
	printSlice[int](sl1)

	//copy to a new slice with twice the cap. The size has cap + 1 for case where capacity is 0
	sl2 := make([]int, 2*(cap(sl1)+1))
	copy(sl2, sl1)
	printSlice[int](sl2)
	
	// appending to slice. Appending to slice appends to last element of
	// underlying array if enough capacity. Else new underlying array is created
	sl3 := make([]int, 1, 4)
	sl3[0] = 1
	printSlice[int](sl3)
	printElementWithAddress(&sl3[len(sl3)-1])

	// Using range to loop over array/slice etc.
	for _,val := range []int{4,5,6,7,8} {

		if len(sl3) >= cap(sl3) {
			fmt.Println("New underlying array will be created")
		}

		sl3 = append(sl3, val)
		printSlice[int](sl3)
		printElementWithAddress(&sl3[0])
		printElementWithAddress(&sl3[len(sl3)-1])
		
	}

	// Can append one slice to other using ... to convert it to set of value 
	// parameters. 
	var sl4 []int
	sl4 = append(sl4, sl3...)
	printSlice[int](sl4)

}

func printSlice[T any](sl []T) {
	fmt.Printf("slice: %#v, length: %v, capacity: %v\n", sl, len(sl), cap(sl))
}

func printElementWithAddress[T any](ele *T) {
	fmt.Printf("element value: %#v, address: %#v\n", *ele, ele)
}

func mapsImplement(){
	// Map is a set of key value pair which can be created using make
	mp := make(map[string]Vehicle)
	fmt.Printf("Map: %#v \n", mp)

	// nil map (cannot assign to a nil map)
	var mpnil map[string]string
	fmt.Printf("Map: %#v \n", mpnil)

	// Insert to map
	mp["1"] = Vehicle{"WagonR",12,"car"}
	mp["3"] = Vehicle{"Ertiga",13,"car"}
	
	fmt.Printf("Map: %v \n", mp)


	// retrieve an element while checking if key exists
	ele, ok := mp["1"]
	fmt.Printf("element: %v, address: %p, key present: %v\n", ele, &ele, ok)

	ele1, ok1 := mp["3"]
	fmt.Printf("element: %v, address: %p, key present: %v\n", ele1, &ele1, ok1)

	ele, ok = mp["2"] // Here we would get a zero
	fmt.Printf("element: %v, address: %p, key present: %v\n", ele, &ele, ok)

	delete(mp,"3") // delete key
	fmt.Printf("Map: %v \n", mp)

	// Map to a slice
	sl := []int{1,2,3}
	mpsl := make(map[string][]int)
	mpsl["car1"] = sl // This would be a reference to the slice instead of copy
	fmt.Printf("Map: %#v \n", mpsl)

	sl[2] = 30 // modifying slice element
	fmt.Printf("Map: %#v \n", mpsl) // this would display the modified slice
}


func initWithNew(){
	// Go has new keyword which can be used to initialize with zeros and return 
	// pointer
	vh := new(Vehicle)

	// this is similar to
	vh = &Vehicle{}

	fmt.Println(vh)
}


func DataStructsPrimitives() {
	// pointers()
	// structImpl()
	// arrays()
	// sliceIntro()
	// sliceResizing()
	mapsImplement()
}
