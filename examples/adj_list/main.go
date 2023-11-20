package main

import (
	"fmt"

	"github.com/vctaragao/graphs/internal/entity"
)

func main() {
	adjList := entity.NewAdjacencyList("adjacency_matrix.csv")
	adjList.Print()
	adjList.RemoveEdge(entity.NewEdge(0, 1, 3))

	if err := adjList.RemoveEdge(entity.NewEdge(0, 2, 6)); err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	adjList.Print()
}
