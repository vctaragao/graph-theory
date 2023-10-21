package entity

import (
	"encoding/csv"
	"io"
	"math"
	"os"
	"strconv"

	"github.com/vctaragao/graphs/pkg/helpers"
)

type AdjacencyMatrix map[int][]int

func NewAdjMatrix(fileName string) AdjacencyMatrix {
	file, err := os.Open(fileName)
	helpers.Chk(err)
	defer file.Close()

	vertice := 0
	adjMatrix := make(AdjacencyMatrix)

	csvReader := csv.NewReader(file)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		helpers.Chk(err)

		for i := 0; i < len(record); i++ {
			edgeWeight, err := strconv.Atoi(record[i])
			helpers.Chk(err)

			if edgeWeight == -1 {
				edgeWeight = math.MaxInt32
			}

			adjMatrix[vertice] = append(adjMatrix[vertice], edgeWeight)
		}

		vertice++
	}

	return adjMatrix
}
