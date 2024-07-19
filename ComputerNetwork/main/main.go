package main

import (
	"fmt"

	//d "github.com/RedFeuer/pet-project988888888s/ComputerNetwork/dialog"
)

type Vertex struct {
	Comp string                   // уникальное имя компьютера
	Port uint                     // номер порта компьютера
	Color int                  // цвет вершины: 0-белый, 1-серый, 2-черный
	Parent *Node                  // предок
	Path_size int                 // длина пути от исходной вершины
	Component int                 // номер компоненты связности, к которой относится вершина
} 

type Edge struct {
	Ports []uint                  // перечень допустимых для связи между собой по данному ребру портов
	Ports_count int               // количество портов для данного ребра
}

/*ЛИНЕЙНЫЙ ОДНОСВЯЗНЫЙ СПИСОК(LINEAR SINGLE-LINKED LIST)*/
type AdjacentVertex struct {
	Vertex *Vertex                // Указатель на данную смежную вершину
	Edge *Edge                    // Указатель на ребро, соединяющее основную и смежную к ней вершины
	Next *AdjacentVertex          // Следующая в списке смежная вершина(смежная к основной)
}

type Node struct {
	Vertex *Vertex                // Указатель на данную основную вершину
	Adjacent *AdjacentVertex      // Указатель на смежные вершины
}

/*ХЭШ-ТАБЛИЦА(MAP)*/
type Graph struct {
	Table map[string]*Node        // Хеш-таблица(мапа) для хранения вершин
}

type COLOR int
const (
	WHITE COLOR = iota            // 0 - белый
	GRAY                          // 1 - серый
 	BLACK                         // 2 - черный
)


func Create_graph() *Graph {
	return &Graph{
		Table: make(map[string]*Node),
	}
}

func Initialize_node(new_vertex *Vertex) *Node {
	var new_node *Node
	new_node.Vertex = new_vertex
	return new_node
}

func Create_vertex(comp string, port uint) *Vertex {
	var new_vertex *Vertex
	new_vertex.Comp = comp
	new_vertex.Port = port
	return new_vertex
}

func D1_Insert_Vertex(graph *Graph) {
	fmt.Printf("Enter unique computer name: ")
	var comp string
	fmt.Scan(&comp)
	if len(comp) == 0 {
		/*ПОДУМАТЬ КАК СДЕЛАТЬ ОБРАБОТКУ ОШИБОК КРАСИВЕЙ*/
		fmt.Printf("ERROR: Computer-name is empty\n");
		return
	}
	fmt.Printf("Enter number of port for computer %s: ", comp)
	var port uint
	fmt.Scanf("%u\n", &port)

	/*ПРОВЕРКА: СУЩЕСТВУЕТ ЛИ ЭЛЕМЕНТ С ТАКИМ КЛЮЧОМ В ГРАФЕ*/
	// var elem *Node
	// var ok bool
	new_vertex := Create_vertex(comp, port)
	new_node := Initialize_node(new_vertex)
	graph.Table[comp] = new_node
}

func main() {
	/*создаем граф*/
	graph := Create_graph()

	var flag int = 1
	for flag == 1 {
		var choice int
		switch choice {
			default : // ВЫХОД ИЗ ПРОГРАММЫ
				flag = 0
			case 1 :
				D1_Insert_Vertex(graph)
		} 
	}
	// for i := 0; i < 1000; i++ {
    //     compName := fmt.Sprintf("computer%d", i)
    //     node := &Node{
    //         Vertex: &Vertex{
    //             Comp:      compName,
    //             Port:      uint(8000 + i),
    //             Color:     0,
    //             Path_size:  0,
    //             Component: i+1,
    //         },
    //     }
    //     graph.Table[compName] = node
    // }

    // // Проверка количества вершин в графе
    // fmt.Println("Number of vertices in the graph:", len(graph.Table))

    // // Доступ к одной из вершин
    // if n, exists := graph.Table["computer500"]; exists {
    //     // Работа с найденной вершиной
    //     fmt.Println("Vertex:", n.Vertex.Comp)
    //     fmt.Println("Port:", n.Vertex.Port)
    // } else {
    //     fmt.Println("Vertex not found")
    // }
}
