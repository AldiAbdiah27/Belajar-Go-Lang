package main

import "fmt"

func main() {
    // Buat sebuah slice dengan 5 data awal
    data := []int{1, 2, 3, 4, 5}

    // Tambahkan 3 data ke dalam slice
    data = append(data, 6, 7, 8)

    // Cetak slice setelah penambahan data
    fmt.Println("Slice setelah penambahan data:", data)
}
