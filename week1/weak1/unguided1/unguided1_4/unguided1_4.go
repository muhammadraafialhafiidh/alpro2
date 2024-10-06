package main

import (
	"fmt"
)

func main() {
	var k int
	var result float64 = 1.0

	// Menerima input K
	fmt.Print("Nilai K = ")
	fmt.Scanln(&k)

	for i := 0; i < k; i++ {
		result = 0.5 * (result + 2/result)
	}

	fk := float64((4*k + 2) * (4*k + 2)) / float64((4*k + 1) * (4*k + 3))
	fmt.Printf("Nilai f(K) = %.10f\n", fk)

	// Menampilkan hasil akar 2
	fmt.Printf("Nilai akar 2 = %.10f\n", result)
}
