package main

import (
	"fmt"
)

func split(sum int64) (x,y int64) {
	x = sum * 5 / 10
	y = sum - x
	return
}

func main() {
	fmt.Println(split(63))
}
