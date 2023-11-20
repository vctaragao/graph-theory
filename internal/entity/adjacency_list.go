package entity

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/vctaragao/graphs/pkg/helpers"
)

var ErrInvalidEdge = errors.New("invalid edge")
var ErrInvalidVertice = errors.New("invalid vertice")

type AdjacencyList map[Vertice][]Edge

func NewAdjacencyList(fileName string) AdjacencyList {
	file, err := os.Open(fileName)
	helpers.Chk(err)
	defer file.Close()

	adjList := make(AdjacencyList, 0)

	var vertice Vertice = 0
	csvReader := csv.NewReader(file)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		helpers.Chk(err)

		for dest := 0; dest < len(record); dest++ {
			edgeWeight, err := strconv.Atoi(record[dest])
			helpers.Chk(err)

			if edgeWeight < 1 {
				continue
			}

			edge := Edge{
				Src:    vertice,
				Dest:   Vertice(dest),
				Weight: edgeWeight,
			}

			adjList[vertice] = append(adjList[vertice], edge)
		}

		vertice++
	}

	return adjList
}

// Type: should return what type of graph is this
func (list AdjacencyList) Type() {

}

// HasAdjacency: should check if two given vertices are adjacency to each other
func (list AdjacencyList) HasAdjacency(v, dest Vertice) (bool, error) {
	adjs, exists := list[v]
	if !exists {
		return false, ErrInvalidVertice
	}

	for _, edge := range adjs {
		if edge.Dest == dest {
			return true, nil
		}
	}

	return false, nil
}

// Density: return the density value of the graph
func (list AdjacencyList) Density() {

}

// InsertEdge: insert a new edge in to the graph
func (list AdjacencyList) InsertEdge() {

}

// InsertVertice: insert a new vertice in to the graph
func (list AdjacencyList) InsertVertice() {

}

// RemoveEdge: remove an existing edge from the graph
func (list AdjacencyList) RemoveEdge(edge Edge) error {
	err := ErrInvalidEdge
	for v, edges := range list {
		for i, e := range edges {
			if e == edge {
				list[v] = append(edges[:i], edges[i+1:]...)
				err = nil
				break
			}

		}
	}

	return err
}

// RemoveVertice: remove an existing vertice from the graph
func (list AdjacencyList) RemoveVertice() {

}

func (list AdjacencyList) Print() {
	for i := 0; i < len(list); i++ {
		v := Vertice(i)
		fmt.Printf("Vertice: %v\n", v)
		fmt.Printf("Adjacências: [\n")

		for _, edge := range list[v] {
			fmt.Println(edge)
		}

		fmt.Printf("]\n")
	}
}
