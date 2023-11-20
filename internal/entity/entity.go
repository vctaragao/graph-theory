package entity

import "fmt"

type Vertice int

type Edge struct {
	Src    Vertice
	Dest   Vertice
	Weight int
}

func NewEdge(src, dest Vertice, w int) Edge {
	return Edge{
		Src:    src,
		Dest:   dest,
		Weight: w,
	}
}

func (e Edge) String() string {
	return fmt.Sprintf("[Src: %v, Dest: %v, Weight: %v]", e.Src, e.Dest, e.Weight)
}
