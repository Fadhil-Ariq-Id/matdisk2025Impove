package main

import "fmt"

func perkalianMatriks(A, B [][]int) [][]int {
	n := len(A)
	R := make([][]int, n)
	for i := 0; i < n; i++ {
		R[i] = make([]int, n)
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				R[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return R
}

func traceMatriks(M [][]int) int {
	sum := 0
	for i := 0; i < len(M); i++ {
		sum += M[i][i]
	}
	return sum
}

func transformasiBaris(M [][]int) {
	n := len(M)
	M[0], M[n-1] = M[n-1], M[0]
	fmt.Println("Matrix setelah tukar baris:", M)

	max := M[0][0]
	for i := range M {
		for j := range M[i] {
			if M[i][j] > max {
				max = M[i][j]
			}
		}
	}
	fmt.Println("Nilai maksimum:", max)
}
