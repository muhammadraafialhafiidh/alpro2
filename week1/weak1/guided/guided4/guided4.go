package main

import "fmt"

func main() {
	var celsius, fahrenheit, reamur, kelvin float64

	fmt.Print("Temperatur Celsius: ")
	fmt.Scanln(&celsius)

	// Menghitung Fahrenheit
	fahrenheit = (celsius * 9 / 5) + 32

	// Menghitung Reamur
	reamur = celsius * 4 / 5

	// Menghitung Kelvin
	kelvin = celsius + 273.15

	fmt.Println("Derajat Reamur:", reamur)
	fmt.Println("Derajat Fahrenheit:", fahrenheit)
	fmt.Println("Derajat Kelvin:", kelvin)
}
