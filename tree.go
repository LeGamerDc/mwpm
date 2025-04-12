package mwpm

type Tree struct {
	g     *WeightedGraph
	roots map[*Node]struct{}
	nodes []*Node
	tight map[*Node]*Node
}

func NewTree(wg *WeightedGraph) *Tree {
	t := &Tree{
		g:     wg,
		roots: make(map[*Node]struct{}),
		nodes: make([]*Node, wg.N()),
		tight: make(map[*Node]*Node),
	}
	for i := int64(0); i < int64(wg.N()); i++ {
		nid := i
		n := &Node{label: 1}
		n.temp = nid
		t.nodes[nid] = n
	}
	return t
}

func (t *Tree) Blossoms() []*Node {
	var nodes []*Node
	unique := make(map[*Node]struct{})
	for n := range t.roots {
		for _, m := range n.Descendents() {
			b := m.Blossom()
			if _, ok := unique[b]; !ok {
				nodes = append(nodes, b)
				unique[b] = struct{}{}
			}
		}
	}
	return nodes
}

// returns the pair of nodes that connects the cycle.
// For example, it returns [[2 1] [1 0] [0 4] [4 3] [3 2]] if the tree looks like:
//		p o  [2]
//	    /   \
// [1] o     o [3]
//     |     |
// [0] o     o [4]
//	   n --  m

func (t *Tree) MakeCycle(n, m *Node) [][2]*Node {
	ansn := n.Anscesters()
	ansm := m.Anscesters()
	var i, j int
	for i = 0; i < len(ansn); i++ {
		for j = 0; j < len(ansm); j++ {
			if ansn[i] == ansm[j] {
				goto FOUND
			}
		}
	}
FOUND:
	var cycle [][2]*Node
	for k := i; k > 0; k-- {
		u := ansn[k].PopChild(ansn[k-1])
		ub := u.Blossom()
		cycle = append(cycle, [2]*Node{ub.parent, u})
		ub.parent = nil
	}
	cycle = append(cycle, [2]*Node{n, m})
	for k := 0; k < j; k++ {
		u := ansm[k+1].PopChild(ansm[k])
		ub := u.Blossom()
		cycle = append(cycle, [2]*Node{u, ub.parent})
		ub.parent = nil
	}
	return cycle
}

// set the node n as tight within the blossom b
func (t *Tree) SetTight(s [2]*Node, b *Node) {
	// fmt.Printf("setting tight edge %d:%d\n", s[0].temp, s[1].temp)
	t.tight[s[0]], t.tight[s[1]] = s[1], s[0]
	for _, l := range s {
		for l.pp != b {
			lb := l.pp
			for i, c := range lb.cycle {
				if c[0].BlossomWithin(lb) == l.BlossomWithin(lb) {
					lb.cycle = append(lb.cycle[i:], lb.cycle[:i]...)
					break
				}
			}
			for i := 1; i < len(lb.cycle); i += 2 {
				t.SetTight(lb.cycle[i], lb)
			}
			l = l.pp
		}
	}
}

func (t *Tree) TightFrom(n *Node) *Node {
	for _, u := range n.All() {
		if v, ok := t.tight[u]; ok {
			if v.Blossom() != n {
				return v
			}
		}
	}
	panic("No tight match found")
	return nil
}

// remove the tight edge within the blossom b
func (t *Tree) RemoveTight(s [2]*Node) {
	for _, u := range s {
		if _, ok := t.tight[u]; ok {
			delete(t.tight, u)
			t.RemoveTightWithin(u.Blossom())
		}
	}
}

func (t *Tree) RemoveTightWithin(b *Node) {
	for _, c := range b.cycle {
		for _, u := range c {
			if u.BlossomWithin(b).pp == b {
				delete(t.tight, u)
			}
		}
		t.RemoveTightWithin(c[0].BlossomWithin(b))
	}
}
