package main

import (
	"fmt"
	"graph"
	//"github.com/RedFeuer/pet-projects/graph"
)

func main() {
	var a int
	var b int
	fmt.Scanf("%d", &a)
	fmt.Printf("%d", a)
	fmt.Scanf("%d", &b)
	fmt.Printf("%d", b)
	fmt.Printf("%d", graph.sum(a, b))
}
