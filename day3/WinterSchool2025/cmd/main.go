package main

import (
	"fmt"
	"network/pkg/graph"
)

func main() {
	// Создаем неориентированный граф
	graph := graph.NewGraph()

	// Добавляем вершины и ребра
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 6)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)

	// Выводим граф
	fmt.Println("Граф (список смежности):")
	graph.Print()

	fmt.Print("DFS: ", graph.DFS(2))
	fmt.Print("\n")
	fmt.Print("BFS: ", graph.BFS(2))

	graph.HasEdge(1, 4)
}
