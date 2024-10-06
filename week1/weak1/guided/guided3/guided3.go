package main

import "fmt"
import "math"

func main() {
	var jariJari int

	fmt.Print("Jejari = ")
	fmt.Scanln(&jariJari)

	volume := (4.0 / 3.0) * math.Pi * math.Pow(float64(jariJari), 3)
	luasKulit := 4 * math.Pi * math.Pow(float64(jariJari), 2)

	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.3f\n", jariJari, volume, luasKulit)
}
