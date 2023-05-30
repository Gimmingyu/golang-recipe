package main

import "fmt"

type Graph struct {
	vertices map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.vertices[u] = append(g.vertices[u], v)
}

func DFS(graph *Graph, start int, visited map[int]bool) {
	visited[start] = true
	fmt.Println(start)

	for _, neighbor := range graph.vertices[start] {
		if !visited[neighbor] {
			DFS(graph, neighbor, visited)
		}
	}
}

func main() {
	graph := NewGraph()

	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 6)

	visited := make(map[int]bool)
	DFS(graph, 1, visited)
}