package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {

	z := float64(1.0)
//	y := z*2

//	delta := 0.0005

	for i := 1 ; i < 10  ; i++ {

//		y = z
		z -= ((z*z)-x)/(2*z)

	}
	return z
}

func main() {
	fmt.Println(Sqrt(20))
}
