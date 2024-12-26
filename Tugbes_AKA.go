package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Fungsi untuk membuat dataset acak
func buatDataset(ukuran int) []int {
	dataset := make([]int, ukuran)
	for i := 0; i < ukuran; i++ {
		dataset[i] = rand.Intn(100) // Nilai acak antara 0 hingga 99999
	}
	return dataset
}

// Fungsi pengurutan secara iteratif
func urutkanIteratif(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] < data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}

// Fungsi pengurutan secara rekursif
func urutkanRekursif(data []int, indeks int) []int {
	if indeks == len(data)-1 {
		return data
	}
	indeksMax := indeks
	for i := indeks + 1; i < len(data); i++ {
		if data[i] > data[indeksMax] {
			indeksMax = i
		}
	}
	data[indeks], data[indeksMax] = data[indeksMax], data[indeks]
	return urutkanRekursif(data, indeks+1)
}

func main() {
	fmt.Println("Selamat datang di aplikasi pengurutan data acak!")

	for {
		var ukuran int
		fmt.Print("Masukkan ukuran dataset (ketik 0 untuk keluar): ")
		_, err := fmt.Scan(&ukuran)
		if err != nil || ukuran <= 0 {
			if ukuran == 0 {
				fmt.Println("Program berhenti. Terima kasih telah menggunakan aplikasi ini!")
				break
			}
			fmt.Println("Input tidak valid. Harap masukkan angka positif.")
			continue
		}

		// Membuat dataset acak
		data := buatDataset(ukuran)

		// Pengurutan iteratif
		start := time.Now()
		urutkanIteratif(data)
		waktuIteratif := time.Since(start).Seconds()

		// Pengurutan rekursif
		start = time.Now()
		urutkanRekursif(data, 0)
		waktuRekursif := time.Since(start).Seconds()

		// Tampilkan hasil waktu
		fmt.Printf("\n=== Hasil Pengurutan ===\n")
		fmt.Printf("Iteratif: Waktu: %.6f detik\n", waktuIteratif)
		fmt.Printf("Rekursif: Waktu: %.6f detik\n", waktuRekursif)
	}
}
