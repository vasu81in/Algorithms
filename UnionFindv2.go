package main

import (
	"fmt"
)

type UnionFind struct {
	N    int
	elem []int
}

// Improved version of Union Find

func (this *UnionFind) Init(n int) {
	this.N = n
	this.elem = make([]int, n, n)
	for i := 0; i < len(this.elem); i++ {
		this.elem[i] = i
	}
}

func (this *UnionFind) Root(i int) int {
	for {
		if this.elem[i] != i {
			i = this.elem[i]
		} else {
			return i
		}
	}
}

func (this *UnionFind) Find(a int, b int) bool {
	return (this.Root(a) == this.Root(b))
}

func (this *UnionFind) Union(a int, b int) {
	i := this.Root(a)
	j := this.Root(b)
	this.elem[i] = j
}

//[0 1 2 3 4 5 6 7 8 9]
//[1 1 2 3 4 5 6 7 8 9]
//[1 3 2 3 4 5 6 7 8 9]
//true

func main() {
	uf := new(UnionFind)
	uf.Init(10)
	fmt.Println(uf.elem)
	uf.Union(0, 1)
	fmt.Println(uf.elem)
	uf.Union(1, 3)
	fmt.Println(uf.elem)
	fmt.Println(uf.Find(1,3))
}

