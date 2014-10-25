package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %v problems.",
		math.Nextafter(5, 6))
}
