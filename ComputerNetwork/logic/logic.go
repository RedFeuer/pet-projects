package logic

import (
	"ComputerNetwork/internal"
	"container/list"
	"fmt"
	"math"
	"os"
)

func Insert_Edge_logic(graph *internal.Graph, comp_src string, comp_dst string, ports_count uint, ports []uint) int {
	/*проверка: одной из вершин нет*/
	if graph.Table[comp_src] == nil {
		return 1
	}
	if graph.Table[comp_dst] == nil {
		return 2
	}
	/*проверка: не соединяем одинаковые вершины*/
	if graph.Table[comp_src] == graph.Table[comp_dst] {
		return 3
	}
	/*проверка: между соединяемыми вершинами не существует ребра*/
	if internal.Find_adjacent_vertex(graph, comp_src, comp_dst) != nil || internal.Find_adjacent_vertex(graph, comp_dst, comp_src) != nil {
		return 4
	}

	/*соединяем основную вершину со смежной: src -> dst*/
	if graph.Table[comp_src].Adjacent == nil {
		graph.Table[comp_src].Adjacent = list.New()
	}
	new_adjacent_for_src := &internal.AdjacentVertex{}
	new_adjacent_for_src.Edge = &internal.Edge{ports, ports_count}
	new_adjacent_for_src.Vertex = graph.Table[comp_dst].Vertex
	graph.Table[comp_src].Adjacent.PushBack(new_adjacent_for_src)

	/*обратное соединение: dst -> src*/
	if graph.Table[comp_dst].Adjacent == nil {
		graph.Table[comp_dst].Adjacent = list.New()
	}
	new_adjacent_for_dst := &internal.AdjacentVertex{}
	new_adjacent_for_dst.Edge = &internal.Edge{ports, ports_count}
	new_adjacent_for_dst.Vertex = graph.Table[comp_src].Vertex
	graph.Table[comp_dst].Adjacent.PushBack(new_adjacent_for_dst)

	return 0
}

func Remove_Edge_Logic(graph *internal.Graph, comp_src string, comp_dst string) int {
	/*ДЕЛАТЬ ПРОВЕРКИ*/
	if graph.Table[comp_src] == nil {
		return 1
	}
	if graph.Table[comp_dst] == nil {
		return 2
	}

	elem_remove_src := internal.Find_adjacent_vertex(graph, comp_src, comp_dst)
	if elem_remove_src == nil {
		return 3
	}
	graph.Table[comp_src].Adjacent.Remove(elem_remove_src)

	elem_remove_dst := internal.Find_adjacent_vertex(graph, comp_dst, comp_src)
	if elem_remove_dst == nil {
		return 3
	}
	graph.Table[comp_dst].Adjacent.Remove(elem_remove_dst)
	return 0
}

func Remove_Vertex_Logic(graph *internal.Graph, comp_remove string) int {
	if graph.Table[comp_remove] == nil {
		return 1
	}

	/*посредством Find_adjacent_vertex нужно проверить все связи с другими вершинами и удалить их*/
	internal.Remove_Adjacent_Vertex(graph, comp_remove)
	/*удалить саму вершину из хэш-таблицы*/
	delete(graph.Table, comp_remove)
	return 0
}

func Change_Edge_Logic(graph *internal.Graph, comp_src string, comp_dst string, new_ports_count uint, new_ports []uint) int {
	new_edge := &internal.Edge{new_ports, new_ports_count}

	elem_src := internal.Find_adjacent_vertex(graph, comp_src, comp_dst)
	if elem_src == nil {
		return 1
	}
	adjacent_vertex_src := elem_src.Value.(*internal.AdjacentVertex)
	if adjacent_vertex_src.Vertex.Comp == comp_dst {
		adjacent_vertex_src.Edge = new_edge
	}

	elem_dst := internal.Find_adjacent_vertex(graph, comp_dst, comp_src)
	if elem_dst == nil {
		return 1
	}
	adjacent_vertex_dst := elem_dst.Value.(*internal.AdjacentVertex)
	if adjacent_vertex_dst.Vertex.Comp == comp_src {
		adjacent_vertex_dst.Edge = new_edge
	}

	return 0
}

func Create_Dot_File(graph *internal.Graph, filename string) int {
	file, err := os.Create(filename)
	if err != nil {
		return 1
		//return err
	}
	defer file.Close()

	file.WriteString("graph G {\n")

	/*хэш-таблица(map) уже записанных в файл вершин*/
	vertices := make(map[string]bool)

	/*хэш-таблица(map) уже записанных в файл ребер*/
	edges := make(map[string]bool)

	for _, node := range graph.Table {
		/*записываем основную вершину из хэш-таблицы*/
		if vertices[node.Vertex.Comp] == false {
			file.WriteString(fmt.Sprintf("  \"%s\" [label=\"%s : %d\"];\n", node.Vertex.Comp, node.Vertex.Comp, node.Vertex.Port))
			vertices[node.Vertex.Comp] = true
		}

		/*записываем смежные вершины из списка для основной вершины*/
		if node.Adjacent != nil {
			for elem := node.Adjacent.Front(); elem != nil; elem = elem.Next() {
				adjacent_vertex := elem.Value.(*internal.AdjacentVertex)

				// Создаём ключ для ребра в виде "vertex1-vertex2" и "vertex2-vertex1"
				edge_key1 := fmt.Sprintf("%s-%s", node.Vertex.Comp, adjacent_vertex.Vertex.Comp)
				edge_key2 := fmt.Sprintf("%s-%s", adjacent_vertex.Vertex.Comp, node.Vertex.Comp)
				if edges[edge_key1] == false && edges[edge_key2] == false {
					file.WriteString(fmt.Sprintf("  \"%s\" -- \"%s\" [label=\"%v\"];\n", node.Vertex.Comp, adjacent_vertex.Vertex.Comp, adjacent_vertex.Edge.Ports))
					edges[edge_key1] = true
					edges[edge_key2] = true
				}
			}
		}
	}
	file.WriteString("}\n")
	return 0
	//return nil
}

func BFS(graph *internal.Graph, source string, port_search uint) *internal.Node {
	/*сделать так, чтобы искало, есть ли в принципе в таблице нода с таким портом*/

	for _, node := range graph.Table {
		if node.Vertex.Comp == source {
			continue
		}
		node.Vertex.Color = internal.WHITE
		node.Vertex.Path_size = math.MaxInt
		node.Vertex.Parent = nil
	}

	graph.Table[source].Vertex.Color = internal.GRAY
	graph.Table[source].Vertex.Path_size = 0
	graph.Table[source].Vertex.Parent = nil
	queue := list.New()
	queue.PushBack(graph.Table[source])
	for queue.Len() > 0 {
		current := queue.Front() //достаем первый элемент из очереди
		current_node := current.Value.(*internal.Node)
		for adjacent := current_node.Adjacent.Front(); adjacent != nil; adjacent = adjacent.Next() {
			// adjacent_vertex := adjacent.Value.(*internal.AdjacentVertex)
			adjacent_vertex := adjacent.Value.(*internal.AdjacentVertex).Vertex
			if adjacent_vertex.Color == internal.WHITE {
				/*сделать проверку по допустимости прохода по ребру(ПРОВЕРЯТЬ ПОРТЫ)*/
				/*сделать BFS_Relax*/
				adjacent_vertex.Color = internal.GRAY
				adjacent_vertex.Path_size = current_node.Vertex.Path_size + 1
				adjacent_vertex.Parent = current_node

				/*добавляем обработанную(серую) ноду в очередь*/
				adjacent_node := graph.Table[adjacent_vertex.Comp]
				queue.PushBack(adjacent_node)
			}
		}
		current_node.Vertex.Color = internal.BLACK
		queue.Remove(current) // удаляем рассмотренный элемент из очереди
	}
	/* ИДЕЯ 1: СДЕЛАТЬ ФУНКЦИЮ, КОТОРАЯ БУДЕТ ПРОХОДИТЬ ПО ТАБЛИЦЕ И ПРОВЕРЯТЬ НА СОВПАДЕНИЕ PORT И ВОЗВРАЩАТЬ COMP С
	МИНИМАЛЬНЫМИ РАЗМЕРОМ ПУТИ PATH_SIZE*/
	comp_result := Port_selection(graph, port_search)
	return graph.Table[comp_result]
	/* ИДЕЯ 2: */
}

func Port_selection(graph *internal.Graph, port_search uint) string {
	path_size_min := math.MaxInt
	var comp_result string
	for _, node := range graph.Table {
		if node.Vertex.Port == port_search && node.Vertex.Path_size < path_size_min {
			path_size_min = node.Vertex.Path_size
			comp_result = node.Vertex.Comp
		}
	}
	return comp_result
}
