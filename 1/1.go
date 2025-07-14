package main

import (
    "fmt"
)

func calculateSum(x, y int) int {
    return x + y
}

func main() {
    x := 48
    y := 48
    result := calculateSum(x, y)
    fmt.Printf("Hasil penjumlahan %d + %d adalah %d\n", x, y, result)
}

// Saya memilih jawaban E
// Alasan : 
// 1. Karena bahasa pemrograman go lebih eksplisit
// 2. Function tersebut sudah jelas yaitu apa yang dilakukan oleh function
// Dan saya coba running dengan menggunakan go run dan hasilnya 
// Sesuai apa yang dilakukan oleh function tersebut