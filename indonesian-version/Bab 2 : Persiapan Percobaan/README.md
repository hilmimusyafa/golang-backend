# Bab 2 : Persiapan Percobaan

Pada bab ini, kita akan mempersiapkan lingkungan pengembangan untuk percobaan menggunakan bahasa pemrograman Go. Pastikan Anda telah menginstal Go di sistem Anda. Jika belum, silakan merujuk ke [dokumentasi resmi Go](https://golang.org/doc/install) untuk panduan instalasi.

## 2.1 Langkah-langkah Persiapan (Terutama Linux)

### 2.1.1 Membuat Folder Proyek

1. Buat folder proyek baru
   
   ```bash
   $ mkdir golang-backend
   $ cd golang-backend
   ```

2. Inisialisasi modul Go
   
   ```bash
   $ go mod init golang-backend
   ```

### 2.2.2 Menambahkan Dependency

1. Tambahkan dependency Gin Framework
   ```bash
   $ go get -u github.com/gin-gonic/gin
   ```

2. Perbarui dan bersihkan dependency
   
   ```bash
   $ go mod tidy
   ```

### 2.2.3 Struktur Folder
Pastikan struktur folder proyek Anda seperti berikut

```
.
├── go.mod
├── go.sum
├── chapter1/
│   ├── basic_http_server.go
│   └── gin_server.go
├── chapter2/
└── README.md
```

### 2.2.4 Menjalankan Server

1. Untuk menjalankan server HTTP dasar
   ```bash
   $ go run chapter1/basic_http_server.go
   ```

2. Untuk menjalankan server menggunakan Gin Framework
   
   ```bash
   $ go run chapter1/gin_server.go
   ```

### 2.2.5 Clone Repository

Jika Anda ingin menggunakan repository yang sudah ada, Anda dapat meng-clone repository berikut :
```bash
$ git clone https://github.com/hilmimusyafa/golang-backend.git
$ cd golang-backend
```

### 2.2.6 Testing
Pastikan server berjalan dengan baik dengan mengakses endpoint berikut:
- HTTP dasar: `http://localhost:8080`
- Gin Framework: `http://localhost:8080/hello`

## 2.2 Catatan Tambahan

- Gunakan nama folder dan file tanpa spasi untuk menghindari error.
- Selalu jalankan `go mod tidy` setelah menambahkan atau menghapus dependency.
- Jika Anda mengalami masalah, periksa dokumentasi resmi Go dan Gin Framework untuk solusi lebih lanjut.

## 2.3 Referensi

- [Dokumentasi Go](https://golang.org/doc/)
- [Dokumentasi Gin Framework](https://gin-gonic.com/)