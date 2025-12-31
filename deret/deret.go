package main

import "fmt"

func rekursif(C1, C2, N int) {
	a := make([]int, N+1)
	a[0], a[1] = 1, 1

	for i := 2; i <= N; i++ {
		a[i] = C1*a[i-1] + C2*a[i-2]
	}
	fmt.Println("Suku ke-", N, "=", a[N])
}

func deretGeometri(a, r float64, N int) {
	sn := a * (1 - pow(r, N)) / (1 - r)
	sinf := a / (1 - r)
	fmt.Println("Sn =", sn)
	fmt.Println("S∞ =", sinf)
	fmt.Printf("Kedekatan = %.2f%%\n", (sn/sinf)*100)
}

func pow(a float64, n int) float64 {
	res := 1.0
	for i := 0; i < n; i++ {
		res *= a
	}
	return res
}
