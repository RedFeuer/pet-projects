package dialog

import (
	"ComputerNetwork/internal"
	"ComputerNetwork/logic"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func D1_Insert_Vertex(graph *internal.Graph) {
	reader := bufio.NewReader(os.Stdin)
	comp := Read_non_empty_string(reader, "Enter unique computer name: ")
	fmt.Printf("Enter number of port for computer %s", comp)
	port := Read_integer(reader, ": ")
	new_vertex := internal.Create_vertex(comp, port)
	new_node := internal.Initialize_node(new_vertex)
	graph.Table[comp] = new_node
}

func D2_Insert_Edge(graph *internal.Graph) {
	reader := bufio.NewReader(os.Stdin)
	comp_src := Read_non_empty_string(reader, "Enter computer name for source computer: ")
	comp_dst := Read_non_empty_string(reader, "Enter computer name for destination computer: ")

	ports_count := Read_integer(reader, "Enter number of ports: ")
	ports := make([]uint, ports_count)
	for i := 0; i < int(ports_count); i++ {
		ports[i] = Read_integer(reader, "Enter port: ")
	}

	termination_status := logic.Insert_Edge_logic(graph, comp_src, comp_dst, ports_count, ports)
	switch termination_status {
	default:
		break
	case 1:
		fmt.Printf("Error: There no computer with name %s in computer network", comp_src)
	case 2:
		fmt.Printf("Error: There no computer with name %s in computer network", comp_dst)
	case 3:
		fmt.Println("Error: can't connect vertex to itself")
	case 4:
		fmt.Println("Error: can't connect vertices that have been already connected")
	}
}

func D3_Remove_Vertex(graph *internal.Graph) {
	reader := bufio.NewReader(os.Stdin)
	comp_remove := Read_non_empty_string(reader, "Enter computer name: ")
	termination_status := logic.Remove_Vertex_Logic(graph, comp_remove)
	switch termination_status {
	case 1:
		fmt.Println("Error: There no such computer in computer network")
	}
}

func D4_Remove_Edge(graph *internal.Graph) {
	reader := bufio.NewReader(os.Stdin)
	comp_src := Read_non_empty_string(reader, "Enter computer name for source computer: ")
	comp_dst := Read_non_empty_string(reader, "Enter computer name for destination computer: ")
	termination_status := logic.Remove_Edge_Logic(graph, comp_src, comp_dst)
	switch termination_status {
	case 1:
		fmt.Printf("Error: There no computer with name %s\n", comp_src)
	case 2:
		fmt.Printf("Error: There no computer with name %s\n", comp_dst)
	case 3:
		fmt.Printf("Error: There is no edge between %s and %s \n", comp_src, comp_dst)
	}
}

func D6_Change_Edge(graph *internal.Graph) {
	reader := bufio.NewReader(os.Stdin)
	comp_src := Read_non_empty_string(reader, "Enter computer name for source computer: ")
	comp_dst := Read_non_empty_string(reader, "Enter computer name for destination computer: ")
	new_ports_count := Read_integer(reader, "Enter number of ports: ")

	new_ports := make([]uint, new_ports_count)
	for i := 0; i < int(new_ports_count); i++ {
		new_ports[i] = Read_integer(reader, "Enter port: ")
	}

	termination_status := logic.Change_Edge_Logic(graph, comp_src, comp_dst, new_ports_count, new_ports)
	switch termination_status {
	case 1:
		fmt.Printf("Error: There no edge between %s and %s \n", comp_src, comp_dst)
	}
}

func D7_Output_as_adjacency_list(graph *internal.Graph) {
	for comp, node := range graph.Table {
		fmt.Printf("Computer name: %s Port: %d  ", comp, node.Vertex.Port)
		if node.Adjacent == nil {
			fmt.Printf("\n")
			continue
		}
		for elem := node.Adjacent.Front(); elem != nil; elem = elem.Next() {
			adjacent_vertex := elem.Value.(*internal.AdjacentVertex)
			fmt.Printf(" ->  %s", adjacent_vertex.Vertex.Comp)
		}
		fmt.Printf("\n")
	}
}

func D8_Graphviz_Output(graph *internal.Graph) {
	reader := bufio.NewReader(os.Stdin)
	filename := Read_non_empty_string(reader, "Enter filename: ")
	termination_status := logic.Create_Dot_File(graph, filename)
	switch termination_status {
	case 1:
		fmt.Println("Error")
	}
}

func D9_BFS_dialog(graph *internal.Graph) {
	reader := bufio.NewReader(os.Stdin)
	source := Read_non_empty_string(reader, "Enter computer name to Breadth-First Search: ")
	logic.BFS(graph, source)
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
	fmt.Printf("7. Output internal as adjacency list\n")
	fmt.Printf("8. Graphical output\n")
	fmt.Printf("9. Graph traversal:BFS\n")
	fmt.Printf("10. Find the shortest path between two vertices of a internal\n")
	fmt.Printf("11. Special operation: partitioning into connected components\n")
}
