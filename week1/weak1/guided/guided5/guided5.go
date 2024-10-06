package main

import (
	"fmt"
)

func main() {
	var integers [5]int
	var chars [3]rune

	fmt.Println("Masukkan 5 buah data integer (32-127):")
	for i := 0; i < 5; i++ {
		fmt.Scan(&integers[i])
	}

	fmt.Println("Masukkan 3 buah karakter:")
	for i := 0; i < 3; i++ {
		fmt.Scanf("%c", &chars[i])
	}

	fmt.Println("\nKeluaran:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%c", rune(integers[i]))
	}
	fmt.Println()
	for i := 0; i < 3; i++ {
		fmt.Printf("%c", chars[i]+1)
	}
}
