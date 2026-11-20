# Program Loket Tiket Wisata (Go)

Aplikasi CLI (*Command-Line Interface*) sederhana berbasis bahasa pemrograman Go untuk mengelola perhitungan tarif, diskon, dan pencetakan struk loket tiket wisata.

---

## 👤 Identitas Pengembang
- **Nama** : Dave Tyan Ronggo Putro Raharjo
- **NIM**  : 42630023

---

## 📋 Aturan Bisnis & Tarif
### Kategori Tiket:
| Kode | Kategori | Tarif / Orang |
|:---:|:---|:---|
| 1 | Dewasa | Rp 20.000 |
| 2 | Anak | Rp 10.000 |
| 3 | Turis Asing | Rp 100.000 |

### Ketentuan Diskon:
Konstanta basis `B = 8`:
- **Turis Asing (Kategori 3)**: Diskon tetap **10%**.
- **Kategori Lain (Dewasa / Anak)**:
  - Jumlah pengunjung $\ge 2 \times B$ ($\ge 16$ orang): Diskon **20%**.
  - Jumlah pengunjung $\ge B$ ($\ge 8$ orang): Diskon **10%**.
  - Kurang dari $B$: Tidak mendapat diskon (0%).

---

## 🚀 Cara Menjalankan Program

### Prasyarat
- Go versi 1.18 atau lebih baru terpasang di komputer.

### Langkah Eksekusi
1. Buka terminal pada folder proyek.
2. Inisialisasi module (opsional jika belum):
   ```bash
   go mod init loket-wisata
   ```
3. Jalankan file program:
   ```bash
   go run main.go
   ```
4. Atau lakukan build menjadi file binary:
   ```bash
   go build -o loket-wisata.exe main.go
   ./loket-wisata.exe
   ```

---

