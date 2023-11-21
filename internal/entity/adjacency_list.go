package entity

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/vctaragao/graphs/pkg/helpers"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

var (
	ErrInvalidVerticef = "invalid vertice: %v"
	ErrInvalidVertice  = errors.New("invalid vertice")

	ErrInvalidEdge         = errors.New("invalid edge")
	ErrVerticeAlredyExists = errors.New("vertice already exists")
)

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
func (list AdjacencyList) InsertEdge(src, dest Vertice, wheight int) error {
	srcAdjs, exists := list[src]
	if !exists {
		return fmt.Errorf(ErrInvalidVerticef, src)
	}

	destAdjs, exists := list[dest]
	if !exists {
		return fmt.Errorf(ErrInvalidVerticef, dest)
	}

	list[src] = append(srcAdjs, NewEdge(src, dest, wheight))
	list[dest] = append(destAdjs, NewEdge(dest, src, wheight))

	return nil
}

// InsertVertice: insert a new vertice in to the graph
func (list AdjacencyList) InsertVertice(v Vertice) error {
	if _, exists := list[v]; exists {
		return ErrVerticeAlredyExists
	}

	list[v] = []Edge{}
	return nil
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
func (list AdjacencyList) RemoveVertice(v Vertice) {
	edges, exists := list[v]
	if !exists {
		return
	}

	for _, e := range edges {
		for i, destEdge := range list[e.Dest] {
			if destEdge.Dest == v {
				list[e.Dest] = append(list[e.Dest][:i], list[e.Dest][i+1:]...)
			}
		}
	}

	delete(list, v)
}

func (list AdjacencyList) Print() {
	vertices := maps.Keys(list)
	slices.Sort(vertices)

	for _, v := range vertices {
		fmt.Printf("Vertice: %v\n", v)
		fmt.Printf("Adjacências: [\n")

		for _, edge := range list[v] {
			fmt.Println(edge)
		}

		fmt.Printf("]\n")
	}
}
