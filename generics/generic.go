package generics

import "fmt"

type KeyValPair struct{
	key string
	val any
}

func GenericFunc[T any](arr *[]T){
	fmt.Printf("Printing slice of type %T\n", arr)
	for _, val := range *arr {
		fmt.Println(val)
	}
}

type GenericType[T any] struct{
	key string
	value T
}

func ImplGeneric(){

	arrInt := []int {1,2,3,5}

	arrKeyVal := []KeyValPair{
		{"abc", 23432},
		{"def", 4567},
	}

	GenericFunc[int](&arrInt)
	GenericFunc[KeyValPair](&arrKeyVal)

	gt := GenericType[KeyValPair]{
		"hello",
		KeyValPair{"tmp", 4567},
	}

	fmt.Printf("Generic type val: %#v \n", gt)
	fmt.Printf("Generic type val.value: %#v \n", gt.value)
}

