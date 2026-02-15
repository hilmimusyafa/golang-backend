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

## 2.2 Basic Routing di Gin (HTTP Methods)

![Route Backend](../../images/2/2.2.0/2.2.0-1.png)

Bayangkan sebuah gedung perkantoran besar dengan satu pintu masuk utama. Di lobi, ada seorang resepsionis yang sangat cekatan. Setiap tamu yang datang pasti ditanya: "Mau kemana?" dan "Apa tujuannya?".

- Jika tamu membawa paket surat, resepsionis mengarahkan ke Ruang Surat.
- Jika tamu ingin bertemu manajer, diarahkan ke Ruang Rapat.
- Jika tamu ingin mendaftar kerja, diarahkan ke HRD.

Dalam Gin, Routing adalah resepsionis tersebut. Ia bertugas memetakan permintaan (request) yang masuk ke URL tertentu menuju ke fungsi penangan (handler) yang tepat berdasarkan metode HTTP yang digunakan.

Tanpa routing yang jelas, aplikasi tidak akan tahu harus berbuat apa saat pengguna mengakses `/users` atau `/products`.

Dalam standar REST API, setiap aksi memiliki metode HTTP yang spesifik. Menggunakan metode yang tepat membuat API mudah dipahami oleh developer lain (semantik).

Berikut adalah pemetaan standar metode HTTP dan kegunaannya :

### 2.2.1 GET Method

GET adalah metode yang paling sering digunakan. Ketika client mengirim permintaan GET, server akan mencari data yang diminta dan mengembalikannya. 

Metode ini bersifat aman (safe) karena tidak boleh mengubah kondisi data di server. Baik data ditemukan maupun tidak, tidak boleh ada data yang berubah akibat metode ini.

Bisa dilihat di kode ini : 

2.2.1-1-GetMethod.go

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	r := gin.Default()

	// 1. GET: Mengambil data
	r.GET("/books", func(c *gin.Context) {

		// Ambil data buku dari fungsi
		bookData := getBookData()

		c.JSON(http.StatusOK, gin.H{
			"message": "Showing list of books",
			"data":    bookData,
		})
	})

	r.Run(":8080")
}


// Fungsi simulasi mengambil data dari database
func getBookData() []Book {

	books := []Book{
		{ID: 1, Title: "Clean Code", Author: "Robert C. Martin"},
		{ID: 2, Title: "Refactoring", Author: "Martin Fowler"},
		{ID: 3, Title: "Domain-Driven Design", Author: "Eric Evans"},
	}

	return books
}
```

Penjelasan singkat bagian kode di atas seperti ini : 

- `type Book struct{}` : Agar data memiliki struktur yang jelas dan bisa dikonversi ke JSON.
- `gin.Default()` : Membuat instance router Gin yang sudah dilengkapi middleware bawaan seperti Logger (mencatat request ke terminal) dan Recovery (mencegah server crash jika terjadi panic)
- `r.GET("/books", handler)` : Mendaftarkan route dengan kelengkapan : 
	
	- Method : GET
	- Endpoint : `/books`
	- Handler : fungsi yang akan dijalankan saat endpoint diakses, 

	Jika ada request GET ke http://localhost:8080/books, maka fungsi di dalamnya akan dieksekusi.
- `c *gin.Context` : Objek Context digunakan untuk Mengambil parameter request, Mengakses body, Mengatur response. Di contoh ini, context digunakan untuk mengirim response JSON.
- `c.JSON(http.StatusOK, gin.H{...})` : Berguna untuk mengirim response informasi ke klien, jika kode di atas maka akan, berstatus code `200 OK` dengan Body berbentuk JSON. Dan sedangkan `gin.H` adalah shortcut untuk membuat map JSON.
- `r.Run(":8080")` : Menjalankan HTTP server pada port 8080. Tanpa baris ini, server tidak akan berjalan.
- `getBookData()` : Fungsi contoh asumsi dia mendapatkan data dari database.

> `getBookData()` merupakan fungsi asumsi untuk mendapatkan data dari database, fungsi ini fleksibel tergantuk developer, pemahaman ini ada di bab lain.

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
  "message": "Showing list of books",
  "data": [
    {
      "id": 1,
      "title": "Clean Code",
      "author": "Robert C. Martin"
    },
    {
      "id": 2,
      "title": "Refactoring",
      "author": "Martin Fowler"
    },
    {
      "id": 3,
      "title": "Domain-Driven Design",
      "author": "Eric Evans"
    }
  ]
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

Untuk mencoba kode di [2.2.1-1-GetMethod.go](../../codes/2/2.2.1/2.2.1-1-GetMethod.go)

### 2.2.2 POST Method

POST digunakan ketika client ingin menambahkan data baru ke server. Biasanya data dikirim melalui body request dalam format JSON. Berbeda dengan GET, metode POST tidak bersifat safe karena akan mengubah kondisi data di server (misalnya menambah entri baru di database).

Bisa dilihat di kode ini : 

2.2.2-1-PostMethod.go

```go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Simulasi database (in-memory)
var books = []Book{
	{ID: 1, Title: "Clean Code", Author: "Robert C. Martin"},
	{ID: 2, Title: "Refactoring", Author: "Martin Fowler"},
}

func main() {
	r := gin.Default()

	// 2. POST: Menambahkan data baru
	r.POST("/books", func(c *gin.Context) {

		var newBook Book

		// Bind JSON dari body request
		if err := c.ShouldBindJSON(&newBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON data",
			})
			return
		}

		// Generate ID sederhana (auto increment)
		newBook.ID = len(books) + 1

		// Tambahkan ke slice (simulasi insert database)
		books = append(books, newBook)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Book created successfully",
			"data":    newBook,
		})
	})

	r.Run(":8080")
}
```

Penjelasan singkat bagian kode di atas seperti ini :

- `var books = []Book{...}` : Berfungsi sebagai simulasi database sementara. Data disimpan dalam slice global.
- `r.POST("/books", handler)` : Mendaftarkan route dengan kelengkapan :
  
  - Method : POST
  - Endpoint : `/books`
  - Handler : fungsi yang akan dijalankan saat endpoint diakses

  Jika ada request POST ke http://localhost:8080/books, maka fungsi di dalamnya akan dieksekusi.

- `var book struct {...}` : Membuat struktur sementara untuk menampung data JSON yang dikirim client.
- `c.ShouldBindJSON(&book)` : Digunakan untuk membaca dan mengikat (bind) data JSON dari body request ke struct book. Jika format JSON tidak sesuai, maka akan mengembalikan error.
- `newBook.ID = len(books) + 1` : Karena belum pakai database, ID dibuat manual dengan auto increment sederhana.
- `books = append(books, newBook)` : Menambahkan data baru ke slice `books`. Ini mensimulasikan proses INSERT ke database.
- `http.StatusBadRequest (400)` : Digunakan jika JSON yang dikirim client tidak valid.
- `http.StatusCreated (201)` : Digunakan jika data berhasil dibuat.
Status ini adalah standar REST API untuk proses pembuatan resource baru.

> Note : Skenario `newBook.ID = len(books) + 1` khusus di kode ini, bagian ini dan pelengkapnya bisa di sesuaikan tergantun kode.

Dari kode di atas bisa dijalankan, akses di `curl` atau API tester : 

```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title":"Domain-Driven Design","author":"Eric Evans"}
```

Maka response jika berhasil yang dihasilkan adalah : 

```json
{
  "message": "Book created successfully",
  "data": {
    "id": 3,
    "title": "Domain-Driven Design",
    "author": "Eric Evans"
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

Kode bisa di akses di [2.2.2-1-PostMethod.go](../../codes/2/2.2.2/2.2.2-1-PostMethod.go)

### 2.2.3 PUT Method

Jika client ingin memperbarui data yang sudah ada, mereka bisa menggunakan PUT. Sifatnya adalah idempotent, artinya permintaan yang sama jika dikirim berulang kali akan menghasilkan efek yang sama. Penting untuk dicatat bahwa PUT biasanya mengharapkan client mengirimkan data yang lengkap untuk menggantikan data lama secara utuh.

Berbeda dengan POST yang membuat data baru, PUT biasanya digunakan untuk mengganti seluruh data pada resource yang sudah ada.

Bisa dilihat di kode ini : 

2.2.3-1-PutMethod.go

```go
package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Simulasi database (in-memory)
var books = []Book{
	{ID: 1, Title: "Clean Code", Author: "Robert C. Martin"},
	{ID: 2, Title: "Refactoring", Author: "Martin Fowler"},
}

func main() {
	r := gin.Default()

	// 3. PUT: Mengganti seluruh data
	r.PUT("/books/:id", func(c *gin.Context) {

		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid book ID",
			})
			return
		}

		var updatedBook Book

		// Bind JSON dari body request
		if err := c.ShouldBindJSON(&updatedBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON data",
			})
			return
		}

		// Panggil fungsi update terpisah
		book, found := updateBookByID(id, updatedBook)

		if !found {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Book not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Book updated successfully",
			"data":    book,
		})
	})

	r.Run(":8080")
}

// Fungsi Terpisah untuk Update Data
func updateBookByID(id int, updatedData Book) (Book, bool) {

	for i, book := range books {
		if book.ID == id {
			updatedData.ID = id // Pastikan ID tetap sama
			books[i] = updatedData
			return updatedData, true
		}
	}

	return Book{}, false
}

```

Penjelasan singkat bagian kode di atas seperti ini :

- `r.PUT("/books/:id", handler)` : Mendaftarkan route dengan kelengkapan :
  - Method : PUT
  - Endpoint : `/books/:id`, `:id` adalah route parameter
- Handler : fungsi yang akan dijalankan saat endpoint diakses, Jika ada request PUT ke http://localhost:8080/books/1, maka fungsi di dalamnya akan dieksekusi.
- `c.Param("id")` : Digunakan untuk mengambil nilai parameter id dari URL.
- `strconv.Atoi(idParam)` : Mengubah ID dari string menjadi integer karena parameter URL selalu berbentuk string.
- `c.ShouldBindJSON(&book)` : Membaca dan mengikat data JSON dari body request ke struct book. Jika format JSON tidak valid, maka akan mengembalikan error.
- `updateBookByID(id, updatedBook)` : Fungsi terpisah yang menangani logika update.

> Catatan : Fungsi `updateBookByID(id, updatedBook)` dinamis tergantung kegunaan kode.

Kita coba jalankan. Untuk menguji PUT, kirim body JSON lengkap karena PUT biasanya mengganti seluruh data.

```bash
$ curl -X PUT http://localhost:8080/books/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"New Clean Code","author":"Robert Martin"}'
```

Maka response yang dihasilkan jika berhasil :

```json
{
  "message": "Book updated successfully",
  "data": {
    "id": 1,
    "title": "New Clean Code",
    "author": "Robert Martin"
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

Kita juga coba dengan ID yang tidak ada : 

```bash
$ curl -X PUT http://localhost:8080/books/99 \
  -H "Content-Type: application/json" \
  -d '{"title":"Test","author":"Test"}'
```

Maka hasilnya adalah : 

```json
{
  "error": "Book not found"
}
```

Dengan status : 

```
404 Not Found
```

Dari contoh di atas, dapat disimpulkan bahwa ketika client mengakses endpoint `/books/:id` menggunakan metode PUT, Gin akan membaca parameter dari URL, mengambil data dari body request, lalu mengganti data sesuai dengan informasi yang dikirim.

Karena PUT bersifat idempotent, pengiriman request yang sama berulang kali tidak akan mengubah hasil akhir dari resource tersebut.

Kode bisa di akses di [2.2.3-1-PutMethod.go](../../codes/2/2.2.3/2.2.3-1-PutMethod.go)

### 2.2.4 PATCH Method

Berbeda dengan PUT yang mengganti seluruh data, PATCH digunakan untuk modifikasi parsial. Client hanya perlu mengirimkan field atau kolom mana yang ingin diubah. Ini lebih efisien untuk pembaruan data yang tidak terlalu besar. PATCH umumnya tidak selalu idempotent, tergantung implementasinya.

Bisa dilihat di kode ini : 

2.2.4-1-PatchMethod.go

```go
package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Simulasi database (in-memory)
var books = []Book{
	{ID: 1, Title: "Clean Code", Author: "Robert C. Martin"},
	{ID: 2, Title: "Refactoring", Author: "Martin Fowler"},
}

// Struct khusus untuk PATCH (pakai pointer agar bisa deteksi field yang dikirim)
type UpdateBookInput struct {
	Title  *string `json:"title"`
	Author *string `json:"author"`
}

func main() {
	r := gin.Default()

	// 4. PATCH: Update sebagian data
	r.PATCH("/books/:id", func(c *gin.Context) {

		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid book ID",
			})
			return
		}

		var input UpdateBookInput

		// Bind JSON parsial
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON data",
			})
			return
		}

		// Panggil fungsi update parsial terpisah
		book, found := patchBookByID(id, input)

		if !found {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Book not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Book partially updated",
			"data":    book,
		})
	})

	r.Run(":8080")
}

// Fungsi Terpisah untuk PATCH
func patchBookByID(id int, input UpdateBookInput) (Book, bool) {

	for i, book := range books {

		if book.ID == id {

			// Update hanya field yang dikirim
			if input.Title != nil {
				books[i].Title = *input.Title
			}

			if input.Author != nil {
				books[i].Author = *input.Author
			}

			return books[i], true
		}
	}

	return Book{}, false
}

```

Penjelasan singkat bagian kode di atas seperti ini :

- `r.PATCH("/books/:id", handler)` : Mendaftarkan route dengan kelengkapan :
  
  - Method : PATCH
  - Endpoint : `/books/:id`, `:id` adalah route parameter
- `Struct UpdateBookInput`: struct khusus untuk inputan.
- Handler : fungsi yang akan dijalankan saat endpoint diakses. Jika ada request PATCH ke http://localhost:8080/books/1, maka fungsi di dalamnya akan dieksekusi.
- `c.Param("id")` : Digunakan untuk mengambil nilai parameter id dari URL.
- `Title *string` dan `Author *string` : Menggunakan pointer agar dapat membedakan antara :

  - Field yang tidak dikirim
  - Field yang dikirim tetapi bernilai kosong

  Ini penting dalam PATCH karena kita hanya ingin mengubah field yang benar-benar dikirim oleh client.

Karena PATCH bersifat parsial, kita cukup mengirim field yang ingin diubah. Contoh hanya mengubah title :

```bash
$ curl -X PATCH http://localhost:8080/books/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Clean Code 2nd Edition"}'
```

Jika berhasil, server akan mengembalikan:

```json
{
  "message": "Book partially updated",
  "data": {
    "id": 1,
    "title": "Clean Code 2nd Edition",
    "author": "Robert C. Martin"
  }
}
```

Dengan status HTTP :

```
200 OK
```

Namun, jika body JSON tidak valid:

```bash
$ curl -X PATCH http://localhost:8080/books/1
```

Maka response yang didapat :

```json
{
  "error": "Book not found"
}
```

Dengan status HTTP :

```
400 Bad Request
```
Coba juga dengan 

```bash
$ curl -X PATCH http://localhost:8080/books/67 \
  -H "Content-Type: application/json" \
  -d '{"title":"Clean Code 2nd Edition"}'
```

Maka responnya adalah : 

```json
{
  "error": "Book not found"
}
```

Dengan status HTTP :

```
404 Not Found
```


Dari contoh di atas, dapat disimpulkan bahwa ketika client mengakses endpoint `/books/:id` menggunakan metode PATCH, Gin akan membaca parameter dari URL dan hanya memperbarui field yang dikirim oleh client.

PATCH digunakan untuk update sebagian data dan lebih efisien dibandingkan PUT ketika perubahan tidak mencakup seluruh resource.

Bisa diakses kode di [2.2.4-1-PatchMethod](../../codes/2/2.2.4/2.2.4-1-PatchMethod.go)

### 2.2.5 DELETE Method

Sesuai namanya, metode ini digunakan untuk menghapus data yang ditentukan dari server. Sama seperti PUT, DELETE juga bersifat idempotent. Permintaan yang sama akan tetap menghasilkan data tersebut telah hilang (meskipun responsnya bisa berbeda, misalnya `200 OK` untuk penghapusan pertama dan `404 Not Found` untuk percobaan kedua). DELETE biasanya tidak memerlukan body request, cukup dengan menentukan resource melalui URL.

Bisa dilihat di kode ini : 

2.2.5-1-DeleteMethod.go

```go
package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Simulasi database (in-memory)
var books = []Book{
	{ID: 1, Title: "Clean Code", Author: "Robert C. Martin"},
	{ID: 2, Title: "Refactoring", Author: "Martin Fowler"},
}

func main() {
	r := gin.Default()

	// 5. DELETE: Menghapus data
	r.DELETE("/books/:id", func(c *gin.Context) {

		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid book ID",
			})
			return
		}

		// Panggil fungsi delete terpisah
		deletedBook, found := deleteBookByID(id)

		if !found {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Book not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Book deleted successfully",
			"data":    deletedBook,
		})
	})

	r.Run(":8080")
}

// Fungsi untuk hapus
func deleteBookByID(id int) (Book, bool) {

	for i, book := range books {
		if book.ID == id {
			// Simpan data yang akan dihapus
			deletedBook := book

			// Hapus dari slice
			books = append(books[:i], books[i+1:]...)

			return deletedBook, true
		}
	}
	return Book{}, false
}
```

Penjelasan dari kode di atas :

- `r.DELETE("/books/:id", handler)` : Mendaftarkan route dengan kelengkapan :
   
   - Method : DELETE
   - Endpoint : `/books/:id`, `:id` adalah route parameter
   - Handler : fungsi yang akan dijalankan saat endpoint diakses
  
  Jika ada request DELETE ke http://localhost:8080/books/1, maka fungsi di dalamnya akan dieksekusi.

- Fungsi `deleteBookByID` : Untuk fungsi hapus data.

Uji test dengan `curl` atau API Tester : 

```bash
$ curl -X DELETE http://localhost:8080/books/1
```

Jika berhasil maka akan mengeluarkan informasi :

```json
{
  "message": "Book deleted successfully",
  "data": {
    "id": 1,
    "title": "Clean Code",
    "author": "Robert C. Martin"
  }
}
```

Status :

```
200 OK
```

Namun jika ID tidak ditemukan :

```json
{
  "error": "Book not found"
}
```

Dengan status : 

```
404 Not Found
```

Dari contoh di atas, dapat disimpulkan bahwa ketika client mengakses endpoint `/books/:id` menggunakan metode DELETE, Gin akan mengambil parameter dari URL dan menghapus resource yang dimaksud.

Karena DELETE bersifat idempotent, pengiriman request yang sama berulang kali tidak akan mengubah hasil akhir yaitu resource tersebut tetap dalam keadaan terhapus.

Kode bisa diakses di [2.2.5-1-DeleteMethod.go](../../codes/2/2.2.5/2.2.5-1-DeleteMethod.go)

## 2.3 Request Handling

Dalam pengembangan backend, salah satu tugas utama adalah menerima data dari klien (frontend, mobile app, atau layanan lain). Data ini bisa berupa JSON, form data, atau file upload.

Mengapa kita butuh mekanisme Request Handling khusus? Bayangkan jika harus memparsing teks JSON mentah (`{"name": "Budi", "age": 20}`) secara manual menjadi variabel Golang. Pasti melelahkan dan rentan error. Gin menyediakan fitur Binding yang bertindak sebagai penerjemah otomatis. Ia mengubah data mentah dari HTTP request langsung menjadi struktur data (Struct) Golang yang aman dan siap pakai.

### 2.3.1 Binding Data & Validasi (JSON, Form, Query)

Konsep dasarnya sederhana adalah seperti menyiapkan wadah (Struct) dengan label (Tag), lalu suruh Gin mengisinya.

#### 2.3.1.1 Struct Tags

Agar Gin tahu field mana yang harus diisi, gunakan tag pada struct.

- `json:"fieldname"` → Untuk data dari Body (Raw JSON).
- `form:"fieldname"` → Untuk data dari Form Data atau Query Param.
- `binding:"required"` → Untuk validasi (wajib diisi).

#### 2.3.1.2 ShouldBind vs MustBind 

Gin memiliki dua keluarga method untuk binding :

- Type Bind (MustBind) : Contohnya `c.BindJSON()`. Jika validasi gagal (misal field required kosong), Gin akan otomatis menghentikan request, mengirim status 400, dan menulis header response. Anda tidak bisa mengontrol format error-nya.
- Type ShouldBind : Contohnya `c.ShouldBindJSON()`. Jika validasi gagal, ia akan mengembalikan error dan membiarkan programmer yang menentukan penanganan error tersebut (misalnya membungkusnya dalam format JSON standar API Anda).

> Rekomendasi : Selalu gunakan ShouldBind di production agar Anda punya kontrol penuh terhadap error handling.

Saatnya buat contoh implementasi yang mencakup Binding JSON, Validasi, dan Query Param : 

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRegister mendefinisikan struktur data yang diharapkan dari client
type UserRegister struct {
	// binding:"required" memastikan field ini tidak boleh kosong
	Username string `json:"username" binding:"required"`
	
	// binding:"email" memvalidasi format email
	Email    string `json:"email" binding:"required,email"`
	
	// binding:"min=8" memastikan panjang minimal 8 karakter
	Password string `json:"password" binding:"required,min=8"`
	
	// Field optional, tidak perlu binding:"required"
	Age      int    `json:"age"` 
}

func main() {
	r := gin.Default()

	r.POST("/register", func(c *gin.Context) {
		var input UserRegister

		// 1. Lakukan Binding & Validasi
		// ShouldBindJSON akan membaca Body request dan mencocokkan dengan struct UserRegister
		if err := c.ShouldBindJSON(&input); err != nil {
			// Jika validasi gagal, kembalikan error 400 Bad Request
			// gin.H{"error": err.Error()} akan menampilkan detail error validasi
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 2. Simulasi Proses Data (misal simpan ke database)
		// Di sini input.Username dll sudah aman digunakan

		c.JSON(http.StatusOK, gin.H{
			"message":  "User berhasil didaftarkan",
			"username": input.Username,
			"email":    input.Email,
		})
	})

	// Contoh Binding Query String (misal: /search?q=golang&page=1)
	r.GET("/search", func(c *gin.Context) {
		var query struct {
			Q    string `form:"q" binding:"required"` // Gunakan tag 'form' untuk query param
			Page int    `form:"page"`
		}

		// ShouldBindQuery khusus untuk mengambil parameter dari URL
		if err := c.ShouldBindQuery(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"search_query": query.Q,
			"page":         query.Page,
		})
	})

	r.Run(":8080")
}
```