package main

import (
	"fmt"
)

func main() {
	var beratParsel, beratKg, sisaBerat, biayaKg, biayaSisa, totalBiaya int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&beratParsel)

	beratKg = beratParsel / 1000
	sisaBerat = beratParsel % 1000
	biayaKg = beratKg * 10000

	if sisaBerat >= 500 {
		biayaSisa = sisaBerat * 5
	} else if beratParsel <= 10000 {
		biayaSisa = sisaBerat * 15
	}

	totalBiaya = biayaKg + biayaSisa

	fmt.Println("Detail berat:", beratKg, "kg +", sisaBerat, "gr")
	fmt.Println("Detail biaya: Rp.", biayaKg, "+ Rp.", biayaSisa)
	fmt.Println("Total biaya: Rp.", totalBiaya)
}