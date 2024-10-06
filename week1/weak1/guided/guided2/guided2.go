package main

import "fmt"

// Fungsi untuk memeriksa apakah suatu tahun adalah tahun kabisat
func tahunKabisat(tahun int) bool {
	// Memeriksa apakah tahun adalah tahun kabisat
	if tahun%400 == 0 {
		return true
	} else if tahun%100 == 0 {
		return false
	} else if tahun%4 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	// Meminta input tahun dari user
	var tahun int
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahun)

	// Memeriksa apakah tahun kabisat dan menampilkan hasilnya
	if tahunKabisat(tahun) {
		fmt.Println("Kabisat: true")
	} else {
		fmt.Println("Kabisat: false")
	}
}
