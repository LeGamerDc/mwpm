package mwpm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newNode(z float64) *Node {
	return &Node{
		z: z, sz: z,
	}
}

func ff(f *Node, ch ...*Node) {
	for _, c := range ch {
		c.link(f)
	}
}

func TestLct(t *testing.T) {
	a := newNode(1)
	b := newNode(2)
	c := newNode(3)
	d := newNode(4)
	e := newNode(5)
	f := newNode(6)
	g := newNode(7)
	h := newNode(8)
	i := newNode(9)
	j := newNode(10)
	k := newNode(11)
	l := newNode(12)
	m := newNode(13)
	n := newNode(14)
	o := newNode(15)

	ff(a, b, c)
	ff(b, d, e, f)
	ff(c, g)
	ff(g, h)
	ff(h, i, j)
	ff(i, k, l, m)
	ff(l, n)
	ff(n, o)

	assert.Equal(t, 1.0, a.getSz())
	assert.Equal(t, 3.0, b.getSz())
	assert.Equal(t, 4.0, c.getSz())
	assert.Equal(t, 7.0, d.getSz())
	assert.Equal(t, 8.0, e.getSz())
	assert.Equal(t, 9.0, f.getSz())
	assert.Equal(t, 11.0, g.getSz())
	assert.Equal(t, 19.0, h.getSz())
	assert.Equal(t, 28.0, i.getSz())
	assert.Equal(t, 29.0, j.getSz())
	assert.Equal(t, 39.0, k.getSz())
	assert.Equal(t, 40.0, l.getSz())
	assert.Equal(t, 41.0, m.getSz())
	assert.Equal(t, 54.0, n.getSz())
	assert.Equal(t, 69.0, o.getSz())

	i.cut()
	assert.Equal(t, 1.0, a.getSz())
	assert.Equal(t, 3.0, b.getSz())
	assert.Equal(t, 4.0, c.getSz())
	assert.Equal(t, 7.0, d.getSz())
	assert.Equal(t, 8.0, e.getSz())
	assert.Equal(t, 9.0, f.getSz())
	assert.Equal(t, 11.0, g.getSz())
	assert.Equal(t, 19.0, h.getSz())
	assert.Equal(t, 9.0, i.getSz())
	assert.Equal(t, 29.0, j.getSz())
	assert.Equal(t, 20.0, k.getSz())
	assert.Equal(t, 21.0, l.getSz())
	assert.Equal(t, 22.0, m.getSz())
	assert.Equal(t, 35.0, n.getSz())
	assert.Equal(t, 50.0, o.getSz())

	i.link(e)
	assert.Equal(t, 1.0, a.getSz())
	assert.Equal(t, 3.0, b.getSz())
	assert.Equal(t, 4.0, c.getSz())
	assert.Equal(t, 7.0, d.getSz())
	assert.Equal(t, 8.0, e.getSz())
	assert.Equal(t, 9.0, f.getSz())
	assert.Equal(t, 11.0, g.getSz())
	assert.Equal(t, 19.0, h.getSz())
	assert.Equal(t, 29.0, j.getSz())
	assert.Equal(t, 8.0+9.0, i.getSz())
	assert.Equal(t, 8.0+20.0, k.getSz())
	assert.Equal(t, 8.0+21.0, l.getSz())
	assert.Equal(t, 8.0+22.0, m.getSz())
	assert.Equal(t, 8.0+35.0, n.getSz())
	assert.Equal(t, 8.0+50.0, o.getSz())

	fmt.Println("ok")
}
