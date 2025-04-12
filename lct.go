package mwpm

//#define ls ch[p][0]
//#define rs ch[p][1]
//#define Get(x) (ch[f[x]][1] == x)

var NIL *Node = nil

func ch(n *Node) int {
	if n.p.ch[1] == n {
		return 1
	}
	return 0
}

func isRoot(n *Node) bool {
	if p := n.p; p == nil || (p.ch[0] != n && p.ch[1] != n) {
		return true
	}
	return false
}

func (n *Node) update() {
	n.sz = n.z
	if l := n.ch[0]; l != nil {
		n.sz += l.sz
	}
	if l := n.ch[1]; l != nil {
		n.sz += l.sz
	}
}

func (n *Node) rotate() {
	p := n.p
	pp := p.p
	k := ch(n)
	if !isRoot(p) {
		pp.ch[ch(p)] = n
	}
	p.ch[k] = n.ch[1-k]
	if n.ch[1-k] != nil {
		n.ch[1-k].p = p
	}
	n.ch[1-k] = p
	p.p = n
	n.p = pp
	p.update()
	n.update()
}

func (n *Node) splay() {
	for !isRoot(n) {
		if !isRoot(n.p) {
			if ch(n.p) == ch(n) {
				n.p.rotate()
			} else {
				n.rotate()
			}
		}
		n.rotate()
	}
}

func (n *Node) access() {
	for x, r := n, NIL; x != nil; r, x = x, x.p {
		x.splay()
		x.ch[1] = r
		x.update()
	}
	n.splay()
}

func (n *Node) findRoot() *Node {
	x := n
	x.access()
	for x.ch[0] != nil {
		x = x.ch[0]
	}
	x.splay()
	return x
}

func (n *Node) link(p *Node) {
	n.access()
	p.access()
	if n.p != nil || n.ch[0] != nil {
		panic("duplicate link")
	}
	n.ch[0] = p
	p.p = n

	//n.p = p
	n.pp = p
	n.access()
}

func (n *Node) cut() {
	n.access()
	if p := n.ch[0]; p != nil {
		p.p = nil
	}
	n.ch[0] = nil
	n.pp = nil
	n.update()
}

func (n *Node) setZ(z float64) {
	n.splay()
	n.z = z
	n.update()
}

func (n *Node) getSz() float64 {
	n.access()
	v := n.sz
	if r := n.ch[1]; r != nil {
		v -= r.sz
	}
	return v
}
