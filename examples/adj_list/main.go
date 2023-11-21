package main

import (
	"fmt"

	"github.com/vctaragao/graphs/internal/entity"
)

func main() {
	adjList := entity.NewAdjacencyList("adjacency_matrix.csv")
	adjList.Print()
	adjList.InsertVertice(entity.Vertice(12))
	if err := adjList.InsertEdge(12, 1, 10); err != nil {
		fmt.Printf("err: %v\n", err)
	}
	adjList.Print()
	adjList.RemoveVertice(entity.Vertice(12))
	adjList.Print()
}
