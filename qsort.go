package main

import (
	"fmt"
)

func partition(arr []int, start int, end int) int {
	pivot := arr[end] // Last element is the pivot
	pIndex := start   // Let the pindex be the first
	for i := start; i < end; i++ {
		if arr[i] <= pivot {
			arr[i], arr[pIndex] = arr[pIndex], arr[i]
			pIndex = pIndex + 1
		}
	}
	arr[pIndex], arr[end] = arr[end], arr[pIndex]
	return pIndex
}

func qsort(arr []int, start int, end int) {
	if start >= end {
		return
	}
	pIndex := partition(arr, start, end)
	fmt.Println("array ", arr, "index ", pIndex)
	qsort(arr, start, pIndex-1)
	qsort(arr, pIndex+1, end)
}

func main() {
	arr := []int{3, 5, 1, 2, 7, 6, 4}
	qsort(arr, 0, len(arr)-1)
	fmt.Println("Sorted Array: ", arr)
}

