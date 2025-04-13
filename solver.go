package mwpm

type SolverI interface {
	AddEdge(x, y int64, w float64)
	Solve(cb func(int64)) ([][2]int64, []int64)
	N() int
}

type BaseSolver struct {
	wg *WeightedGraph
}

func NewBaseSolver(n int64) *BaseSolver {
	return &BaseSolver{
		wg: NewWeightedGraph(n),
	}
}

func (b *BaseSolver) N() int {
	return b.wg.N()
}

func (b *BaseSolver) AddEdge(x, y int64, w float64) {
	b.wg.AddEdge(x, y, w)
}

func (b *BaseSolver) Solve(cb func(int64)) ([][2]int64, []int64) {
	return run(b.wg, cb)
}

type MaxSolver struct {
	s   SolverI
	max float64
}

func NewMaxSolver(max float64, s SolverI) *MaxSolver {
	return &MaxSolver{
		s:   s,
		max: max + 1.0,
	}
}

func (m *MaxSolver) N() int {
	return m.s.N()
}

func (m *MaxSolver) AddEdge(x, y int64, w float64) {
	m.s.AddEdge(x, y, m.max-w)
}

func (m *MaxSolver) Solve(cb func(int64)) ([][2]int64, []int64) {
	return m.s.Solve(cb)
}

type xy struct {
	x, y int64
}

type EvaluateSolver struct {
	s SolverI
	e map[xy]float64

	weighted float64
}

func NewEvaluateSolver(s SolverI) *EvaluateSolver {
	return &EvaluateSolver{
		s: s,
		e: make(map[xy]float64),
	}
}

func (e *EvaluateSolver) N() int {
	return e.s.N()
}

func (e *EvaluateSolver) AddEdge(x, y int64, w float64) {
	e.s.AddEdge(x, y, w)
	if x > y {
		x, y = y, x
	}
	e.e[xy{x, y}] = w
}

func (e *EvaluateSolver) Solve(cb func(int64)) (pair [][2]int64, rest []int64) {
	pair, rest = e.s.Solve(cb)
	if len(rest) > 0 {
		return
	}
	for i := range pair {
		x, y := pair[i][0], pair[i][1]
		if x > y {
			x, y = y, x
		}
		e.weighted += e.e[xy{x, y}]
	}
	return
}

func (e *EvaluateSolver) Weighted() float64 {
	return e.weighted
}
