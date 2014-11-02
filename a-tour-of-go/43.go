package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	var countsMap = make(map[string]int)
	for _, y := range strings.Fields(s) {
		countsMap[y]++
	}

	return countsMap
}

func main() {
	wc.Test(WordCount)
}
