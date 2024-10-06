package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)

	fmt.Print("Masukkan input string pertama: ")
	fmt.Scanln(&satu)

	fmt.Print("Masukkan input string kedua: ")
	fmt.Scanln(&dua)

	fmt.Print("Masukkan input string ketiga: ")
	fmt.Scanln(&tiga)

	// Print the initial output
	fmt.Println("Output awal: " + satu + "  " + dua + "  " + tiga)

	// Swap the values
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	// Print the final output after swapping
	fmt.Println("Output akhir: " + satu + "  " + dua + "  " + tiga)
}
