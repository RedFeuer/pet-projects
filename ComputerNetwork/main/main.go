package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
	//d "github.com/RedFeuer/pet-project988888888s/ComputerNetwork/dialog"
)

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
	// Next   *AdjacentVertex // Следующая в списке смежная вершина(смежная к основной)
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
	WHITE COLOR = iota // 0 - белый
	GRAY               // 1 - серый
	BLACK              // 2 - черный
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

func Insert_Edge_logic(graph *Graph, comp_src string, comp_dst string, ports_count uint, ports []uint) int {
	/*проверка: одной из вершин нет*/
	if graph.Table[comp_src] == nil || graph.Table[comp_dst] == nil {
		return 1
	}
	/*проверка: не соединяем одинаковые вершины*/
	if graph.Table[comp_src] == graph.Table[comp_dst] {
		return 2
	}
	/*проверка: между соединяемыми вершинами не существует ребра*/
	if Find_adjacent_vertex(graph, comp_src, comp_dst) != nil || Find_adjacent_vertex(graph, comp_dst, comp_src) != nil {
		return 3
	}
	//if graph.Table[comp_src].Adjacent != nil && graph.Table[comp_dst].Adjacent != nil {
	//	return 3
	//}

	/*соединяем основную вершину со смежной: src -> dst*/
	if graph.Table[comp_src].Adjacent == nil {
		graph.Table[comp_src].Adjacent = list.New()
	}
	new_adjacent_for_src := &AdjacentVertex{}
	new_adjacent_for_src.Edge = &Edge{ports, ports_count}
	new_adjacent_for_src.Vertex = graph.Table[comp_dst].Vertex
	graph.Table[comp_src].Adjacent.PushBack(new_adjacent_for_src)

	/*обратное соединение: dst -> src*/
	if graph.Table[comp_dst].Adjacent == nil {
		graph.Table[comp_dst].Adjacent = list.New()
	}
	new_adjacent_for_dst := &AdjacentVertex{}
	new_adjacent_for_dst.Edge = &Edge{ports, ports_count}
	new_adjacent_for_dst.Vertex = graph.Table[comp_src].Vertex
	graph.Table[comp_dst].Adjacent.PushBack(new_adjacent_for_dst)

	return 0
}

func Remove_Edge_Logic(graph *Graph, comp_src string, comp_dst string) int {
	/*ДЕЛАТЬ ПРОВЕРКИ*/
	if graph.Table[comp_src] == nil {
		return 1
	}
	if graph.Table[comp_dst] == nil {
		return 2
	}

	elem_remove_src := Find_adjacent_vertex(graph, comp_src, comp_dst)
	if elem_remove_src == nil {
		return 3
	}
	graph.Table[comp_src].Adjacent.Remove(elem_remove_src)

	elem_remove_dst := Find_adjacent_vertex(graph, comp_dst, comp_src)
	if elem_remove_dst == nil {
		return 3
	}
	graph.Table[comp_dst].Adjacent.Remove(elem_remove_dst)
	return 0
}

func D1_Insert_Vertex(graph *Graph) {
	reader := bufio.NewReader(os.Stdin)
	comp := Read_non_empty_string(reader, "Enter unique computer name: ")
	fmt.Printf("Enter number of port for computer %s", comp)
	port := Read_integer(reader, ": ")
	new_vertex := Create_vertex(comp, port)
	new_node := Initialize_node(new_vertex)
	graph.Table[comp] = new_node
}

func D2_Insert_Edge(graph *Graph) {
	reader := bufio.NewReader(os.Stdin)
	comp_src := Read_non_empty_string(reader, "Enter computer name for source computer: ")
	comp_dst := Read_non_empty_string(reader, "Enter computer name for destination computer: ")

	ports_count := Read_integer(reader, "Enter number of ports: ")
	ports := make([]uint, ports_count)
	for i := 0; i < int(ports_count); i++ {
		ports[i] = Read_integer(reader, "Enter port: ")
	}

	termination_status := Insert_Edge_logic(graph, comp_src, comp_dst, ports_count, ports)
	switch termination_status {
	default:
		break
	case 1:
		fmt.Println("Error: There no such vertex in graph")
	case 2:
		fmt.Println("Error: can't connect vertex to itself")
	case 3:
		fmt.Println("Error: can't connect vertices that have been already connected")
	}
}

func D4_Remove_Edge(graph *Graph) {
	reader := bufio.NewReader(os.Stdin)
	comp_src := Read_non_empty_string(reader, "Enter computer name for source computer: ")
	comp_dst := Read_non_empty_string(reader, "Enter computer name for destination computer: ")
	termination_status := Remove_Edge_Logic(graph, comp_src, comp_dst)
	switch termination_status {
	case 1:
		fmt.Printf("Error: There no computer with name %s\n", comp_src)
	case 2:
		fmt.Printf("Error: There no computer with name %s\n", comp_dst)
	case 3:
		fmt.Printf("Error: There is no edge between %s and %s \n", comp_src, comp_dst)
	}
}

func D7_Output_as_adjacency_list(graph *Graph) {
	for comp, node := range graph.Table {
		fmt.Printf("Computer name: %s Port: %d  ", comp, node.Vertex.Port)
		if node.Adjacent == nil {
			fmt.Printf("\n")
			continue
		}
		for elem := node.Adjacent.Front(); elem != nil; elem = elem.Next() {
			adjacent_vertex := elem.Value.(*AdjacentVertex)
			fmt.Printf(" ->  %s", adjacent_vertex.Vertex.Comp)
		}
		fmt.Printf("\n")
	}
}

func Read_non_empty_string(reader *bufio.Reader, prompt string) string {
	var result string
	var err error

	for {
		fmt.Print(prompt)
		result, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Try again.")
			continue
		}
		result = strings.TrimSpace(result)
		if result == "" {
			fmt.Println("Error: Empty string. Try again.")
			continue
		}
		break
	}
	return result
}

func Read_integer(reader *bufio.Reader, prompt string) uint {
	var result uint
	var err error

	for {
		fmt.Print(prompt)
		_, err = fmt.Scan(&result)
		if err != nil {
			fmt.Println("Error reading input. Try again.")
			reader.ReadString('\n') // очистка буфера ввода от неверных данных
			continue
		}
		break
	}
	reader.ReadString('\n') // очистка буфера от \n
	return result
}

func Print_menu() {
	fmt.Print("\n")
	fmt.Printf("MENU:\n")
	fmt.Printf("0. Exit\n")
	fmt.Printf("1. Insert vertex\n")
	fmt.Printf("2. Insert edge\n")
	fmt.Printf("3. Remove vertex\n")
	fmt.Printf("4. Remove edge\n")
	fmt.Printf("5. Change vertex\n")
	fmt.Printf("6. Change edge\n")
	fmt.Printf("7. Output graph as adjacency list\n")
	fmt.Printf("8. Graphical output\n")
	fmt.Printf("9. Graph traversal:BFS\n")
	fmt.Printf("10. Find the shortest path between two vertices of a graph\n")
	fmt.Printf("11. Special operation: partitioning into connected components\n")
}

func main() {
	/*создаем граф*/
	graph := Create_graph()

	var flag int = 1
	for flag == 1 {
		Print_menu()
		reader := bufio.NewReader(os.Stdin)
		choice := Read_integer(reader, "Your choice: ")
		switch choice {
		default:
			fmt.Printf("Error: Invalid input. Try again. Enter number from 0 to 10\n")
			continue
		case 0: // ВЫХОД ИЗ ПРОГРАММЫ
			flag = 0
		case 1:
			D1_Insert_Vertex(graph)
		case 2:
			D2_Insert_Edge(graph)
		case 4:
			D4_Remove_Edge(graph)
		case 7:
			D7_Output_as_adjacency_list(graph)
		}
	}
}
