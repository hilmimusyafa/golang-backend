# Bab 2 : Gin Framework Fundamentals

## 2.1 Mengenal Gin

### 2.1.1 Apa itu Gin Framework

Dalam dunia pengembangan backend menggunakan Golang, Gin (atau sering disebut Gin-Gonic) adalah salah satu web framework yang paling populer dan banyak digunakan. Gin dikenal karena kecepatannya, kesederhanaannya, dan performanya yang tinggi.

![Gin Gonic Framework](../../images/2/2.1.1/2.1.1-1.png)

Mengapa perlu framework? Bukankah Go sudah punya `net/http`?

Bayangkan `net/http` itu seperti membangun rumah dengan memotong kayu dan mencampur semen sendiri. Punya kontrol penuh, tapi butuh waktu lama dan rawan kesalahan untuk fitur dasar seperti routing kompleks atau middleware. Gin hadir seperti kontraktor yang membawa peralatan canggih dan modul siap pasang. Gin dibangun di atas `net/http`, tapi mempermudah hal-hal yang rumit sehingga pengembangan aplikasi menjadi jauh lebih cepat dan terstruktur.

### 2.1.2 Kenapa Memilih Gin?

Ada tiga alasan utama kenapa Gin menjadi pilihan de facto bagi banyak perusahaan dan developer Go :

- Performa Tinggi : Gin menggunakan `httprouter`, sebuah multiplexer HTTP kustom yang sangat cepat. Gin mengklaim dirinya 40x lebih cepat daripada framework sejenis lainnya (seperti Martini).
- Popularitas & Ekosistem : Karena sangat populer, mencari solusi saat stuck atau mencari library pendukung (seperti CORS, Gzip, Auth) sangat mudah. Dukungan komunitasnya masif.
- Kesederhanaan (Simplicity) : API yang ditawarkan Gin sangat intuitif. Menangani request JSON, validasi data, dan routing bisa dilakukan dengan kode yang sangat ringkas.

### 2.1.3 Perbandingan : Gin vs Echo vs Fiber vs `net/http`

Agar lebih yakin, mari lihat perbandingan singkatnya :

| Framework   | Kelebihan Utama                                                                                              | Kekurangan / Catatan                                                                                                 |
| :---------- | :----------------------------------------------------------------------------------------------------------- | :------------------------------------------------------------------------------------------------------------------- |
| **`net/http`**  | Standar bawaan Go. Stabil, tidak ada dependensi eksternal.                                                   | Kode menjadi sangat panjang (verbose) untuk fitur standar seperti dynamic routing atau middleware chains.            |
| **Gin**       | Keseimbangan sempurna antara performa, kemudahan, dan fitur. Ekosistem terbesar.                            | Sedikit lebih lambat dari Fiber (meski bedanya dalam nanodetik, jarang terasa di aplikasi bisnis umum).             |
| **Echo**      | Sangat mirip dengan Gin, dokumentasi sangat rapi.                                                            | Komunitasnya sedikit lebih kecil dibanding Gin, meski secara fitur sangat bersaing.                                  |
| **Fiber**     | Tercepat (berbasis `fasthttp`, bukan `net/http`). Sangat mirip Express.js.                                       | Tidak kompatibel 100% dengan `net/http`. Ini bisa menyulitkan jika ingin menggunakan library standar Go lain yang butuh kompatibilitas `net/http`. |

Untuk proyek backend profesional yang membutuhkan kestabilan, kompatibilitas luas, dan kemudahan maintenance, Gin adalah pilihan yang sangat aman dan solid.

### 2.1.4 Instalasi dan Setup Gin

Sebelum mulai, pastikan Go sudah terinstal di komputer. Langkah pertama dalam setiap proyek Go modern adalah inisialisasi modul.

Buka terminal dan jalankan perintah berikut untuk membuat folder proyek dan inisialisasi `go.mod` :

```bash
$ cd <folder>
$ go mod init <folder>
```

Setelah itu, unduh package Gin :

```bash
$ go get -u github.com/gin-gonic/gin
```

Perintah ini akan mengunduh Gin dan dependensinya, serta mencatatnya di file `go.mod` dan `go.sum`.

### 2.1.5 First Gin Application (Hello World)

Tujuan aplikasi pertama ini sederhana, yaitu membuat server HTTP yang merespons dengan format JSON ketika diakses.

Mengapa JSON? Karena dalam konteks Backend Developer, hampir 90% pekerjaan adalah membuat REST API yang berkomunikasi menggunakan JSON, bukan merender HTML.

Silakan buat file baru dan jalankan kode berikut :

2.1.5-1-SimpleServer.go

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Inisialisasi Router Gin
	// gin.Default() membuat instance router dengan middleware bawaan:
	// - Logger (mencatat log request ke console)
	// - Recovery (mencegah server crash/panic jika ada error fatal)
	r := gin.Default()

	// 2. Definisi Route
	// Ketika user akses GET ke root url "/", jalankan fungsi ini
	r.GET("/", func(c *gin.Context) {
		// c.JSON akan mengubah map/struct menjadi format JSON secara otomatis
		// dan mengatur Content-Type menjadi application/json
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello! This is my first Gin Application",
			"status":  "success",
		})
	})

	// Route tambahan dengan parameter
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 3. Menjalankan Server
	// Secara default akan berjalan di port 8080
	// Bisa diubah dengan r.Run(":3000")
	r.Run() 
}
```

Setelah kode disimpan, jalankan lewat terminal :

```bash
$ go run 2.1.5-1-SimpleServer.go
```

Output di terminal akan terlihat seperti ini (menandakan server aktif) :

```
[GIN-debug] Listening and serving HTTP on :8080
```

Sekarang, buka browser atau Postman dan akses http://localhost:8080. Hasilnya akan berupa JSON :

```json
{
  "message": "Hello! This is my first Gin Application",
  "status": "success"
}
```

Mari bedah bagian-bagian vital dari kode di atas :

- `gin.Default()` Ini adalah constructor standar. Gin sebenarnya punya `gin.New()` yang membuat router kosong tanpa middleware apa pun. Namun, `gin.Default()` sangat disarankan untuk pemula karena sudah menyertakan Logger (untuk melihat trafik di terminal) dan Recovery (agar satu bug fatal tidak mematikan seluruh server).
- `func(c *gin.Context)`: Ini adalah handler function. `*gin.Context`    adalah bagian paling penting di Gin. Objek c ini membawa semua informasi tentang :

    - Request (parameter URL, body JSON, header, cookies).
    - Response (cara mengirim balik data, validasi, kode status).
    - Menggantikan `http.ResponseWriter` dan `*http.Request` di `net/http` standar.

- `gin.H{}` : Ini hanyalah shortcut atau alias untuk `map[string]interface`{}. Daripada menulis tipe data yang panjang untuk membuat objek JSON sederhana, Gin menyediakan `gin.H` agar penulisan kode lebih rapi dan cepat.
- `r.Run()`: Metode ini memblokir eksekusi program dan mulai mendengarkan request HTTP. Jika tidak diberi parameter, ia akan menggunakan port `:8080` secara default.

Dengan pemahaman dasar ini, fondasi untuk membangun API yang lebih kompleks sudah terbentuk. Selanjutnya, materi akan membahas bagaimana cara menangani berbagai jenis request method dan memproses data yang dikirim oleh pengguna.

Untuk mencoba kodenya, silakan akses : [2.1.5-1-SimpleServer.go](2.1-SimpleServer.go)

## 2.2 Routing di Gin

![Route Backend](../../images/2/2.2.0/2.2.0-1.png)

Bayangkan sebuah gedung perkantoran besar dengan satu pintu masuk utama. Di lobi, ada seorang resepsionis yang sangat cekatan. Setiap tamu yang datang pasti ditanya: "Mau kemana?" dan "Apa tujuannya?".

- Jika tamu membawa paket surat, resepsionis mengarahkan ke Ruang Surat.
- Jika tamu ingin bertemu manajer, diarahkan ke Ruang Rapat.
- Jika tamu ingin mendaftar kerja, diarahkan ke HRD.

Dalam Gin, Routing adalah resepsionis tersebut. Ia bertugas memetakan permintaan (request) yang masuk ke URL tertentu menuju ke fungsi penangan (handler) yang tepat berdasarkan metode HTTP yang digunakan.

Tanpa routing yang jelas, aplikasi tidak akan tahu harus berbuat apa saat pengguna mengakses `/users` atau `/products`.

### 2.2.1 Basic Routing (HTTP Methods)

Dalam standar REST API, setiap aksi memiliki metode HTTP yang spesifik. Menggunakan metode yang tepat membuat API mudah dipahami oleh developer lain (semantik).

Berikut adalah pemetaan standar metode HTTP dan kegunaannya :

#### 2.2.1.1 GET Method

GET adalah metode yang paling sering digunakan. Ketika client mengirim permintaan GET, server akan mencari data yang diminta dan mengembalikannya. 

Metode ini bersifat aman (safe) karena tidak boleh mengubah kondisi data di server. Baik data ditemukan maupun tidak, tidak boleh ada data yang berubah akibat metode ini.

Bisa dilihat di kode ini : 

2.2.1.1-1-GetMethod.go

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 1. GET: Mengambil data
	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Showing list of books",
		})
	})
	
	r.Run(":8080") // menjalankan server di port 8080
}
```

Penjelasan singkat bagian kode di atas seperti ini : 

- `gin.Default()` : Membuat instance router Gin yang sudah dilengkapi middleware bawaan seperti Logger (mencatat request ke terminal) dan Recovery (mencegah server crash jika terjadi panic)
- `r.GET("/books", handler)` : Mendaftarkan route dengan kelengkapan : 
	
	- Method : GET
	- Endpoint : `/books`
	- Handler : fungsi yang akan dijalankan saat endpoint diakses, 

	Jika ada request GET ke http://localhost:8080/books, maka fungsi di dalamnya akan dieksekusi.
- `c *gin.Context` : Objek Context digunakan untuk Mengambil parameter request, Mengakses body, Mengatur response. Di contoh ini, context digunakan untuk mengirim response JSON.
- `c.JSON(http.StatusOK, gin.H{...})` : Berguna untuk mengirim response informasi ke klien, jika kode di atas maka akan, berstatus code `200 OK` dengan Body berbentuk JSON. Dan sedangkan `gin.H` adalah shortcut untuk membuat map JSON.
- `r.Run(":8080")` : Menjalankan HTTP server pada port 8080. Tanpa baris ini, server tidak akan berjalan.

Dari kode di atas bisa dijalankan, akses di `curl` atau API tester : 

```bash
$ curl http://localhost:8080/books
```

Atau lebih eksplisit dengan metode GET :

```bash
$ curl -X GET http://localhost:8080/books
```

Maka server akan mengembalikan response :

```json
{
  "message": "Showing list of books"
}
```

Dengan status HTTP :

```
200 OK
```

Dari contoh di atas, dapat disimpulkan bahwa ketika client mengakses endpoint `/books` menggunakan metode GET, Gin akan menjalankan handler yang sesuai dan mengembalikan data dalam format JSON tanpa mengubah kondisi data di server.

Namun, jika di coba lagi dengan sembarang routing, misal `/unkown` :

```bash
$ curl http://localhost:8080/unknown
```

Maka response yang di dapat pada API tester :

```
404 page not found
```

Ini jelas, jika route tidak terdaftar, Gin akan mengembalikan 404 Not Found. Ini menunjukkan pentingnya mendefinisikan routing dengan benar.

Untuk mencoba kode di [2.2.1.1-1-GetMethod.go](../../codes/2/2.2.1.1/2.2.1.1-1-GetMethod.go)

#### 2.2.1.2 POST Method

POST digunakan ketika client ingin menambahkan data baru ke server. Biasanya data dikirim melalui body request dalam format JSON. Berbeda dengan GET, metode POST tidak bersifat safe karena akan mengubah kondisi data di server (misalnya menambah entri baru di database).

Bisa dilihat di kode ini : 

2.2.1.2-1-PostMethod.go

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 2. POST: Menambahkan data baru
	r.POST("/books", func(c *gin.Context) {
		var book struct {
			Title  string `json:"title"`
			Author string `json:"author"`
		}

		// Mengambil data JSON dari body request
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON data",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Book created successfully",
			"data":    book,
		})
	})
	
	r.Run(":8080")
}
```

Penjelasan singkat bagian kode di atas seperti ini :

- `r.POST("/books", handler)` : Mendaftarkan route dengan kelengkapan :
  
  - Method : POST
  - Endpoint : `/books`
  - Handler : fungsi yang akan dijalankan saat endpoint diakses

  Jika ada request POST ke http://localhost:8080/books, maka fungsi di dalamnya akan dieksekusi.

- `var book struct {...}` : Membuat struktur sementara untuk menampung data JSON yang dikirim client.
- `c.ShouldBindJSON(&book)` : Digunakan untuk membaca dan mengikat (bind) data JSON dari body request ke struct book. Jika format JSON tidak sesuai, maka akan mengembalikan error.
- `http.StatusBadRequest (400)` : Digunakan jika JSON yang dikirim client tidak valid.
- `http.StatusCreated (201)` : Digunakan jika data berhasil dibuat.
Status ini adalah standar REST API untuk proses pembuatan resource baru.

Dari kode di atas bisa dijalankan, akses di `curl` atau API tester : 

```bash
$ curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title":"Clean Code","author":"Robert C. Martin"}'
```

Maka response jika berhasil yang dihasilkan adalah : 

```json
{
  "message": "Book created successfully",
  "data": {
    "title": "Clean Code",
    "author": "Robert C. Martin"
  }
}
```

Dengan status HTTP :

```
201 Created
```

Namun, jika body JSON tidak valid atau tidak dikirim :

```bash
$ curl -X POST http://localhost:8080/books
```

Maka response yang didapat :

```json
{
  "error": "Invalid JSON data"
}
```

Dengan status HTTP :

```
400 Bad Request
```

Dari contoh di atas, dapat disimpulkan bahwa ketika client mengakses endpoint `/books` menggunakan metode POST, Gin akan membaca data dari body request, memvalidasi format JSON, lalu mengembalikan response sesuai hasil proses tersebut.

Berbeda dengan GET, metode POST digunakan untuk membuat resource baru dan akan mengubah kondisi data di server.

Kode bisa di akses di []()

#### 2.2.1.3 PUT Method

Jika client ingin memperbarui data yang sudah ada, mereka bisa menggunakan PUT. Sifatnya adalah idempotent, artinya permintaan yang sama jika dikirim berulang kali akan menghasilkan efek yang sama. Penting untuk dicatat bahwa PUT biasanya mengharapkan client mengirimkan data yang lengkap untuk menggantikan data lama secara utuh.

Berbeda dengan POST yang membuat data baru, PUT biasanya digunakan untuk mengganti seluruh data pada resource yang sudah ada.

Bisa dilihat di kode ini : 

2.2.1.3-1-PutMethod.go

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 3. PUT: Mengganti seluruh data
	r.PUT("/books/:id", func(c *gin.Context) {
		id := c.Param("id")

		var book struct {
			Title  string `json:"title"`
			Author string `json:"author"`
		}

		// Mengambil data JSON dari body request
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON data",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Book with id " + id + " updated successfully",
			"data":    book,
		})
	})
	
	r.Run(":8080")
}
```

Penjelasan singkat bagian kode di atas seperti ini :

- `r.PUT("/books/:id", handler)` : Mendaftarkan route dengan kelengkapan :
  - Method : PUT
  - Endpoint : `/books/:id`, `:id` adalah route parameter
- Handler : fungsi yang akan dijalankan saat endpoint diakses, Jika ada request PUT ke http://localhost:8080/books/1, maka fungsi di dalamnya akan dieksekusi.
- `c.Param("id")` : Digunakan untuk mengambil nilai parameter id dari URL.
- `c.ShouldBindJSON(&book)` : Membaca dan mengikat data JSON dari body request ke struct book. Jika format JSON tidak valid, maka akan mengembalikan error.

Kita coba jalankan. Untuk menguji PUT, kirim body JSON lengkap karena PUT biasanya mengganti seluruh data.

```bash
$ curl -X PUT http://localhost:8080/books/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Refactoring","author":"Martin Fowler"}'
```

Maka response yang dihasilkan jika berhasil :

```json
{
  "message": "Book with id 1 updated successfully",
  "data": {
    "title": "Refactoring",
    "author": "Martin Fowler"
  }
}
```

Dengan status HTTP :

```
200 OK
```

Namun, jika body JSON tidak valid :

```bash
$ curl -X PUT http://localhost:8080/books/1
```

Maka response yang didapat :

```json
{
  "error": "Invalid JSON data"
}
```

Dengan status HTTP :

```
400 Bad Request
```

Dari contoh di atas, dapat disimpulkan bahwa ketika client mengakses endpoint `/books/:id` menggunakan metode PUT, Gin akan membaca parameter dari URL, mengambil data dari body request, lalu mengganti data sesuai dengan informasi yang dikirim.

Karena PUT bersifat idempotent, pengiriman request yang sama berulang kali tidak akan mengubah hasil akhir dari resource tersebut.

Kode bisa di akses di []()

#### 2.2.1.4 PATCH Method

Berbeda dengan PUT yang mengganti seluruh data, PATCH digunakan untuk modifikasi parsial. Client hanya perlu mengirimkan field atau kolom mana yang ingin diubah. Ini lebih efisien untuk pembaruan data yang tidak terlalu besar. PATCH umumnya tidak selalu idempotent, tergantung implementasinya.

Bisa dilihat di kode ini : 

2.2.1.4-1-PatchMethod.go

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 4. PATCH: Mengubah sebagian data
	r.PATCH("/books/:id", func(c *gin.Context) {
		id := c.Param("id")

		var book struct {
			Title  *string `json:"title"`
			Author *string `json:"author"`
		}

		// Mengambil data JSON dari body request
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON data",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Book with id " + id + " partially updated",
			"data":    book,
		})
	})
	
	r.Run(":8080")
}
```

Penjelasan singkat bagian kode di atas seperti ini :

- `r.PATCH("/books/:id", handler)` : Mendaftarkan route dengan kelengkapan :
  
  - Method : PATCH
  - Endpoint : `/books/:id`, `:id` adalah route parameter

- Handler : fungsi yang akan dijalankan saat endpoint diakses. Jika ada request PATCH ke http://localhost:8080/books/1, maka fungsi di dalamnya akan dieksekusi.
- `c.Param("id")` : Digunakan untuk mengambil nilai parameter id dari URL.
- `Title *string` dan `Author *string` : Menggunakan pointer agar dapat membedakan antara :

  - Field yang tidak dikirim
  - Field yang dikirim tetapi bernilai kosong

  Ini penting dalam PATCH karena kita hanya ingin mengubah field yang benar-benar dikirim oleh client.

#### 2.2.1.5 DELETE Method

Sesuai namanya, metode ini digunakan untuk menghapus data yang ditentukan dari server. Sama seperti PUT, DELETE juga bersifat idempotent. Permintaan yang sama akan tetap menghasilkan data tersebut telah hilang (meskipun responsnya bisa berbeda, misalnya `200 OK` untuk penghapusan pertama dan `404 Not Found` untuk percobaan kedua).

Bisa dilihat di kode ini : 

2.2.1.1-1-GetMethod.go

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

```