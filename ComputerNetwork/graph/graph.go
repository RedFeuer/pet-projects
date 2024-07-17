package graph

import "fmt"

var D = 4

func AddBDS(x, y int) int {
	return x + y
}

func PrintSlice(s []int) {
	fmt.Printf("len = %d cap = %d %v", len(s), cap(s), s)	
}
