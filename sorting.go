package main

import (
	"fmt"
	"sort"
)

func main() {
    // Buat slice dengan data yang akan diurutkan
    data := []int{10, 5, 8, 2, 7}

    // Menggunakan fungsi sort.Slice untuk mengurutkan slice secara descending
    sort.Slice(data, func(i, j int) bool {
        return data[i] > data[j]
    })

    // Cetak hasil pengurutan
    fmt.Println("Slice setelah diurutkan dari terbesar ke terkecil:", data)
}
