package main

import (
	"ComputerNetwork/dialog"
	"ComputerNetwork/internal"
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*создаем граф*/
	graph := internal.Create_graph()

	var flag int = 1
	for flag == 1 {
		dialog.Print_menu()
		reader := bufio.NewReader(os.Stdin)
		choice := dialog.Read_integer(reader, "Your choice: ")
		switch choice {
		default:
			fmt.Printf("Error: Invalid input. Try again. Enter number from 0 to 10\n")
			continue
		case 0: // ВЫХОД ИЗ ПРОГРАММЫ
			flag = 0
		case 1:
			dialog.D1_Insert_Vertex(graph)
		case 2:
			dialog.D2_Insert_Edge(graph)
		case 3:
			dialog.D3_Remove_Vertex(graph)
		case 4:
			dialog.D4_Remove_Edge(graph)
		case 6:
			dialog.D6_Change_Edge(graph)
		case 7:
			dialog.D7_Output_as_adjacency_list(graph)
		case 8:
			dialog.D8_Graphviz_Output(graph)
		}
	}
}
