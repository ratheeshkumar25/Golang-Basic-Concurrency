package main

import (
	"fmt"
)

type Graph struct {
	Vertex map[int][]int
}

func NewGraph() *Graph {
	return &Graph{Vertex: make(map[int][]int)}
}

func (g *Graph) Insert(v1, v2 int) {
	g.Vertex[v1] = append(g.Vertex[v1], v2)
	g.Vertex[v2] = append(g.Vertex[v2], v1)
}

func (g *Graph) RemoveEdge(v1, v2 int) {
	arr := g.Vertex[v1]
	for i := 0; i < len(arr); i++ {
		if arr[i] == v2 {
			g.Vertex[v1] = append(arr[:i], arr[i+1:]...)
			break
		}
	}
	arr = g.Vertex[v2]
	for i := 0; i < len(arr); i++ {
		if arr[i] == v1 {
			g.Vertex[v2] = append(arr[:i], arr[i+1:]...)
			break
		}
	}

}

func (g *Graph) RemoveVertex(v int) {
	arr := g.Vertex[v]
	for i := 0; i < len(arr); i++ {
		edges := g.Vertex[arr[i]]
		for j := 0; j < len(edges); j++ {
			if edges[j] == v {
				g.Vertex[arr[i]] = append(arr[:j], arr[j+1:]...)
				break
			}
		}
	}
	delete(g.Vertex, v)
}

func (g *Graph) BFS(val int) {
	arr := []int{val}
	visited := make(map[int]bool)

	visited[val] = true

	for len(arr) > 0 {
		v := arr[0]
		fmt.Printf("%d", v)
		arr = arr[1:]

		for _, e := range g.Vertex[v] {
			if !visited[e] {
				arr = append(arr, e)
				visited[e] = true
			}
		}
	}
}

func (g *Graph) ShortestPath(start, end int) ([]int, int) {
	queue := []int{start}
	visited := make(map[int]bool)
	parent := make(map[int]int)

	visited[start] = true
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == end {
			break
		}
		for _, neighbor := range g.Vertex[current] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
				parent[neighbor] = current
			}
		}
	}
	if !visited[end] {
		return nil, -1
	}

	path := []int{}
	current := end
	for current != start {
		path = append([]int{current}, path...)
		current = parent[current]
	}
	path = append([]int{current}, path...)
	return path, len(path) - 1
}

func (g *Graph) Display() {
	fmt.Println(g.Vertex)
}

func main() {
	g := NewGraph()
	g.Insert(1, 2)
	g.Insert(1, 3)
	g.Insert(2, 3)
	g.Insert(2, 4)
	g.Insert(3, 5)
	g.Insert(4, 5)
	g.Insert(4, 6)
	g.Insert(5, 6)

	g.BFS(6)
	g.Display()

	path, distance := g.ShortestPath(1, 6)
	if distance == -1 {
		fmt.Println("NO PATH FOUND")
	} else {
		fmt.Printf("Shortest path from 1 to 4 %v\n", path)
		fmt.Printf("Distance %v\n", distance)
	}

}
