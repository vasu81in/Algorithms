package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	for layer := 0; layer < N; layer++ {
		first := layer
		last := N - layer - 1
		for i := first; i < last; i++ {
			offset := i - first
			top := matrix[first][i]
			matrix[first][i] = matrix[last-offset][first]
			matrix[last-offset][first] = matrix[last][last-offset]
			matrix[last][last-offset] = matrix[i][last]
			matrix[i][last] = top
		}

	}
}

var N int = 6
var mat = [][]int{
	{1, 2, 3, 4, 5, 6},
	{7, 8, 9, 10, 11, 12},
	{13, 14, 15, 16, 17, 18},
	{19, 20, 21, 22, 23, 24},
	{25, 26, 27, 28, 29, 30},
	{31, 32, 33, 34, 35, 36},
}

func main() {
	fmt.Println("Matrix\n======")
	for i := 0; i < N; i++ {
		fmt.Printf("\n\n")
		for j := 0; j < N; j++ {
			fmt.Printf("%-3d", mat[i][j])
		}
	}
	rotate(mat)
	fmt.Println("\n\n\nMatrix rotated by 90\n====================")
	for i := 0; i < 6; i++ {
		fmt.Printf("\n\n")
		for j := 0; j < 6; j++ {
			fmt.Printf("%-3d", mat[i][j])
		}
	}

}

