package main

import (
	"fmt"
)

func merge(a []int, b []int) []int {
	var i, j, k int
	c := make([]int, len(a)+len(b))
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			c[k] = a[i]
			k++
			i++
		} else {
			c[k] = b[j]
			k++
			j++
		}
	}
	for i < len(a) {
		c[k] = a[i]
		k++
		i++
	}
	for j < len(b) {
		c[k] = b[j]
		k++
		j++
	}
	fmt.Println("Merging ", c)
	return c
}

func mergesort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	mid := len(a) / 2
	left := mergesort(a[:mid])
	right := mergesort(a[mid:])
	return merge(left, right)
}

func main() {
	fmt.Println("Hello, playground", mergesort([]int{5, 3, 4, 1, 2}))
}
