package mwpm

import (
	"fmt"
	"math"
)

type Node struct {
	label    int
	parent   *Node   // directs to non-blossom node
	children []*Node // directs to non-blossom node
	blossom  *Node   // immediate blossom that contains this node (nil if the node is the outermost blossom)
	// blossom *Blossom
	cycle [][2]*Node // cyclic pair of nodes (start, end)
	// chain []*Node    // all nodes in the blossom in the same cyclic order as cycle ()
	dval float64
	temp int64
}

const Eps = 1e-6

func NewNode() *Node {
	return &Node{
		label:    1,
		children: []*Node{},
		cycle:    [][2]*Node{},
	}
}

// returns the outermost blossom
func (n *Node) Blossom() *Node {
	if n.blossom == nil {
		return n
	}
	return n.blossom.Blossom()
}

func (n *Node) Root() *Node {
	n = n.Blossom()
	for n.parent != nil {
		n = n.parent.Blossom()
	}
	return n
}

func (n *Node) anscesters() []*Node {
	n = n.Blossom()
	chain := []*Node{n}
	for n.parent != nil {
		chain = append(chain, n.parent.Blossom())
		n = n.parent.Blossom()
	}
	return chain
}

// return ALL child blossom nodes from n
func (n *Node) descendents() []*Node {
	n = n.Blossom()
	if len(n.children) == 0 {
		return []*Node{n}
	}
	nodes := []*Node{n}
	for _, c := range n.children {
		nodes = append(nodes, c.Blossom().descendents()...)
	}
	return nodes
}

// returns all nodes (not blossom) in the blossom n
func (n *Node) all() []*Node {
	if n.blossom == nil {
		return []*Node{n}
	}
	nodes := []*Node{}
	for _, c := range n.cycle {
		nodes = append(nodes, c[0].all()...)
	}
	return nodes
}

func (n *Node) IsDvalZero() bool {
	return math.Abs(n.dval) < Eps
}

func (n *Node) IsBlossom() bool {
	return len(n.cycle) > 1
}

func (n *Node) RemoveChild(m *Node) {
	for i, c := range n.children {
		if c == m {
			n.children = append(n.children[:i], n.children[i+1:]...)
			return
		}
	}
}

func (n *Node) BlossomWithin(b *Node) *Node {
	for n.blossom != b {
		n = n.blossom
	}
	if n.blossom == nil {
		panic("invalid search for blossom within")
	}
	return n
}

func (n *Node) AllBlossoms() []*Node {
	blossoms := []*Node{n}
	for n.blossom != nil {
		blossoms = append(blossoms, n.blossom)
		n = n.blossom
	}
	return blossoms
}

func (n *Node) RemoveParent() {
	for n.blossom != nil {
		n.parent = nil
		n = n.blossom
	}
}

func (n *Node) show() {
	fmt.Printf("Node id %d label %d\n", n.temp, n.label)
}
