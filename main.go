package main

import (
    "fmt"
    "math"
)

// Fungsi untuk menghitung luas persegi
func hitungLuasPersegi(sisi float64) float64 {
    return sisi * sisi
}

// Fungsi untuk menghitung luas segitiga
func hitungLuasSegitiga(alas float64, tinggi float64) float64 {
    return 0.5 * alas * tinggi
}

// Fungsi untuk menghitung luas lingkaran
func hitungLuasLingkaran(jariJari float64) float64 {
    return math.Pi * jariJari * jariJari
}

// Fungsi untuk menghitung energi potensial
func hitungEnergiPotensial(massa float64, tinggi float64) float64 {
    g := 9.81 // Percepatan gravitasi bumi (m/s^2)
    return massa * g * tinggi
}

// Fungsi untuk menghitung energi kinetik
func hitungEnergiKinetik(massa float64, kecepatan float64) float64 {
    return 0.5 * massa * kecepatan * kecepatan
}

// Fungsi untuk menghitung frekuensi dari periode
func hitungFrekuensi(periode float64) float64 {
    return 1 / periode
}

// Fungsi untuk mengkonversi suhu dari Celsius ke Fahrenheit
func konversiCelsiusKeFahrenheit(celsius float64) float64 {
    return (celsius * 9/5) + 32
}

// Fungsi untuk mengkonversi suhu dari Fahrenheit ke Celsius
func konversiFahrenheitKeCelsius(fahrenheit float64) float64 {
    return (fahrenheit - 32) * 5/9
}

func main() {
    // Contoh penggunaan fungsi-fungsi di atas
    fmt.Println("Luas Persegi:", hitungLuasPersegi(5))
    fmt.Println("Luas Segitiga:", hitungLuasSegitiga(6, 8))
    fmt.Println("Luas Lingkaran:", hitungLuasLingkaran(4))
    fmt.Println("Energi Potensial:", hitungEnergiPotensial(10, 15))
    fmt.Println("Energi Kinetik:", hitungEnergiKinetik(5, 10))
    fmt.Println("Frekuensi:", hitungFrekuensi(0.02))
    fmt.Println("Konversi Celsius ke Fahrenheit:", konversiCelsiusKeFahrenheit(25))
    fmt.Println("Konversi Fahrenheit ke Celsius:", konversiFahrenheitKeCelsius(98.6))
}
