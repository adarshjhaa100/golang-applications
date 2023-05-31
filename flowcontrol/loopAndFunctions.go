package flowcontrol

import (
	"fmt"
	"math"
	"time"
)

func Sqrt(x float64) (float64, int32, int64) {

	z := float64(1)
	maxDelta := float64(1e-5)
	iter := int32(0)
	unixNow := time.Now().UnixNano()

	for ; iter<= 1000; iter++ {
		
		delta := (z*z - x)/(2*z)
		// fmt.Printf("z = %#v, delta = %#v\n", z, math.Abs(delta))
		z = z - delta
		
		if(math.Abs(delta) < maxDelta) {
			return z, iter, time.Now().UnixNano() - unixNow
		}
	}

	return z, iter, time.Now().UnixNano() - unixNow
}


func CallSqrt(){
	nums := []float64 {10, 20, 30, 1000,  56, 2, 100_000_000, 47350894350943534753045 }
	
	for _, num:=range(nums){
		sqrt, iters, timeElapsed := Sqrt(num)
		fmt.Printf("SQRT of %v is %v, iterations: %v, timeElapsed: %vns\n", num, sqrt, iters, timeElapsed )
	}

}