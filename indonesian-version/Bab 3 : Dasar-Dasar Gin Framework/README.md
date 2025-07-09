# Bab 3 : Dasar-Dasar Gin Framework

## 3.1 Bagian - Bagian Dasar Framework Gin

Sebelum masuk ke detail, ada baiknya mehami dulu bagian-bagian penting dari kode Gin yang sering ditemukan saat memulai proyek :

### 3.1.1 Import Package Gin

```go
import (
	"net/http"
	"github.com/gin-gonic/gin"
)
```
Berikut penjelasan singkat :

- `net/hhtp` :  Digunakan untuk konstanta status HTTP seperti http.StatusOK
- `github.com/gin-gonic/gin` : Import package utama framework Gin
  
Setiap aplikasi Go memerlukan package yang relevan untuk fungsionalitasnya. `net/http` adalah package standar Go yang menyediakan fungsi dasar untuk protokol HTTP, termasuk kode status (misalnya, 200 OK, 404 Not Found). 

Sementara itu, `github.com/gin-gonic/gin` adalah package Gin itu sendiri, yang berisi semua tools dan fungsi yang diperlukan untuk membangun API dengan Gin.

### 3.1.2 Inisialisasi Gin Engine

```go
func main() {
	r := gin.Default()
	
    // ... kode routing lainnya ...
}
```

Baris `r := gin.Default()` adalah titik awal aplikasi Gin, dengan keterangan berikut :
1. `gin.Default()` : Ini adalah fungsi yang mengembalikan instance dari `*gin.Engine`. `*gin.Engine` adalah objek utama yang akan digunakan untuk mendefinisikan rute, middleware, dan menjalankan server. `Default()` secara otomatis menyertakan dua middleware bawaan yang sangat berguna :
   - Logger : Menampilkan log dari setiap permintaan yang masuk ke konsol, sangat membantu untuk debugging.
    - Recovery : Menangkap panic (kesalahan runtime) yang mungkin terjadi selama pemrosesan permintaan, sehingga server tidak crash sepenuhnya dan bisa mengirimkan respons error yang sesuai ke klien.
2. Jika ingin memulai dengan engine yang benar-benar "kosong" tanpa middleware bawaan, bisa gunakan `r := gin.New()`. Namun, `gin.Default()` adalah pilihan paling umum dan direkomendasikan untuk sebagian besar kasus.
3. `r`: Ini adalah variabel (umumnya dinamakan r untuk router) yang menyimpan instance `*gin.Engine`. Semua definisi rute dan konfigurasi server akan melekat pada objek `r` ini.

### 3.2.3 Menjalankan Server

```go
func main() {
	// ... kode inisialisasi dan routing ...
	r.Run(":8080")
}
```

`r.Run(":8080")` adalah baris terakhir yang harus dipanggil di fungsi main Gin. Fungsi ini :

- Memulai server HTTP Gin.
- endengarkan permintaan masuk pada alamat dan port yang ditentukan (dalam contoh ini, `:8080` berarti mendengarkan di semua interface jaringan yang tersedia pada port 8080).
- Aplikasi akan terus berjalan dan menunggu permintaan hingga dihentikan secara manual (misalnya, dengan Ctrl+C).

Untuk memahami kode dasar Go Gin bisa menggunakan kode dasar ini :

3.2-BasicGin.go

```go
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    // Inisialisasi Gin dengan middleware default (Logger dan Recovery)
    r := gin.Default()

    // Definisikan route untuk path root ("/") dengan metode GET
    r.GET("/", func(c *gin.Context) {
        // Kirim respons JSON dengan status 200 OK
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello From Gin!",
        })
    })

    // Jalankan server di port 8080
    r.Run(":8080")
}
```

Jalankan dengaan menggunakan perintah : 

```bash
$ go run 3.2-BasicGin.go
```
Lalu buka browser dengan URL :

```url
localhost:port/
```

atau karena port yang ada di code adalah `8080` :

```
localhost:8080/
```

Maka akan keluar dengan di browser seperti yang ada di gambar : 

![3.2.1-test](../../images/chapter3/3.2.1-1.BasicGin.png)

Dan pada log terminal akan menampilkan data transaksi :

```bash
$ go run 3.2-BasicGin.go 
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2025/07/01 - 15:46:55 | 200 |      36.007µs |       127.0.0.1 | GET      "/"
[GIN] 2025/07/01 - 15:46:55 | 404 |       1.122µs |       127.0.0.1 | GET      "/favicon.ico"
```

Test telah berhasil, terlihat terdapat transaksi, dan di repository sudah di sediakan source code tinggal buka [3.2-BasicGin.go]()

## 3.3 Routing

### 3.3.1 Basic routing (GET, POST, PUT, DELETE)

Setiap permintaan HTTP memiliki metode atau verb yang menunjukkan jenis operasi yang ingin dilakukan klien. Gin menyediakan fungsi yang sesuai untuk setiap metode ini, memungkinkan kita untuk mendefinisikan handler atau fungsi yang akan dieksekusi ketika permintaan dengan metode dan path tertentu diterima. Berikut adalah contoh penggunaan metode routing dasar di Gin :

#### 3.2.1.1 GET Method

Metode GET digunakan untuk meminta data dari sumber daya tertentu. Ini adalah metode yang paling umum digunakan dan biasanya digunakan untuk mengambil halaman web, gambar, atau data API.

3.2.1.1-1.TryGet.go

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/welcome", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Welcome to Server Gin!",
        })
    })
    router.Run(":8080")
}
```

Kode di atas akan bermakna, jika server mendapatkan `/welcome` pada browser maka server akan mengirmkan JSON message dengan status OK dan isi pada key `message` akan berisi `Welcome to Server Gin!`. Coba untuk run :

```bash
$ go run 3.2.1.1-1.TryGet.go
```

Ketika di akses dengan menggunakan URL `localhost:8080/welcome` akan mengeluarkan  :



Kita akan coba dengan menambah GET dengan routing `/checkserver' :

3.2.1.1-2.TryGet.go

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/welcome", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Welcome to Server Gin!",
        })
    })

    router.GET("/checkserver", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Server is running well!",
        })
    })
    
    router.Run(":8080")
}
```

Jika akses di browser antara `localhost:8080/welcome` dan `localhost:8080/checkserver` maka pasti akan berbeda output yang di keluarkan :

Contoh kasus nyata lagi meminta data berdasarkan pada sebuah ID, kita panggil dengan penggunaan URL `localhost:8080/users/1` dengan kode :

3.2.1.1-3.TryGet.go

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    // Mendefinisikan route dengan parameter dinamis :id
    router.GET("/users/:id", func(c *gin.Context) {
        // Mengambil nilai dari parameter 'id' dari URL
        id := c.Param("id")
        c.JSON(http.StatusOK, gin.H{
            "message": "Mengambil data user dengan ID: " + id,
        })
    })
    router.Run(":8080")
}
```

Maka ketika di panggil dengan `localhost:8080/user/1` maka akan menjawab :



Dan disitulah, dengan uji coba code di atas dapat menjelasakan dari GET Method. Untuk source cdoe bisa di lihat di []() dan []()

#### 3.2.1.2 POST Method

Metode POST digunakan untuk mengirim data dari pengguna ke server untuk membuat sumber daya baru. Data yang dikirim biasanya berada di dalam body dari permintaan (request body). 

> Keterangan : Pengujian POST memerlukan tools khusus seperti `cURL`, Postman, atau Insomnia untuk mengirim permintaan.

3.2.1.2-1.TryPost.go

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    // Mendefinisikan route untuk metode POST di path "/create-user"
    router.POST("/create-user", func(c *gin.Context) {
        // ... Logika untuk membuat user baru akan ada di sini ...
        // Untuk saat ini, kita hanya mengirim respons konfirmasi
        c.JSON(http.StatusCreated, gin.H{
            "message": "User berhasil dibuat.",
        })
    })
    router.Run(":8080")
}
```

Kode di atas akan membuat sebuah endpoint `/create-user` yang hanya menerima metode POST. Jika endpoint ini diakses dengan metode POST, server akan merespons dengan status `201 Created` dan sebuah pesan JSON.

Jalankan server :

```bash
$ go run 3.2.1.2-1.TryPost.go
```

Untuk mengujinya, gunakan `cURL` di terminal :

```bash
$ curl -X POST http://localhost:8080/create-user
```

Anda akan mendapatkan output JSON berikut :

```json
{"message":"User berhasil dibuat."}
```

Sama seperti GET, kita bisa mendefinisikan beberapa route POST dalam satu aplikasi.

3.2.1.2-2.TryPost.go

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    
    router.POST("/create-user", func(c *gin.Context) {
        c.JSON(http.StatusCreated, gin.H{
            "message": "User berhasil dibuat.",
        })
    })

    router.POST("/create-product", func(c *gin.Context) {
        c.JSON(http.StatusCreated, gin.H{
            "message": "Produk berhasil dibuat.",
        })
    })
    
    router.Run(":8080")
}
```

Sekarang, jika Anda menjalankan kode di atas dan mengirim permintaan POST ke endpoint yang berbeda, Anda akan mendapatkan respons yang berbeda pula.

Uji endpoint `/create-product` :
```bash
$ curl -X POST http://localhost:8080/create-product
```

Output :

```json
{"message":"Produk berhasil dibuat."}
```

Ini menunjukkan bagaimana Gin dapat dengan mudah memetakan permintaan POST ke handler yang berbeda berdasarkan path URL. Untuk source code bisa dilihat di []() dan []().

#### 3.2.1.3 PUT Method

Metode PUT digunakan untuk memperbarui sumber daya yang sudah ada di server. Biasanya, permintaan PUT menyertakan ID dari sumber daya yang akan diubah di URL dan data baru di dalam request body.

> Keterangan : Sama seperti POST, pengujian PUT memerlukan tools khusus seperti `cURL`, Postman, atau Insomnia.

3.2.1.3-1.TryPut.go
```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    // Mendefinisikan route untuk metode PUT di path "/users/:id"
    // :id adalah parameter dinamis yang bisa diambil dari URL
    router.PUT("/users/:id", func(c *gin.Context) {
        id := c.Param("id") // Mengambil nilai parameter "id" dari URL
        // Logika untuk memperbarui user dengan ID tertentu akan ada di sini
        c.JSON(http.StatusOK, gin.H{
            "message": "User dengan ID " + id + " berhasil diperbarui.",
        })
    })
    router.Run(":8080")
}
```
Kode di atas mendefinisikan endpoint `/users/:id` yang merespons metode PUT. Bagian `:id` adalah *route parameter* yang memungkinkan URL menjadi dinamis. Nilai dari `id` bisa diambil menggunakan `c.Param("id")`.

Jalankan server:
```bash
$ go run 3.2.1.3-1.TryPut.go
```

Untuk mengujinya, gunakan `cURL` dan berikan ID user yang ingin di-update, misalnya `123`:
```bash
$ curl -X PUT http://localhost:8080/users/123
```

Anda akan mendapatkan output JSON berikut :
```json
{"message":"User dengan ID 123 berhasil diperbarui."}
```

Kita juga bisa memiliki beberapa route PUT yang berbeda.

3.2.1.3-2.TryPut.go
```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    
    router.PUT("/users/:id", func(c *gin.Context) {
        id := c.Param("id")
        c.JSON(http.StatusOK, gin.H{
            "message": "User dengan ID " + id + " berhasil diperbarui.",
        })
    })

    router.PUT("/products/:productId", func(c *gin.Context) {
        productId := c.Param("productId")
        c.JSON(http.StatusOK, gin.H{
            "message": "Produk dengan ID " + productId + " berhasil diperbarui.",
        })
    })
    
    router.Run(":8080")
}
```
Uji endpoint `/products/:productId`:
```bash
$ curl -X PUT http://localhost:8080/products/abc-456
```

Output:
```json
{"message":"Produk dengan ID abc-456 berhasil diperbarui."}
```

Ini menunjukkan fleksibilitas Gin dalam menangani pembaruan data untuk berbagai jenis sumber daya. Untuk source code bisa dilihat di []() dan []().

#### 3.2.1.4 DELETE Method

Metode DELETE digunakan untuk menghapus sumber daya tertentu dari server. Sama seperti PUT, permintaan DELETE biasanya menyertakan ID dari sumber daya yang akan dihapus di URL. Endpoint DELETE sangat umum digunakan pada API untuk menghapus data berdasarkan parameter unik, seperti ID.

Contoh implementasi endpoint DELETE pada Gin :

3.2.1.4-1.TryDelete.go

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    // Mendefinisikan route untuk metode DELETE di path "/users/:id"
    router.DELETE("/users/:id", func(c *gin.Context) {
        id := c.Param("id") // Mengambil ID dari URL
        // Logika untuk menghapus user dengan ID tertentu
        c.JSON(http.StatusOK, gin.H{
            "message": "User dengan ID " + id + " berhasil dihapus.",
        })
    })
    router.Run(":8080")
}
```

Jalankan server dengan perintah berikutb :

```bash
$ go run 3.2.1.4-1.TryDelete.go
```

Untuk menguji endpoint DELETE, gunakan `cURL` di terminal:

```bash
$ curl -X DELETE http://localhost:8080/users/42
```

Output yang dihasilkan:

```json
{"message":"User dengan ID 42 berhasil dihapus."}
```

Anda juga dapat menambahkan beberapa endpoint DELETE untuk sumber daya lain, misalnya produk:

3.2.1.4-2.TryDelete.go

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    
    router.DELETE("/users/:id", func(c *gin.Context) {
        id := c.Param("id")
        c.JSON(http.StatusOK, gin.H{
            "message": "User dengan ID " + id + " berhasil dihapus.",
        })
    })

    router.DELETE("/products/:productId", func(c *gin.Context) {
        productId := c.Param("productId")
        c.JSON(http.StatusOK, gin.H{
            "message": "Produk dengan ID " + productId + " berhasil dihapus.",
        })
    })
    
    router.Run(":8080")
}
```

Uji endpoint `/products/:productId`:

```bash
$ curl -X DELETE http://localhost:8080/products/abc-123
```

Output:

```json
{"message":"Produk dengan ID abc-123 berhasil dihapus."}
```

Dengan demikian, Gin memudahkan pembuatan endpoint DELETE untuk berbagai jenis sumber daya. Untuk source code lengkap dapat dilihat di []() dan []().

### 3.2.2 Route Parameters dan Query parameters

Seringkali, kita perlu menangani permintaan yang bervariasi berdasarkan data spesifik dalam URL. Gin menyediakan dua cara utama untuk menangani ini yaitu Route Parameters dan Query Parameters.

#### 3.2.2.1 Route Parameters

Route parameters adalah bagian dari URL yang memungkinkan kita menangkap nilai dinamis pada path tertentu. Biasanya, route parameters didefinisikan dengan awalan titik dua (`:`) di dalam pola rute. Contohnya, pada endpoint `/users/:id`, bagian `:id` akan menangkap nilai apa pun yang diberikan pada posisi tersebut di URL, misalnya `/users/5` atau `/users/abc123`.

Route parameters sangat berguna ketika kita ingin mengakses data spesifik berdasarkan identitas unik, seperti ID user, kode produk, atau slug artikel. Gunakan route parameters jika nilai tersebut merupakan bagian utama dari identitas sumber daya yang diakses.

Berikut contoh implementasi route parameters di Gin:

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Mendefinisikan route dengan parameter dinamis ":id"
    r.GET("/users/:id", func(c *gin.Context) {
        id := c.Param("id") // Mengambil nilai parameter "id" dari URL
        c.JSON(http.StatusOK, gin.H{
            "message": "Detail user dengan ID: " + id,
        })
    })

    r.Run(":8080")
}
```

Ketika coba di akses

```bash

```

Pada contoh di atas, jika ada permintaan ke `/users/42`, maka nilai `42` akan diambil melalui `c.Param("id")` dan dapat digunakan di dalam handler. Dengan demikian, route parameters memudahkan pembuatan endpoint yang fleksibel dan dinamis sesuai kebutuhan aplikasi.

Untuk mencoba Routes Parameter, bisa menggunakan kode di atas, atau bisa di akses di 

#### 3.2.2.2 Query Parameters

Query parameters adalah pasangan kunci-nilai yang ditambahkan di akhir URL setelah tanda tanya (?). Mereka dipisahkan oleh ampersand (&). Contoh URL dengan query parameters adalah `/products?category=elektronik&sort=price_asc`. Di sini, category dan sort adalah query parameters.

Untuk kapan menggunakan query parameters Gunakan query parameters untuk memfilter, mengurutkan, atau menyediakan data opsional yang tidak secara langsung mengidentifikasi sumber daya. Misalnya, untuk paginasi (page=1&limit=10), pencarian (q=laptop), atau filtering (status=active). Berikut contoh kodenya :

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Mendefinisikan rute yang akan menggunakan query parameters
	r.GET("/products", func(c *gin.Context) {
		category := c.Query("category") // Mengambil nilai query parameter "category"
		sort := c.Query("sort")         // Mengambil nilai query parameter "sort"

		if category != "" && sort != "" {
			c.JSON(http.StatusOK, gin.H{"message": "Mendapatkan produk kategori: " + category + " dengan sort: " + sort})
		} else if category != "" {
			c.JSON(http.StatusOK, gin.H{"message": "Mendapatkan produk kategori: " + category})
		} else if sort != "" {
			c.JSON(http.StatusOK, gin.H{"message": "Mendapatkan produk dengan sort: " + sort})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Mendapatkan semua produk"})
		}
	})

	r.Run(":8080")
}
```

Di sini, c.Query("category") dan c.Query("sort") digunakan untuk mengambil nilai dari query parameters. Kita juga bisa menggunakan c.DefaultQuery("paramName", "defaultValue") untuk memberikan nilai default jika parameter tidak ada.

### 3.2.3 Route groups dan Middlewares

### 3.2.4 Static file serving

## 3.2 Request Handling

### 3.2.1 Binding request data (JSON, Form, Query)

### 3.2.2 Request validation

### 3.2.3 Response formatting (JSON, XML, HTML)

### 3.2.4 Error handling patterns