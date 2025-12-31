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
