package graph

import "fmt"

var D = 4

func AddBDS(x, y int) int {
	return x + y
}

func PrintSlice(s string,x []int) {
	fmt.Printf("%s: len = %d cap = %d %v\n", s, len(x), cap(x), x)	
}
