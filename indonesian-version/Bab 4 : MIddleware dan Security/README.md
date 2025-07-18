# Bab 4 : Middleware dan Security

Dalam pengembangan aplikasi web, middleware adalah tulang punggung yang memungkinkan kita menambahkan fungsionalitas di antara penerimaan permintaan dan pengiriman respons, tanpa harus menulis ulang kode di setiap handler.

Sementara itu, keamanan adalah aspek krusial yang tidak bisa diabaikan. Gin menyediakan berbagai built-in middleware yang sangat berguna untuk kedua tujuan ini. Bab ini akan membahas middleware bawaan Gin yang sering digunakan dan bagaimana penerapannya membantu dalam membangun aplikasi yang lebih robust dan aman.

## 4.1 Built-in Middlewares

Middleware bawaan adalah fungsi-fungsi yang sudah disediakan oleh framework Gin, siap pakai untuk mempermudah tugas-tugas umum seperti pencatatan (logging), penanganan error, atau bahkan pengaturan keamanan dasar. Menggunakan middleware ini akan sangat menghemat waktu pengembangan dan membantu menjaga konsistensi aplikasi.

### 4.1.1 Logger Middleware

Logger middleware adalah salah satu middleware paling dasar dan paling sering digunakan. Fungsinya adalah untuk mencatat setiap permintaan HTTP yang masuk ke server. Catatan ini sangat berguna untuk debugging, pemantauan kinerja, dan memahami pola lalu lintas ke aplikasi.

Secara default, ketika menginisialisasi Gin dengan `gin.Default()`, Logger middleware sudah disertakan secara otomatis. Ini berarti tidak perlu secara eksplisit menambahkannya jika menggunakan `gin.Default()`. Informasi yang Dicatat oleh Logger Middleware :

- Metode HTTP : `GET`, `POST`, `PUT`, `DELETE`, dll.
- Path URL : Jalur yang diakses oleh klien (misalnya `/users`, `/products/1`).
- Alamat IP Klien : Alamat IP dari mana permintaan berasal.
- Kode Status HTTP Respons : `200 OK`, `404 Not Found`, `500 Internal Server Error`, dll.
- Waktu Pemrosesan : Durasi yang dibutuhkan server untuk memproses permintaan (dari awal hingga respons dikirim).
- Ukuran Respons : Ukuran body respons yang dikirimkan.

Untuk contoh tampilan log (dari terminal) :

```bash
[GIN] 2025/07/01 - 15:46:55 | 200 |      36.007µs |       127.0.0.1 | GET      "/"
```

Keterangan dari baris log di atas menunjukkan :

- `[GIN]` : Awalan log dari Gin.
- `2025/07/01 - 15:46:55` : Tanggal dan waktu permintaan diterima.
- `200` : Kode status HTTP respons (sukses).
- `36.007µs` : Waktu yang dibutuhkan untuk memproses permintaan.
- `127.0.0.1` : Alamat IP klien (localhost).
- `GET` : Metode HTTP.
- `/` : Path URL yang diakses.

Logger Middleware sudah otomatis aktif jika menggunakan `gin.Default()`. Namun, jika ingin kontrol lebih, bisa menginisialisasi engine dengan `gin.New()` (yang tidak menyertakan middleware bawaan), lalu menambahkan `gin.Logger()` secara manual sesuai kebutuhan. Berikut contoh langsungnya :

4.1.1.TryLoggerMiddleware.go

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize a Gin machine without built-in middleware
	r := gin.New()

	// Add Logger middleware manually
	r.Use(gin.Logger())

	// Simple route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from Gin!"})
	})

	r.Run(":8080")
}
```

Kita coba kode tersebut :

```bash
$ go run 4.1.1.TryLoggerMiddleware.go
```

Saat sudah jalan kita akses seperti biasa :

```
http://localhost:8080
```

Dan saat akses maka akan muncul pada log seperti di bawah :

```bash
$ go run 4.1.1.TryLoggerMiddleware.go
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func1 (2 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2025/07/14 - 21:15:00 | 200 |       26.55µs |       127.0.0.1 | GET      "/"
[GIN] 2025/07/14 - 21:15:00 | 404 |         130ns |       127.0.0.1 | GET      "/favicon.ico"
```

Logger middleware sangat berguna digunakan dalam berbagai situasi, seperti debugging untuk mengidentifikasi request dan respons yang masuk dan keluar, pemantauan aktivitas server secara real-time, serta analisis data log untuk memahami lalu lintas atau kinerja aplikasi. 

Dengan memberikan visibilitas yang diperlukan ke dalam interaksi server-klien, logger middleware menjadi alat yang tidak ternilai bagi setiap aplikasi Gin.

Untuk mencoba kode tersebut bisa akses [4.1.1.TryLoggerMiddleware.go](../../source-code/chapter4/4.1.1.TryLoggerMiddleware.go)

### 4.1.2 Recovery Middleware

Recovery middleware adalah penjaga gawang aplikasi. Dalam Go, panic adalah error runtime yang tidak tertangkap yang dapat menyebabkan program berhenti total (crash). Dalam aplikasi web, ini berarti server akan mati dan tidak bisa lagi melayani permintaan. Recovery middleware Gin dirancang untuk mencegah hal ini.

Sama seperti Logger middleware, Recovery middleware juga disertakan secara otomatis ketika menggunakan `gin.Default()`. Fungsi Utama Recovery Middleware :

1. Menangkap Panic : Jika terjadi panic di salah satu handler atau middleware dalam rantai pemrosesan permintaan, Recovery middleware akan menangkapnya.
2. Mencegah Crash : Dengan menangkap panic, middleware ini mencegah seluruh server berhenti berfungsi. Hanya permintaan yang menyebabkan panic itulah yang akan terpengaruh.
3. Mencatat Panic : Recovery middleware akan mencatat detail panic (termasuk stack trace) ke konsol atau log server. Ini sangat membantu untuk debugging masalah yang tidak terduga.
4. Mengirim Respons Error 500: Kepada klien yang membuat permintaan, Recovery middleware akan mengirimkan respons HTTP 500 Internal Server Error, memastikan klien mendapatkan indikasi bahwa ada masalah di sisi server.

Contoh implementasi ada di kode :

4.1.2.TryRecoveryMidddleware.go

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Using gin.Default() which includes Recovery middleware
	r := gin.Default()

	r.GET("/safe", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This is a safe route"})
	})

	r.GET("/cause-panic", func(c *gin.Context) {
		// Simulating panic: attempting to access index outside slice bounds
		var data []string
		_ = data[0] // This will cause panic!
		c.JSON(http.StatusOK, gin.H{"message": "This line will never be reached"})
	})

	r.Run(":8080")
}
```

Bagian kode di atas yang perlu di bahas dan menjadi catatan : 

- `var data []string` : Mendefinisikan slice kosong.
- `data[0]` : Mengakses elemen pertama dari slice kosong menyebabkan **panic: index out of range**.
- Baris `c.JSON(...)` tidak akan pernah dijalankan karena eksekusi sudah berhenti akibat panic.
- **Namun server tidak akan mati** karena middleware `Recovery` dari Gin akan menangani panic dan merespons error 500.

Kita coba kode tersebut :

```bash
$ go run 4.1.2.TryRecoveryMidddleware.go
```

Sesuai kode kita coba akses endpoint dengan URL :

```
http://localhost:8080/safe
```

Maka server dengan normal akan merespons sukses seperti log di bawah : 

```bash
$ go run 4.1.2.TryRecoveryMidddleware.go 
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /safe                     --> main.main.func1 (3 handlers)
[GIN-debug] GET    /cause-panic              --> main.main.func2 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2025/07/17 - 22:08:02 | 200 |      35.156µs |       127.0.0.1 | GET      "/safe"
[GIN] 2025/07/17 - 22:08:02 | 404 |         732ns |       127.0.0.1 | GET      "/favicon.ico"
```

dengan hasil yang sesusi seperti di bawah :

```json
{"message":"This is a safe route"}
```

Oke, sekarang kita akan coba dengan menggunakan akses yang satunya, yaitu :

```
http://localhost:8080/cause-panic
```

Maka hasilnya akan berbeda : 

```json
`{"message":"Internal Server Error"}`
```

dengan log :
```bash
2025/07/17 22:18:37 [Recovery] 2025/07/17 - 22:18:37 panic recovered:
GET /cause-panic HTTP/1.1
Host: localhost:8080
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate, br, zstd
Accept-Language: en-US,en;q=0.5
Connection: keep-alive
Priority: u=0, i
Sec-Fetch-Dest: document
Sec-Fetch-Mode: navigate
Sec-Fetch-Site: none
Sec-Fetch-User: ?1
Sec-Gpc: 1
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:140.0) Gecko/20100101 Firefox/140.0


runtime error: index out of range [0] with length 0
/usr/lib/go/src/runtime/panic.go:115 (0x439b53)
        goPanicIndex: panic(boundsError{x: int64(x), signed: true, y: y, code: boundsIndex})
/workspaces/Course/Golang/Golang Backend/golang-backend/source-code/chapter4/4.1.2.TryRecoveryMidddleware.go:20 (0x735931)
        main.func2: _ = data[0] // This will cause panic!
/home/hilmimusyafa/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185 (0x72faae)
        (*Context).Next: c.handlers[c.index](c)
/home/hilmimusyafa/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/recovery.go:102 (0x72fa9b)
        CustomRecoveryWithWriter.func1: c.Next()
/home/hilmimusyafa/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185 (0x72ebe4)
        (*Context).Next: c.handlers[c.index](c)
/home/hilmimusyafa/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/logger.go:249 (0x72ebcb)
        LoggerWithConfig.func1: c.Next()
/home/hilmimusyafa/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185 (0x72dff1)
        (*Context).Next: c.handlers[c.index](c)
/home/hilmimusyafa/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/gin.go:633 (0x72da80)
        (*Engine).handleHTTPRequest: c.Next()
/home/hilmimusyafa/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/gin.go:589 (0x72d709)
        (*Engine).ServeHTTP: engine.handleHTTPRequest(c)
/usr/lib/go/src/net/http/server.go:3301 (0x62318d)
        serverHandler.ServeHTTP: handler.ServeHTTP(rw, req)
/usr/lib/go/src/net/http/server.go:2102 (0x6191c4)
        (*conn).serve: serverHandler{c.server}.ServeHTTP(w, w.req)
/usr/lib/go/src/runtime/asm_amd64.s:1700 (0x478ac0)
        goexit: BYTE    $0x90   // NOP

[GIN] 2025/07/17 - 22:18:37 | 500 |     1.86669ms |       127.0.0.1 | GET      "/cause-panic"
```

Di terminal tempat  menjalankan server, akan melihat _log_ yang mencatat _panic_ tersebut, lengkap dengan _stack trace_ (urutan fungsi yang dipanggil hingga terjadi _panic_). Namun, server itu sendiri **tidak akan mati**; ia akan terus berjalan dan siap melayani permintaan lainnya (misalnya, jika mengakses `/safe` lagi).

Selalu gunakan _Recovery middleware_! Ini adalah _middleware_ esensial untuk aplikasi produksi karena secara signifikan meningkatkan ketahanan aplikasi  terhadap _bug_ atau kondisi tak terduga yang bisa menyebabkan _panic_. Tanpanya, _panic_ sekecil apapun bisa membuat server Anda lumpuh. Dengan _Recovery middleware_,  mengetahui server akan tetap tegak meskipun ada _error_ yang tidak terduga.

Untuk mencoba kode bisa akses pada [4.1.2.TryRecoveryMidddleware.go](../../source-code/chapter4/4.1.2.TryRecoveryMidddleware.go)

### 4.1.3 CORS Middleware

**CORS** (_Cross-Origin Resource Sharing_) adalah mekanisme keamanan _browser_ yang mencegah halaman web membuat permintaan ke domain lain selain domain asalnya. Ini adalah fitur keamanan penting untuk melindungi pengguna dari serangan _Cross-Site Request Forgery_ (CSRF) dan serangan _web_ lainnya. 

Namun, dalam pengembangan API modern, terutama ketika _frontend_ (misalnya aplikasi React, Vue, Angular) berjalan di domain yang berbeda dengan _backend_ API Anda, CORS bisa menjadi hambatan.

_CORS middleware_ Gin hadir untuk menyelesaikan masalah ini dengan menambahkan _header HTTP_ yang diperlukan ke respons server, sehingga _browser_ klien mengizinkan permintaan lintas-asal.

Kira kira bagaiamana CORS bisa bekerja secara sederhana? Ketika sebuah _browser_ membuat permintaan lintas-asal (misalnya, JavaScript di `app.example.com` mencoba mengakses API di `api.example.com`), _browser_ akan mengirimkan _header_ `Origin` yang menunjukkan dari mana permintaan itu berasal. Server kemudian harus merespons dengan _header_ `Access-Control-Allow-Origin` yang mengindikasikan apakah `Origin` tersebut diizinkan untuk mengakses sumber daya.

Jika _Origin_ tidak diizinkan, _browser_ akan memblokir permintaan tersebut, bahkan jika server telah memprosesnya dan menghasilkan respons. 

Untuk implementasi CORS Middleware, Gin tidak memiliki _middleware_ CORS bawaan langsung seperti _Logger_ atau _Recovery_ di paket intinya. Namun, ada _middleware_ komunitas yang sangat populer dan direkomendasikan untuk CORS: `github.com/gin-contrib/cors`. Ini adalah _middleware_ yang sangat fleksibel dan mudah digunakan.

Tapi karena penting, jadi tetap akan di coba, dan untuk mulai pastikan melakukan Instalasi package `gin-contrib/cors` : 

```bash
$ go get github.com/gin-contrib/cors
```

Lalu setelah sudah install, bisa test dengan menjalankan kode di bawah :

4.1.3-TryCORSMiddleware.go

```go
package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors" // Import the CORS middleware
)

func main() {
	r := gin.Default()

	// CORS Middleware Configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "https://your-frontend-domain.com"} // Izinkan hanya domain ini
	// config.AllowAllOrigins = true // Atau izinkan semua domain (TIDAK disarankan untuk produksi)
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"} // Header yang boleh diekspos ke browser
	config.AllowCredentials = true // Izinkan pengiriman cookies dan header otorisasi
	config.MaxAge = 12 * time.Hour // Durasi hasil preflight request di-cache

	r.Use(cors.New(config)) // Terapkan CORS middleware

	// Endpoint API yang akan diakses dari frontend
	r.GET("/api/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Data from API"})
	})

	r.POST("/api/submit", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Data submitted successfully"})
	})

	r.Run(":8080")
}
```

Buat penjelasan Konfigurasi CORS di atas : 
- `cors.DefaultConfig()`: Mengembalikan konfigurasi CORS _default_ yang bisa dimodifikasi.
- `config.AllowOrigins`: Sangat penting untuk menentukan daftar domain yang diizinkan untuk mengakses API Anda. Menggunakan `AllowAllOrigins = true` harus dihindari di lingkungan produksi karena membuka API Anda untuk setiap domain, yang merupakan risiko keamanan.
- `config.AllowMethods`: Metode HTTP (GET, POST, dll.) yang diizinkan dari domain yang diizinkan.
- `config.AllowHeaders`: _Header_ kustom yang diizinkan untuk dikirim oleh klien. `Authorization` seringkali perlu ditambahkan jika Anda menggunakan _token_ autentikasi.
- `config.ExposeHeaders`: _Header_ yang diizinkan untuk diekspos ke _browser_ klien.
- `config.AllowCredentials`: Mengizinkan _browser_ untuk menyertakan _cookies_ atau _header_ otorisasi (`Authorization`) saat membuat permintaan lintas-asal. Ini penting untuk sesi berbasis _cookie_ atau _token_ otorisasi.
- `config.MaxAge`: Berapa lama hasil dari permintaan _preflight_ (permintaan `OPTIONS` yang dikirim _browser_ sebelum _request_ yang sebenarnya untuk memeriksa kebijakan CORS) dapat di-_cache_ oleh _browser_.

Dan untuk kapan sih penggunaan CORS Middleware? 

- Ketika memiliki _frontend_ aplikasi web (misalnya SPA berbasis JavaScript) yang disajikan dari domain/port yang berbeda dengan _backend_ Gin.
- Saat mengintegrasikan API dengan aplikasi pihak ketiga yang berjalan di domain berbeda.
    
Mengatur CORS dengan benar sangat penting untuk fungsionalitas API Anda sekaligus menjaga keamanan dasar terhadap permintaan lintas-asal yang tidak diinginkan.

Untuk mencoba bisa akses [4.1.3-TryCORSMiddleware.go](../../source-code/chapter4/4.1.3-TryCORSMiddleware.go)
### 4.1.4 Rate Limiting

## 4.2 Custom Middlewares

### 4.2.1 Authentication Middleware
### 4.2.2 Authorization Middleware
### 4.2.3 Request Logging dan Monitoring
### 4.2.4 Input Sanitization

## 4.3 Security Best Practices

### 4.3.1 JWT Authentication Implementation
### 4.3.2 Password Hashing (bcrypt)
### 4.3.3 Input Validation dan Sanitization
### 4.3.4 HTTPS dan Security Headers