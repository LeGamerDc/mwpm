package mwpm

import "slices"

type edge struct {
	x, y int64
	w    float64
}

type FilterSolver struct {
	n      int
	fe     int
	count  []int
	edges  []edge
	solver SolverI
}

func NewFilterSolver(fe int, s SolverI) *FilterSolver {
	return &FilterSolver{
		n:      s.N(),
		fe:     max(20, fe),
		count:  make([]int, s.N()),
		edges:  make([]edge, 0, 1000),
		solver: s,
	}
}

func (f *FilterSolver) AddEdge(x, y int64, w float64) {
	f.edges = append(f.edges, edge{x, y, w})
}

func (f *FilterSolver) N() int {
	return f.n
}

func (f *FilterSolver) Solve(cb func(i int64)) ([][2]int64, []int64) {
	slices.SortFunc(f.edges, func(i, j edge) int {
		return int(i.w - j.w)
	})
	after := make([]edge, 0, len(f.edges))
	for _, e := range f.edges {
		if f.count[e.x] < f.fe || f.count[e.y] < f.fe {
			after = append(after, e)
			f.count[e.x]++
			f.count[e.y]++
		}
	}
	slices.SortFunc(after, func(a, b edge) int {
		if a.x == b.x {
			return int(a.y - b.y)
		}
		return int(a.x - b.x)
	})

	wg := NewWeightedGraph(int64(f.n))
	for _, e := range after {
		wg.AddEdge(e.x, e.y, e.w)
	}
	return run(wg, cb)
}
