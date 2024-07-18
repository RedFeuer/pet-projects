package graph

// import "fmt"


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