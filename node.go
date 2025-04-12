package mwpm

type Node struct {
	label    int
	parent   *Node      // directs to non-blossom node
	children []*Node    // directs to non-blossom node
	cycle    [][2]*Node // cyclic pair of nodes (start, end)
	temp     int64

	//blossom *Node // immediate blossom that contains this node (nil if the node is the outermost blossom)
	//dval    float64

	// for lct, maintain a blossom forest
	p, pp *Node    // blossom parent
	ch    [2]*Node // child left, right
	z, sz float64  // node's z value and blossom tree sum z value
}

// returns the outermost blossom
func (n *Node) Blossom() *Node {
	return n.findRoot()
	//if n.blossom == nil {
	//	return n
	//}
	//return n.blossom.Blossom()
}

func (n *Node) Root() *Node {
	n = n.Blossom()
	for n.parent != nil {
		n = n.parent.Blossom()
	}
	return n
}

func (n *Node) Anscesters() []*Node {
	n = n.Blossom()
	chain := []*Node{n}
	for n.parent != nil {
		chain = append(chain, n.parent.Blossom())
		n = n.parent.Blossom()
	}
	return chain
}

func (n *Node) Anscestary() [][2]*Node {
	n = n.Blossom()
	cycle := [][2]*Node{}
	for n.parent != nil {
		for _, u := range n.parent.Blossom().children {
			if u.Blossom() == n {
				cycle = append(cycle, [2]*Node{u, n.parent})
				break
			}
		}
		n = n.parent.Blossom()
	}
	return cycle
}

// return ALL child blossom nodes from n
func (n *Node) Descendents() []*Node {
	n = n.Blossom()
	if len(n.children) == 0 {
		return []*Node{n}
	}
	nodes := []*Node{n}
	for _, c := range n.children {
		nodes = append(nodes, c.Blossom().Descendents()...)
	}
	return nodes
}

// returns All nodes (not blossom) in the blossom n
func (n *Node) All() []*Node {
	if len(n.cycle) == 0 {
		return []*Node{n}
	}
	nodes := []*Node{}
	for _, c := range n.cycle {
		nodes = append(nodes, c[0].BlossomWithin(n).All()...)
	}
	return nodes
}

func (n *Node) BlossomWithin(b *Node) *Node {
	var x = n
	for x.pp != b {
		x = x.pp
	}
	if x.pp == nil {
		panic("invalid search for blossom within")
	}
	return x
}

func (n *Node) RemoveParent() {
	for x := n; x.pp != nil; {
		x.parent = nil
		x = x.pp
	}
}

func (n *Node) PopChild(m *Node) *Node {
	for l, u := range n.children {
		ub := u.Blossom()
		if ub == m {
			n.children = append(n.children[:l], n.children[l+1:]...)
			return u
		}
	}
	return nil
}

func (n *Node) SetAlone() {
	n.parent = nil
	n.children = []*Node{}
}

// set the blossom (or nodes) as a free node
func (n *Node) SetFree() {
	// fmt.Printf("set free %d (%v)\n", b.temp, b)
	n.label = 0
	n.SetAlone()
	for _, c := range n.cycle {
		c[0].BlossomWithin(n).SetFree()
	}
}

func (n *Node) Update(delta float64) {
	n.setZ(n.z + float64(n.label)*delta)
	for _, c := range n.cycle {
		cb := c[0].BlossomWithin(n)
		cb.Update(delta)
	}
}
