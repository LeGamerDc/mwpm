package mwpm

import "iter"

type link struct {
	to     int64
	weight float64
}

type WeightedGraph struct {
	n  int64
	ww []float64
	nn [][]link
}

func NewWeightedGraph(n int64) *WeightedGraph {
	return &WeightedGraph{
		n:  n,
		ww: make([]float64, n*n),
		nn: make([][]link, n),
	}
}

func (w *WeightedGraph) AddEdge(u, v int64, wi float64) {
	if u != v {
		p1, p2 := u*w.n+v, v*w.n+u
		if w.ww[p1] == 0 {
			w.ww[p1] = wi
			w.ww[p2] = wi
			w.nn[u] = append(w.nn[u], link{to: v, weight: wi})
			w.nn[v] = append(w.nn[v], link{to: u, weight: wi})
		}
	}
}

func (w *WeightedGraph) Connect(u int64) iter.Seq2[int64, float64] {
	return func(yield func(int64, float64) bool) {
		for _, l := range w.nn[u] {
			if !yield(l.to, l.weight) {
				return
			}
		}
	}
}

func (w *WeightedGraph) N() int {
	return int(w.n)
}

func (w *WeightedGraph) HasEdgeBetween(xid, yid int64) bool {
	return w.ww[yid*w.n+xid] > 0
}

func (w *WeightedGraph) Weight(xid, yid int64) (wi float64, ok bool) {
	wi = w.ww[yid*w.n+xid]
	return wi, wi > 0
}
