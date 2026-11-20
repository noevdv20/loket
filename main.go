package main

import "fmt"

const B = 8

func getTarif(kategori int) (string, int, bool) {
	switch kategori {
	case 1:
		return "Dewasa", 20000, true
	case 2:
		return "Anak", 10000, true
	case 3:
		return "Turis Asing", 100000, true
	default:
		return "", 0, false
	}
}

func hitungDiskon(kategori, n int) int {
	if kategori == 3 {
		return 10
	} else if n >= 2*B {
		return 20
	} else if n >= B {
		return 10
	}
	return 0
}

func main() {
	fmt.Println("Nama : Dave Tyan Ronggo Putro Raharjo")
	fmt.Println("NIM  : 42630023")

	var kategori, n int

	fmt.Println("\n===== LOKET TIKET WISATA =====")
	fmt.Print("Masukkan kode kategori (1/2/3): ")
	if _, err := fmt.Scan(&kategori); err != nil {
		fmt.Println("Error: Masukan harus berupa angka!")
		return
	}

	namaKategori, tarif, valid := getTarif(kategori)
	if !valid {
		fmt.Println("Kategori tidak valid!")
		return
	}

	fmt.Print("Masukkan jumlah pengunjung: ")
	if _, err := fmt.Scan(&n); err != nil || n < 1 || n > 100 {
		fmt.Println("Jumlah pengunjung tidak valid (harus antara 1 - 100)!")
		return
	}

	total := tarif * n
	diskonPersen := hitungDiskon(kategori, n)
	potongan := total * diskonPersen / 100
	totalBayar := total - potongan

	fmt.Println("\n===== STRUK TIKET =====")
	fmt.Println("Kategori       :", namaKategori)
	fmt.Println("Tarif per orang: Rp", tarif)
	fmt.Println("Jumlah         :", n)
	fmt.Println("Diskon         :", diskonPersen, "%")
	fmt.Println("Total bayar    : Rp", totalBayar)
}