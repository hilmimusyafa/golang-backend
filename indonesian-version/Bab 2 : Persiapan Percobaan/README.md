# Bab 2 : Persiapan Percobaan

Pada bab ini, kita akan mempersiapkan lingkungan pengembangan untuk percobaan menggunakan bahasa pemrograman Go. Pastikan telah menginstal Go di sistem. Jika belum, silakan merujuk ke [dokumentasi resmi Go](https://golang.org/doc/install) untuk panduan instalasi.

## 2.1 Instalasi dan Setup Gin

### 2.1.1 Instalasi Gin Framework

Gin adalah web framework yang ditulis dalam bahasa Go. Framework ini dirancang untuk menjadi cepat dan mudah digunakan, dengan performa yang sangat baik untuk membangun REST API dan web service.

#### Langkah 1 : Membuat Folder Proyek

1. Buat folder proyek baru
   
   ```bash
   $ mkdir golang-backend
   $ cd golang-backend
   ```

2. Inisialisasi modul Go
   
   ```bash
   $ go mod init golang-backend
   ```

#### Langkah 2 : Instalasi Gin Framework

1. Tambahkan dependency Gin Framework
   ```bash
   $ go get -u github.com/gin-gonic/gin
   ```

2. Perbarui dan bersihkan dependency
   
   ```bash
   $ go mod tidy
   ```

### 2.1.2 Struktur Project Backend yang Baik

Untuk project backend yang scalable dan maintainable, berikut adalah struktur folder yang direkomendasikan :

```
golang-backend/
├── go.mod
├── go.sum
├── main.go
├── config/
│   └── config.go
├── controllers/
│   ├── auth_controller.go
│   └── user_controller.go
├── models/
│   ├── user.go
│   └── response.go
├── middlewares/
│   ├── auth_middleware.go
│   └── cors_middleware.go
├── routes/
│   ├── auth_routes.go
│   ├── user_routes.go
│   └── router.go
├── services/
│   ├── auth_service.go
│   └── user_service.go
├── utils/
│   ├── jwt.go
│   └── validator.go
├── database/
│   └── connection.go
├── migrations/
│   └── create_users_table.sql
├── static/
│   ├── css/
│   ├── js/
│   └── images/
├── templates/
├── tests/
│   ├── controllers/
│   └── services/
├── docs/
│   └── api_documentation.md
└── README.md
```

#### Penjelasan Struktur :
- **config/** : Konfigurasi aplikasi (database, environment variables, dll)
- **controllers/** : Handler untuk HTTP requests
- **models/** : Struktur data dan models
- **middlewares/** : Middleware untuk authentication, logging, dll
- **routes/** : Definisi routing API
- **services/** : Business logic layer
- **utils/** : Utility functions dan helpers
- **database/** : Koneksi dan setup database
- **migrations/** : Database migration files
- **static/** : File statis (CSS, JS, gambar)
- **templates/** : Template HTML jika diperlukan
- **tests/** : Unit tests dan integration tests
- **docs/** : Dokumentasi API

### 2.1.3 First Gin Application

Mari buat aplikasi Gin pertama kita dengan struktur yang sederhana:

```go
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize Gin with default middleware (Logger and Recovery)
    r := gin.Default()

    // Define route for root path ("/") with GET method
    r.GET("/", func(c *gin.Context) {
        // Send JSON response with status 200 OK
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello From Gin!",
        })
    })

    // Run server on port 8080
    r.Run(":8080")
}
```

### 2.1.4 Menjalankan Server

1. Untuk menjalankan aplikasi Gin pertama Anda :
   
   ```bash
   $ go run main.go
   ```

2. Untuk menjalankan server HTTP dasar (jika ada) :
   
   ```bash
   $ go run chapter1/basic_http_server.go
   ```

3. Untuk menjalankan server menggunakan Gin Framework (jika ada) :
   
   ```bash
   $ go run chapter1/gin_server.go
   ```

### 2.1.5 Testing Endpoints

Setelah server berjalan, Anda dapat menguji endpoint berikut :

1. **GET Route dasar**
   
   ```bash
   curl http://localhost:8080/
   ```

2. **GET dengan parameter**
   
   ```bash
   curl http://localhost:8080/hello/John
   ```

3. **GET dengan query parameter**
   
   ```bash
   curl "http://localhost:8080/search?q=golang"
   ```

4. **POST untuk membuat user:**
   
   ```bash
   curl -X POST http://localhost:8080/users \
     -H "Content-Type: application/json" \
     -d '{"name":"John Doe","email":"john@example.com"}'
   ```

## 2.2 Clone Repository (Opsional)

Jika Anda ingin menggunakan repository yang sudah ada, Anda dapat meng-clone repository berikut :

```bash
$ git clone https://github.com/hilmimusyafa/golang-backend.git
$ cd golang-backend
```

## 2.3 Catatan Tambahan

- Gunakan nama folder dan file tanpa spasi untuk menghindari error.
- Selalu jalankan `go mod tidy` setelah menambahkan atau menghapus dependency.
- Gin menyediakan hot reload dengan menggunakan tools seperti `air` atau `gin` untuk development.
- Untuk production, pastikan untuk menggunakan `gin.SetMode(gin.ReleaseMode)` untuk performance yang optimal.
- Jika Anda mengalami masalah, periksa dokumentasi resmi Go dan Gin Framework untuk solusi lebih lanjut.

### Tips Development

1. Hot Reload dengan Air
   
   ```bash
   $ go install github.com/cosmtrek/air@latest
   $ air
   ```

2. Environment Variables 
   
   Gunakan package seperti `godotenv` untuk mengelola environment variables :
   
   ```bash
   $ go get github.com/joho/godotenv
   ```

3. Middleware yang Berguna
   
   - CORS : `github.com/gin-contrib/cors`
   - Logger : Built-in di Gin
   - Recovery : Built-in di Gin
   - JWT Authentication : `github.com/golang-jwt/jwt`

## 2.4 Referensi

- [Dokumentasi Go](https://golang.org/doc/)
- [Dokumentasi Gin Framework](https://gin-gonic.com/)
- [Gin GitHub Repository](https://github.com/gin-gonic/gin)
- [Go Modules Documentation](https://golang.org/ref/mod)
- [Best Practices untuk Go Project Structure](https://github.com/golang-standards/project-layout)