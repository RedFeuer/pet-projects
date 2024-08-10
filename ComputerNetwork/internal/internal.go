package internal

import "container/list"

type Vertex struct {
	Comp      string // уникальное имя компьютера
	Port      uint   // номер порта компьютера
	Color     int    // цвет вершины: 0-белый, 1-серый, 2-черный
	Parent    *Node  // предок
	Path_size int    // длина пути от исходной вершины
	Component int    // номер компоненты связности, к которой относится вершина
}

type Edge struct {
	Ports       []uint // перечень допустимых для связи между собой по данному ребру портов
	Ports_count uint   // количество портов для данного ребра
}

/*ЛИНЕЙНЫЙ ОДНОСВЯЗНЫЙ СПИСОК(LINEAR SINGLE-LINKED LIST)*/
type AdjacentVertex struct {
	Vertex *Vertex // Указатель на данную смежную вершину
	Edge   *Edge   // Указатель на ребро, соединяющее основную и смежную к ней вершины
}

type Node struct {
	Vertex   *Vertex    // Указатель на данную основную вершину
	Adjacent *list.List // Указатель на смежные вершины
}

/*ХЭШ-ТАБЛИЦА(MAP)*/
type Graph struct {
	Table map[string]*Node // Хеш-таблица(мапа) для хранения вершин
}

type COLOR int

const (
	WHITE int = iota // 0 - белый
	GRAY             // 1 - серый
	BLACK            // 2 - черный
)

func Create_graph() *Graph {
	return &Graph{
		Table: make(map[string]*Node),
	}
}

func Initialize_node(new_vertex *Vertex) *Node {
	new_node := &Node{}
	new_node.Vertex = new_vertex
	return new_node
}

func Create_vertex(comp string, port uint) *Vertex {
	new_vertex := &Vertex{}
	new_vertex.Comp = comp
	new_vertex.Port = port
	return new_vertex
}

/*эта функция ищет, если ли у данной вершины с именем comp_src смежная вершина с именем comp_dst*/
/*возвращает элемент из списка, содержащий в поле value вершину с именем comp_dst*/
/*чтобы получить *AdjacentVertex, нужно сделать elem.value.(*AdjacentVertex)*/
func Find_adjacent_vertex(graph *Graph, comp_src string, comp_dst string) *list.Element {
	for comp, node := range graph.Table {
		if comp == comp_src {
			if node.Adjacent == nil {
				return nil
			}
			for elem := node.Adjacent.Front(); elem != nil; elem = elem.Next() {
				adjacent_vertex := elem.Value.(*AdjacentVertex)
				if adjacent_vertex.Vertex.Comp == comp_dst {
					return elem
				}
			}
			/*скорее всего return 1 лучше сделать тут, чтобы быстрее работало*/
			/*ЕСЛИ КАКАЯ-ТО НЕВЕРНАЯ ОШИБКА ПРИ ДОБАВЛЕНИИ РЕБРА, ТО ОНА СКОРЕЕ ВСЕГО ТУТ*/
			return nil
		}
	}
	return nil
}

func Remove_Adjacent_Vertex(graph *Graph, comp_remove string) {
	for _, node := range graph.Table {
		if node.Adjacent != nil {
			for elem := node.Adjacent.Front(); elem != nil; elem = elem.Next() {
				adjacent_vertex := elem.Value.(*AdjacentVertex)
				if adjacent_vertex.Vertex.Comp == comp_remove {
					node.Adjacent.Remove(elem)
				}
			}
		}
	}
}
