package main

import (
	"fmt"
)

func main() {
	var beratKiri, beratKanan float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&beratKiri, &beratKanan)

		if beratKiri < 0 || beratKanan < 0 || beratKiri+beratKanan > 150 {
			break
		}

		selisih := beratKiri - beratKanan
		if selisih < 0 {
			selisih = -selisih
		}

		fmt.Printf("Sepeda motor pak Andi akan oleng: %t\n", selisih >= 9)
	}

	fmt.Println("Proses selesai.")
}
