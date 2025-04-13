package mwpm

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMWPM(t *testing.T) {
	adj := [][]int{
		{0, 1, 1, 0, 0, 0},
		{1, 0, 1, 1, 0, 0},
		{1, 1, 0, 1, 1, 0},
		{0, 1, 1, 0, 1, 1},
		{0, 0, 1, 1, 0, 1},
		{0, 0, 0, 1, 1, 0},
	}
	wg := NewWeightedGraph(6)
	for i := 0; i < len(adj); i++ {
		for j := i + 1; j < len(adj[i]); j++ {
			if adj[i][j] != 0 {
				wg.AddEdge(int64(i), int64(j), float64(adj[i][j]))
			}
		}
	}
	fmt.Println("----------------NEW----------------")
	pair, _ := run(wg, func(i int64) {})
	fmt.Println(pair)
}

func TestMWPM2(t *testing.T) {
	wg := NewWeightedGraph(4)
	wg.AddEdge(0, 1, 2)
	wg.AddEdge(2, 3, 3)
	wg.AddEdge(0, 2, 1)
	wg.AddEdge(1, 3, 2)
	pair, _ := run(wg, func(i int64) {})
	fmt.Println(pair)
}

func BenchmarkSolver(b *testing.B) {
	for x := 0; x < b.N; x++ {
		n := 1000
		g := NewEvaluateSolver(NewMaxSolver(101, NewFilterSolver(0, NewBaseSolver(int64(n)))))
		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				g.AddEdge(int64(i), int64(j), float64(rand.Intn(99)+1))
			}
		}
		_, ok := g.Solve(func(i int64) {})
		fmt.Println(ok, g.Weighted())
	}
}
