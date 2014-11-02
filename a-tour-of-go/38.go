package main

import (
	"code.google.com/p/go-tour/pic"
	"math/rand"
)

func Pic(dx, dy int) [][]uint8 {

	data:= make([][]uint8,dy)
	for i := range data {
		data[i] = make([]uint8,dx)
		for j := range data[i] {
			data[i][j] = uint8(rand.Int())
		}
	}

	return data
}

func main() {
	pic.Show(Pic)
}
