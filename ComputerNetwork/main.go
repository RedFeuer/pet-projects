package main

import (
	"fmt"
    g "ComputerNetwork/graph"
	// "ComputerNetwork/logic"
)

// type Vertex struct {
// 	comp string                   // уникальное имя компьютера
// 	port uint                     // номер порта компьютера
// 	color COLOR                   // цвет вершины: 0-белый, 1-серый, 2-черный
// 	parent *Node                  // предок
// 	path_size int                 // длина пути от исходной вершины
// 	component int                 // номер компоненты связности, к которой относится вершина
// } 

// type Edge struct {
// 	ports []uint                  // перечень допустимых для связи между собой по данному ребру портов
// 	ports_count int               // количество портов для данного ребра
// }

// /*ЛИНЕЙНЫЙ ОДНОСВЯЗНЫЙ СПИСОК(LINEAR SINGLE-LINKED LIST)*/
// type AdjacentVertex struct {
// 	vertex *Vertex                // Указатель на данную смежную вершину
// 	edge *Edge                    // Указатель на ребро, соединяющее основную и смежную к ней вершины
// 	next *AdjacentVertex          // Следующая в списке смежная вершина(смежная к основной)
// }

// type Node struct {
// 	vertex *Vertex                // Указатель на данную основную вершину
// 	adjacent *AdjacentVertex      // Указатель на смежные вершины
// }

// /*ХЭШ-ТАБЛИЦА(MAP)*/
// type Graph struct {
// 	table map[string]*Node        // Хеш-таблица(мапа) для хранения вершин
// }

type COLOR int
const (
	WHITE COLOR = iota            // 0 - белый
	GRAY                          // 1 - серый
 	BLACK                         // 2 - черный
)

func main() {
	/*создаем граф*/
	graph := g.Create_graph()

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
