package graph

import (
	"fmt"
	"network/pkg/queue"
	"network/pkg/stack"
)

// Graph представляет собой граф с использованием списка смежности
type Graph struct {
	vertices map[int][]int
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

func (g *Graph) HasEdge(u, v int) bool {
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

	fmt.Printf("friend? : %t", buf)

	return buf

}

// BFS реализует алгоритм поиска в ширину
func (g *Graph) BFS(start int) []int {
	// Создаем очередь для хранения вершин, которые нужно посетить
	q := &queue.Queue{}

	q.Enqueue(start)

	// Создаем мапу для отслеживания посещенных вершин
	visited := make(map[int]bool)
	visited[start] = true

	// Создаем список для хранения порядка обхода вершин
	result := []int{}

	// Пока очередь не пуста, продолжаем обход
	for len(q.Data) > 0 {
		// Извлекаем вершину из очереди
		current, err := q.Dequeue()
		if err != nil {
			break
		}

		// Добавляем текущую вершину в результат
		result = append(result, current)

		// Перебираем всех соседей текущей вершины
		for _, neighbor := range g.vertices[current] {
			// Если соседняя вершина еще не посещена, добавляем ее в очередь
			if !visited[neighbor] {
				visited[neighbor] = true
				q.Enqueue(neighbor)
			}
		}
	}

	return result
}

// DFS реализует алгоритм поиска в глубину (итеративный)
func (g *Graph) DFS(start int) []int {
	// Создаем стек для хранения вершин, которые нужно посетить
	stack := &stack.Stack{}
	stack.Push(start)

	// Создаем мапу для отслеживания посещенных вершин
	visited := make(map[int]bool)
	visited[start] = true

	// Создаем список для хранения порядка обхода вершин
	result := []int{}

	// Пока стек не пуст, продолжаем обход
	for !stack.IsEmpty() {
		// Извлекаем вершину из стека
		current, err := stack.Pop()
		if err != nil {
			break
		}

		// Добавляем текущую вершину в результат
		result = append(result, current)

		// Перебираем всех соседей текущей вершины
		for _, neighbor := range g.vertices[current] {
			// Если соседняя вершина еще не посещена, добавляем ее в стек
			if !visited[neighbor] {
				visited[neighbor] = true
				stack.Push(neighbor)
			}
		}
	}

	return result
}
