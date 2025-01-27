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
func NewGraph(directed bool) *Graph {
	return &Graph{
		vertices: make(map[int][]int),
		directed: directed,
	}
}

// AddVertex добавляет вершину в граф
func (g *Graph) AddVertex(vertex int) {
	if _, exists := g.vertices[vertex]; !exists {
		g.vertices[vertex] = []int{}
	}
}

// AddEdge добавляет ребро между двумя вершинами
func (g *Graph) AddEdge(from, to int) {
	// Добавляем вершины, если их еще нет
	g.AddVertex(from)
	g.AddVertex(to)

	// Добавляем ребро from -> to
	g.vertices[from] = append(g.vertices[from], to)

	// Если граф неориентированный, добавляем обратное ребро to -> from
	if !g.directed {
		g.vertices[to] = append(g.vertices[to], from)
	}
}

// Print выводит граф в виде списка смежности
func (g *Graph) Print() {
	for vertex, edges := range g.vertices {
		fmt.Printf("%d -> %v\n", vertex, edges)
	}
}

// BFS выполняет обход графа в ширину (Breadth-First Search)
func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := []int{start}

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		if !visited[vertex] {
			fmt.Printf("%d ", vertex)
			visited[vertex] = true

			for _, neighbor := range g.vertices[vertex] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
				}
			}
		}
	}
	fmt.Println()
}

// DFS выполняет обход графа в глубину (Depth-First Search)
func (g *Graph) DFS(start int) {
	visited := make(map[int]bool)
	g.dfsRecursive(start, visited)
	fmt.Println()
}

// dfsRecursive — вспомогательная функция для рекурсивного DFS
func (g *Graph) dfsRecursive(vertex int, visited map[int]bool) {
	if !visited[vertex] {
		fmt.Printf("%d ", vertex)
		visited[vertex] = true

		for _, neighbor := range g.vertices[vertex] {
			g.dfsRecursive(neighbor, visited)
		}
	}
}

func main() {
	// Создаем неориентированный граф
	graph := NewGraph(false)

	// Добавляем вершины и ребра
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)

	// Выводим граф
	fmt.Println("Граф (список смежности):")
	graph.Print()

	// Обход в ширину (BFS)
	fmt.Println("\nBFS, начиная с вершины 1:")
	graph.BFS(1)

	// Обход в глубину (DFS)
	fmt.Println("\nDFS, начиная с вершины 1:")
	graph.DFS(1)
}
