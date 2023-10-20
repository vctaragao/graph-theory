package entity

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/vctaragao/graphs/pkg/helpers"
)

type adjacencyMatrix map[int][]int

func NewAdjMatrix(fileName string) adjacencyMatrix {
	file, err := os.Open(fileName)
	helpers.Chk(err)
	defer file.Close()

	vertice := 0
	adjMatrix := make(adjacencyMatrix)

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

			adjMatrix[vertice] = append(adjMatrix[vertice], edgeWeight)
		}

		vertice++
	}

	return adjMatrix
}
