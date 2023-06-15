package methodsinterface

import (
	"fmt"
)

type Vehicle struct {
	model string	
	typeOf string
	hp int
	price float32
}

// v is a pointer receiver
func (v *Vehicle) display() string{
	fmt.Printf("Vehicle Description:\nModel: %v\nType: %v\nHorse Power: %v\nPrice: %v\n", 
			v.model, 
			v.typeOf,
			v.hp,
			v.price)

			return fmt.Sprintf("Vehicle Description:\nModel: %v\nType: %v\nHorse Power: %v\nPrice: %v\n", 
			v.model, 
			v.typeOf,
			v.hp,
			v.price)

}

func CallMethodVehicle(){
	var v Vehicle = Vehicle{
		model : "FE1221GCV",
		typeOf: "CAR",
		hp: 1234,
		price: 3204332,  
	}

	// Implicitly means ( (&v).display() )
	v.display()


}
