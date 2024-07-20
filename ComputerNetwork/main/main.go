package main

import (
	"bufio"
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
	Ports_count int    // количество портов для данного ребра
}

/*ЛИНЕЙНЫЙ ОДНОСВЯЗНЫЙ СПИСОК(LINEAR SINGLE-LINKED LIST)*/
type AdjacentVertex struct {
	Vertex *Vertex         // Указатель на данную смежную вершину
	Edge   *Edge           // Указатель на ребро, соединяющее основную и смежную к ней вершины
	Next   *AdjacentVertex // Следующая в списке смежная вершина(смежная к основной)
}

type Node struct {
	Vertex   *Vertex         // Указатель на данную основную вершину
	Adjacent *AdjacentVertex // Указатель на смежные вершины
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

func D1_Insert_Vertex(graph *Graph) {
	reader := bufio.NewReader(os.Stdin)
	comp := Read_non_empty_string(reader, "Enter unique computer name: ")
	fmt.Printf("Enter number of port for computer %s", comp)
	port := Read_integer(reader, ": ")
	new_vertex := Create_vertex(comp, port)
	new_node := Initialize_node(new_vertex)
	graph.Table[comp] = new_node
}

func D7_Output_as_adjacency_list(graph *Graph) {
	for comp, node := range graph.Table {
		fmt.Printf("Computer name: %s    Port: %d\n", comp, node.Vertex.Port)
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
		case 7:
			D7_Output_as_adjacency_list(graph)
		}
	}
}
