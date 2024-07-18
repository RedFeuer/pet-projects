package main

import (
	"fmt"

	d "github.com/RedFeuer/pet-projects/ComputerNetwork/dialog"
)

// type COLOR int
// const (
// 	WHITE COLOR = iota            // 0 - белый
// 	GRAY                          // 1 - серый
//  	BLACK                         // 2 - черный
// )

func main() {
	/*создаем граф*/
	graph := g.Create_graph()

	d.D1_Insert_Vertex(graph)

	for i := 0; i < 1000; i++ {
        compName := fmt.Sprintf("computer%d", i)
        node := &g.Node{
            Vertex: &g.Vertex{
                Comp:      compName,
                Port:      uint(8000 + i),
                Color:     0,
                Path_size:  0,
                Component: i+1,
            },
        }
        graph.Table[compName] = node
    }

    // Проверка количества вершин в графе
    fmt.Println("Number of vertices in the graph:", len(graph.Table))

    // Доступ к одной из вершин
    if n, exists := graph.Table["computer500"]; exists {
        // Работа с найденной вершиной
        fmt.Println("Vertex:", n.Vertex.Comp)
        fmt.Println("Port:", n.Vertex.Port)
    } else {
        fmt.Println("Vertex not found")
    }
}
