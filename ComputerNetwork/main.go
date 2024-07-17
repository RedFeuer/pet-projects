package main

import (
	"fmt"
	"ComputerNetwork/graph"
	"ComputerNetwork/logic"
	//"github.com/RedFeuer/pet-projects/graph"
)

// func add(x, y int) int {
// 	return x + y
// }

func main() {
	var a int
	fmt.Scanf("%d", &a)
	var b int
	fmt.Scanf("%d", &b)
	res1 := graph.AddBDS(a,b)
	res2 := logic.SumBDS(a, b)
	fmt.Printf("%d == %d", res1, res2)
}
