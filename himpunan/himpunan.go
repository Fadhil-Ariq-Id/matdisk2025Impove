package main

import "fmt"

// Soal 1
func operasiHimpunan(A, B, C []int) []int {
	result := []int{}

	for _, a := range A {
		found := false
		for _, b := range B {
			if a == b {
				found = true
			}
		}
		if !found {
			result = append(result, a)
		}
	}

	for _, b := range B {
		found := false
		for _, a := range A {
			if b == a {
				found = true
			}
		}
		if !found {
			result = append(result, b)
		}
	}

	final := []int{}
	for _, x := range result {
		found := false
		for _, c := range C {
			if x == c {
				found = true
			}
		}
		if !found {
			final = append(final, x)
		}
	}
	return final
}

// Soal 2
func pasanganSubset(S []int, K int) {
	fmt.Println("Pasangan (x,y) dengan x+y >", K)
	total := 0
	for i := 0; i < len(S); i++ {
		for j := i + 1; j < len(S); j++ {
			if S[i]+S[j] > K {
				fmt.Println(S[i], S[j])
				total++
			}
		}
	}
	fmt.Println("Total pasangan:", total)
}
