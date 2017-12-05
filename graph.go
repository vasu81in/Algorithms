package main

import (
	"fmt"
)

type elem struct {
	v    int
	next *elem
}

type graph struct {
	size int
	list []*elem
}

func (g *graph) new_elem(v int) *elem {
	return &elem{v: v, next: nil}
}

func (g *graph) new_graph(size int) *graph {
	return &graph{size: size,
		list: make([]*elem, size)}
}
func (g *graph) init_graph(vertices int) *graph {
	return (g.new_graph(vertices))
}

func (g *graph) add_edge(v1 int, v2 int) {
	if g == nil {

	}
	if g.list[v1] == nil {
		g.list[v1] = g.new_elem(v2)
		return
	}
	walk := g.list[v1]
	for {
		if walk.next == nil {
			break
		}
		walk = walk.next
	}
	walk.next = g.new_elem(v2)
}

func (g *graph) display() {
	for i, j := range g.list {
		walk := j
		for {
			if walk == nil {
				break
			}
			fmt.Println(i, walk.v)
			walk = walk.next
		}
	}
}

func (g *graph) ()

func main() {
	g := &graph{}
	f := g.init_graph(10)
	f.add_edge(1, 2)
	f.add_edge(1, 4)
	f.add_edge(2, 3)
	f.add_edge(2, 4)
	f.add_edge(3, 4)
	f.display()
}

