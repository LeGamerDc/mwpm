package mwpm

import "slices"

type edge struct {
	x, y int64
	w    float64
}

type FilterSolver struct {
	n      int
	fp     float64
	count  []int
	edges  []edge
	solver SolverI
}

func NewFilterSolver(fp float64, s SolverI) *FilterSolver {
	return &FilterSolver{
		n:      s.N(),
		fp:     fp,
		count:  make([]int, s.N()),
		edges:  make([]edge, 0, 1000),
		solver: s,
	}
}

func (f *FilterSolver) AddEdge(x, y int64, w float64) {
	f.edges = append(f.edges, edge{x, y, w})
	f.count[x]++
	f.count[y]++
}

func (f *FilterSolver) N() int {
	return f.n
}

func (f *FilterSolver) Solve(cb func(i int64)) ([][2]int64, bool) {
	slices.SortFunc(f.edges, func(i, j edge) int {
		return int(j.w - i.w)
	})
	after := make([]edge, 0, len(f.edges))
	maxEdge := max(20, (int(float64(f.n)*f.fp)+1)/2*2)
	for _, e := range f.edges {
		if f.count[e.x] > maxEdge && f.count[e.y] > maxEdge {
			f.count[e.x]--
			f.count[e.y]--
		} else {
			after = append(after, e)
		}
	}

	wg := NewWeightedGraph(int64(f.n))
	for _, e := range after {
		wg.AddEdge(int64(e.x), int64(e.y), e.w)
	}
	return Run(wg, cb)
}
