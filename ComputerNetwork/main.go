package main

import (
	// "fmt"
	"ComputerNetwork/graph"
	// "ComputerNetwork/logic"
)

func main() {
	a := make([]int, 5)	
	graph.PrintSlice("a", a)

	b := make([]int, 0, 5)
	graph.PrintSlice("b", b)

	c := a[:2]
	graph.PrintSlice("c", c)
	c[0] = 1
	c[1] = 2
	graph.PrintSlice("new c", c)
	graph.PrintSlice("new a", a)

	d := b[:2]
	graph.PrintSlice("d", d)

	e := c[2:5]
	graph.PrintSlice("e", e)

	
}
