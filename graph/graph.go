package graph

type Graph[E any] interface {
	NewGraph(verticies []int, initialEdgeValue E) Graph[E]
	SetEdge(x, y int, newEdge E) bool
	GetEdge(x, y int) (*E, bool)
}
