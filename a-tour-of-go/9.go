package main

import (
	"fmt"
)

func swap(x, y string) (string,string) {
	return y,x
}

func main() {
	a,b := swap("This goes last","this goes first")

	fmt.Println(a,b)

	fmt.Println(swap("2","1"))
}
