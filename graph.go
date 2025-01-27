package main

import (
	"fmt"
)

// Graph представляет собой граф с использованием списка смежности
type Graph struct {
	vertices map[int][]int
	directed bool // Показывает, ориентированный ли граф
}

// NewGraph создает новый граф (ориентированный или неориентированный)
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[int][]int),
	}
}

// AddVertex добавляет вершину в граф
func (g *Graph) AddVertex(vertex int) {
	if _, exists := g.vertices[vertex]; !exists {
		g.vertices[vertex] = []int{}
	}
}

// AddEdge добавляет ребро между двумя вершинами
func (g *Graph) AddEdge(u, v int) {
	// Добавляем вершины, если их еще нет
	g.AddVertex(u)
	g.AddVertex(v)

	// Добавляем ребро
	g.vertices[u] = append(g.vertices[u], v)
	g.vertices[v] = append(g.vertices[v], u)

}

// Print выводит граф в виде списка смежности
func (g *Graph) Print() {
	for vertex, edges := range g.vertices {
		fmt.Printf("%d -> %v\n", vertex, edges)
	}
}

func HasEdge(g *Graph, u, v int) bool {
	for i := 0; i < len(g.vertices); i++ {
		for j := 0; j < len(g.vertices[i]); j++ {
			fmt.Printf("%d ", g.vertices[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("%d \n", len(g.vertices))
	buf := false
	for j := 0; j < len(g.vertices[u]); j++ {
		fmt.Printf("%d ", g.vertices[u][j])
		if g.vertices[u][j] == v {
			buf = true
		}

	}

	fmt.Printf("friend? : %d", buf)

	return buf

}

func main() {
	// Создаем неориентированный граф
	graph := NewGraph()

	// Добавляем вершины и ребра
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)

	// Выводим граф
	fmt.Println("Граф (список смежности):")
	graph.Print()

	HasEdge(graph, 1, 4)
}
