package internal

import (
	"math"

	"github.com/vctaragao/graphs/internal/entity"
	"golang.org/x/exp/slices"
)

func Dijsktra(adjMatrix entity.AdjacencyMatrix, src, dest int) []int {
	nVertices := len(adjMatrix)

	cost := make([]int, nVertices)
	route := make([]int, nVertices)
	vToBeChecked := make([]int, nVertices)

	for i := 0; i < nVertices; i++ {
		route[i] = 0
		vToBeChecked[i] = i
		cost[i] = math.MaxInt32
	}

	cost[src] = 0
	vAlredyChecked := make([]int, 0, nVertices)

	for len(vToBeChecked) != 0 {
		currI, currV := min(vToBeChecked, cost)

		vAlredyChecked = append(vAlredyChecked, currV)
		vToBeChecked = slices.Delete(vToBeChecked, currI, currI+1)

		for adjV, eWeight := range adjMatrix[currV] {
			if !slices.Contains(vAlredyChecked, adjV) && cost[currV]+eWeight < cost[adjV] {
				route[adjV] = currV
				cost[adjV] = cost[currV] + eWeight
			}
		}
	}
	return buildMinPath(route, src, dest)
}

func min(vToBeChecked, cost []int) (i, c int) {
	minIndex := 0
	minVertice := 0
	minimun := math.MaxInt32

	for i, v := range vToBeChecked {
		if cost[v] < minimun {
			minIndex = i
			minVertice = v
			minimun = cost[v]
		}
	}

	return minIndex, minVertice
}

func buildMinPath(route []int, src, dest int) []int {
	minPath := make([]int, 0, len(route))
	for v := route[dest]; v != src; v = route[v] {
		minPath = append(minPath, v)
	}
	slices.Reverse(minPath)
	minPath = append(minPath, dest)

	return append([]int{src}, minPath...)
}
