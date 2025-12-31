package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// ================= HIMPUNAN =================
	fmt.Println("========== HIMPUNAN ==========")

	A := rand.Perm(105)[:5]
	B := rand.Perm(105)[:5]
	C := rand.Perm(105)[:3]

	fmt.Println("A:", A)
	fmt.Println("B:", B)
	fmt.Println("C:", C)
	fmt.Println("Hasil operasi himpunan:", operasiHimpunan(A, B, C))

	S := []int{1, 5, 7, 9, 12, 15, 18, 20, 22}
	pasanganSubset(S, 24)

	// ================= MATRIX =================
	fmt.Println("\n========== MATRIX ==========")

	Amt := [][]int{{1, 2}, {3, 4}}
	Bmt := [][]int{{5, 6}, {7, 8}}

	R := perkalianMatriks(Amt, Bmt)
	fmt.Println("Hasil perkalian matriks:", R)
	fmt.Println("Trace matriks:", traceMatriks(R))

	transformasiBaris([][]int{
		{1, 2},
		{10, 5},
		{4, 6},
		{7, 8},
		{9, 3},
	})

	// ================= DERET =================
	fmt.Println("\n========== DERET ==========")

	rekursif(3, 4, 11)
	deretGeometri(7, 0.3, 12)
}
