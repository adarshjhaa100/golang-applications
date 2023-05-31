package flowcontrol

import (
	"fmt"
)

func ForLoop() {
	for i := 0; i<10; i++ {
		fmt.Println("Iteration: #", i)
	}

	// for can be used as a while with only the conditional
	i:=0
	for i<5 {
		fmt.Println("Using as a while loop...")
		i++
		break
	}
}


// class based on salary. A if salary > 80000, B if salary >50000 else C
func UsingSwitchInsteadIf() {
	
	salary := 12000
	salaryClass := ""

	switch {
		case salary > 80_000:
			salaryClass = "A"
		case salary > 50_000:
			salaryClass = "B"
		case salary <= 50_000:
			salaryClass = "C"
		default:
			fmt.Println("Some issue in value of salary") 		
	}
	fmt.Println("Salary class is: ", salaryClass)

}