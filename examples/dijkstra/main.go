package main

import (
	"fmt"

	"github.com/vctaragao/graphs/internal"
	"github.com/vctaragao/graphs/internal/entity"
)

func main() {
	adjMatrix := entity.NewAdjMatrix("adjacency_matrix.csv")
	minPath := internal.Dijsktra(adjMatrix, 0, 5)
	fmt.Println("Min Path:", minPath)
}
