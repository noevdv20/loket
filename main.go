package main

import (
	"fmt"
	"strconv"
)

const B = 8

// Struct data tiket wisata
type TiketWisata struct {
	KodeKategori int
	NamaKategori string
	Tarif        int
	Jumlah       int
	Diskon       int
	TotalAwal    int
	TotalBayar   int
}

func formatRupiah(n int) string {
	in := strconv.Itoa(n)
	out := ""
	length := len(in)

	for i, c := range in {
		out += string(c)
		sisa := length - i - 1
		if sisa > 0 && sisa%3 == 0 {
			out += "."
		}
	}
	return "Rp " + out
}

func buatTiket(kategori, jumlah int) (*TiketWisata, error) {
	var nama string
	var tarif int

	switch kategori {
	case 1:
		nama, tarif = "Dewasa", 20000
	case 2:
		nama, tarif = "Anak", 10000
	case 3:
		nama, tarif = "Turis Asing", 100000
	default:
		return nil, fmt.Errorf("kategori tidak valid")
	}

	diskon := 0
	if kategori == 3 {
		diskon = 10
	} else if jumlah >= 2*B {
		diskon = 20
	} else if jumlah >= B {
		diskon = 10
	}

	total := tarif * jumlah
	totalBayar := total - (total * diskon / 100)

	return &TiketWisata{
		KodeKategori: kategori,
		NamaKategori: nama,
		Tarif:        tarif,
		Jumlah:       jumlah,
		Diskon:       diskon,
		TotalAwal:    total,
		TotalBayar:   totalBayar,
	}, nil
}

func cetakStruk(t *TiketWisata) {
	fmt.Println("\n==============================")
	fmt.Println("         STRUK TIKET          ")
	fmt.Println("==============================")
	fmt.Printf("%-16s: %s\n", "Kategori", t.NamaKategori)
	fmt.Printf("%-16s: %s\n", "Tarif Satuan", formatRupiah(t.Tarif))
	fmt.Printf("%-16s: %d orang\n", "Jumlah", t.Jumlah)
	fmt.Printf("%-16s: %d%%\n", "Diskon", t.Diskon)
	fmt.Println("------------------------------")
	fmt.Printf("%-16s: %s\n", "Total Bayar", formatRupiah(t.TotalBayar))
	fmt.Println("==============================")
}

func main() {
	fmt.Println("Nama : Dave Tyan Ronggo Putro Raharjo")
	fmt.Println("NIM  : 42630023")

	var kategori, n int
	fmt.Println("\n===== LOKET TIKET WISATA =====")
	fmt.Print("Masukkan kode kategori (1: Dewasa, 2: Anak, 3: Asing): ")
	fmt.Scan(&kategori)

	fmt.Print("Masukkan jumlah pengunjung (1-100): ")
	fmt.Scan(&n)

	if n < 1 || n > 100 {
		fmt.Println("Error: Jumlah pengunjung di luar batas valid!")
		return
	}

	tiket, err := buatTiket(kategori, n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cetakStruk(tiket)
}