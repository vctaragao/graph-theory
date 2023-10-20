package main

import (
	"fmt"

	"github.com/vctaragao/graphs/internal/entity"
)

func main() {
	adjMatrix := entity.NewAdjMatrix("adjacency_matrix.csv")
	fmt.Println(adjMatrix)
}
