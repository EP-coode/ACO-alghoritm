package graph

type UndirectedNoLoopGraph[E any] struct {
	vericies *[]int
	edges    [][]E
}

func NewGraph[V any, E any](verticies []int, initialEdgeValue E) *UndirectedNoLoopGraph[E] {
	edges := make([][]E, len(verticies)-1)
	for i := range edges {
		edges[i] = make([]E, len(verticies)-i)
		for j := range edges[i] {
			edges[i][j] = initialEdgeValue
		}
	}

	return &UndirectedNoLoopGraph[E]{
		vericies: &verticies,
		edges:    edges,
	}
}

func (g *UndirectedNoLoopGraph[E]) SetEdge(x, y int, newEdge E) bool {
	if y == x {
		return false
	}

	// make sure y is bigger than x
	if x > y {
		tmp := x
		x = y
		y = tmp
	}

	if y > len(g.edges) || y < 0 {
		return false
	}

	if x >= len(g.edges) || y < 0 {
		return false
	}

	// add ofset becaouse diagonal is skipped
	(g.edges)[x][y-1-x] = newEdge
	return true
}

func (g *UndirectedNoLoopGraph[E]) GetEdge(x, y int) (*E, bool) {
	if y == x {
		return nil, false
	}

	// make sure y is bigger than x
	if x > y {
		tmp := x
		x = y
		y = tmp
	}

	if y > len(g.edges) || y < 0 {
		return nil, false
	}

	if x >= len(g.edges) || y < 0 {
		return nil, false
	}
	// add ofset becaouse diagonal is skipped
	return &(g.edges)[x][y-1-x], true
}
