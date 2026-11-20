package main

import "fmt"

func main() {
    fmt.Println("Nama:Dave Tyan Ronggo Putro Raharjo")
    fmt.Println("NIM: 42630023")

    const B = 8

    var kategori int
    var n int
    var namaKategori string
    var tarif int

    fmt.Println("\n===== LOKET TIKET WISATA =====")

    fmt.Print("Masukkan kode kategori (1/2/3): ")
    fmt.Scan(&kategori)

    switch kategori {
    case 1:
        namaKategori = "Dewasa"
        tarif = 20000
    case 2:
        namaKategori = "Anak"
        tarif = 10000
    case 3:
        namaKategori = "Turis Asing"
        tarif = 100000
    default:
        fmt.Println("Kategori tidak valid")
        return
    }

    fmt.Print("Masukkan jumlah pengunjung: ")
    fmt.Scan(&n)

    if n < 1 || n > 100 {
        fmt.Println("Jumlah tidak valid")
        return
    }

    total := tarif * n
    diskon := 0

    if kategori == 3 {
        diskon = 10
    } else if n >= 2*B {
        diskon = 20
    } else if n >= B {
        diskon = 10
    }

    totalBayar := total - (total * diskon / 100)

    fmt.Println("\n===== STRUK TIKET =====")
    fmt.Println("Kategori      :", namaKategori)
    fmt.Println("Tarif per orang: Rp", tarif)
    fmt.Println("Jumlah         :", n)
    fmt.Println("Diskon         :", diskon, "%")
    fmt.Println("Total bayar    : Rp", totalBayar)
}