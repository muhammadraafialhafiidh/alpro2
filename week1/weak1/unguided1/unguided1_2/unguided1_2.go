package main

import "fmt"

func main() {
    var N int
    var bunga string
    var pita string = ""
    var jumlahBunga int = 0

    fmt.Print("N: ")
    fmt.Scanln(&N)

    for i := 1; i <= N; i++ {
        fmt.Printf("Bunga %d: ", i)
        fmt.Scanln(&bunga)

        //cek input user apakah "SELESAI" atau bukan
        if bunga == "SELESAI" || bunga == "selesai" {
            break
        }
        //apabila input bukan "SELESAI" maka akan di tambahkan ke pita
        if pita == "" {
            pita += bunga
        } else {
            pita += " - " + bunga
        }
        jumlahBunga++ //menghitung jumlah bunga
    }

    fmt.Println("Pita:", pita)
    fmt.Println("Bunga:", jumlahBunga)
}