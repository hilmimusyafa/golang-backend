# Bab 1 : Penerapan Go Dasar ke Backend Development

## 1.1 Konsep Backend Development

Sebelum masuk dalam pemrograman aplikasi Backend dalam Go, harus memahami apa itu konsep Sistem Backend.

### 1.1.1 Pengertian backend dan perannya dalam aplikasi

Backend secara sederhana mengacu pada sisi server dari sebuah aplikasi. Backend merupakan otak dalam aplikasi. Tidak bisa di lihat langsung oleh pengguna aplikasi. Backend memiliki beberapa tanggung jawab pada aplikasi :

- **Menyimpan dan mengelola data** : Semua informasi yang dimasukkan atau lihat di aplikasi, seperti postingan, foto, profil pengguna, atau riwayat transaksi, disimpan dan diatur oleh backend di dalam database.
- **Logika bisnis** : Ini adalah aturan dan proses di balik layar yang membuat aplikasi berfungsi. Misalnya, ketika  login, backend akan memverifikasi kredensial. Ketika membeli sesuatu, backend akan memproses pembayaran dan memperbarui inventaris.
- **Keamanan** : Backend memastikan bahwa data aman dan terlindungi dari akses yang tidak sah.
- **Integrasi dengan sistem lain** : Backend bisa berkomunikasi dengan layanan eksternal, seperti layanan pembayaran, API pihak ketiga, atau sistem lainnya.

Peran backend sangat krusial karena tanpanya, aplikasi hanya akan menjadi antarmuka statis tanpa fungsionalitas. Berikut adalah beberapa peran utama backend :

1. **Menyediakan data untuk frontend** : Disaaat membuka aplikasi, sebenarnya frontend akan meminta data dari backend. Misal pada Instagram, backend akan mengirinkan data postingan terbaru ke aplikasi agar di tampilkan oleh frontend.
2. **Memproses permintaan dari user** : Setiap kali melakukan tindakan di aplikasi (misalnya, menekan tombol "suka", mengirim pesan, atau melakukan pencarian), permintaan tersebut akan dikirim ke backend untuk diproses.
3. **Melaksanakan Logika Bisnis** : Backend adalah tempat semua aturan dan logika yang kompleks diimplementasikan. Ini memastikan bahwa setiap tindakan pengguna sesuai dengan aturan aplikasi.
4. **Menjaga Konsistensi Data** : Backend memastikan bahwa data selalu akurat dan up-to-date di seluruh aplikasi. Jika ada perubahan data, backend akan mengelolanya sehingga semua pengguna melihat informasi yang sama.
5. **Menangani Autentikasi dan Otorisasi** : Backend bertanggung jawab untuk memverifikasi identitas pengguna (autentikasi) dan menentukan apa yang diizinkan untuk dilakukan oleh pengguna tersebut (otorisasi).
6. **Skalabilitas** : Backend dirancang untuk menangani sejumlah besar pengguna dan permintaan secara bersamaan, memastikan aplikasi tetap responsif meskipun banyak orang menggunakannya

Bisa dianalogikan, Bayangkan sebuah restoran. Frontend adalah meja, menu, dan pelayan yang berinteraksi langsung dengan pelanggan. Pelanggan memesan makanan (mengirim permintaan). Backend adalah dapur tempat koki menyiapkan makanan (memproses permintaan), gudang tempat bahan makanan disimpan (database), dan manajer yang memastikan semua proses berjalan lancar. Tanpa dapur (backend), Pelanggan tidak akan mendapatkan makanan yang pelanggan pesan.

### 1.1.2 Request-Response cycle

Siklus Permintaan-Respons (Request-Response Cycle) adalah dasar bagaimana internet dan sebagian besar aplikasi modern bekerja. Ini adalah proses komunikasi dua arah yang terjadi antara klien (misalnya, browser web atau aplikasi di ponsel) dan server (tempat backend aplikasi berada).

Untuk mempermudah akan menggunakan skenario Restoran sebelumnya. Ada beberapa tahapan yang perlu di perhatikan :

1. Mengirim Permintaan (Request)

Pelanggan memanggil pelayan dan bilang, "Saya mau nasi goreng!" Ini seperti Permintaan (Request) dari browser atau aplikasi Pelanggan ke server, meminta data atau melakukan sesuatu.

2. Pemrosesan di Server

Pelayan membawa pesanan Pelanggan (Client) ke dapur (server). Di sana, koki (backend) meracik nasi goreng Pesanan, mungkin mengambil bahan dari lemari es (database) dan memasaknya sesuai resep (logika bisnis). Ini adalah tahap Pemrosesan di Server.

3. Mengirim Respons

Begitu nasi goreng siap, pelayan mengantarkannya kembali ke meja Pelanggan. Ini adalah Respons (Response) dari server yang mengirimkan data atau hasil kembali ke klien.

4. Menampilkan di Klien

Pelanggan menerima nasi goreng dan langsung menyantapnya. Ini seperti browser atau aplikasi yang menampilkan informasi atau hasil yang dikirim server, agar bisa dilihat dan digunakan.

### 1.1.3 HTTP Protocol dan REST API principles

Untuk memungkinkan klien dan server berkomunikasi dalam siklus permintaan-respons, mereka membutuhkan bahasa yang sama. Di sinilah Protokol HTTP dan prinsip-prinsip REST API berperan penting.

#### 1.1.3.1 Protokol HTTP (Hypertext Transfer Protocol)

HTTP adalah protokol dasar yang digunakan untuk komunikasi data di World Wide Web. Ini adalah protokol tanpa status (stateless), artinya setiap permintaan dari klien ke server dianggap independen, dan server tidak "mengingat" permintaan sebelumnya dari klien yang sama. HTTP dirancang untuk komunikasi antara klien dan server. Klien (misalnya, browser web) mengirimkan permintaan HTTP, dan server mengirimkan respons HTTP.

Permintaan HTTP terdiri dari beberapa bagian :

1. Metode HTTP : Menunjukkan jenis tindakan yang ingin dilakukan. Metode yang paling umum adalah:
    - GET: Meminta data dari sumber daya yang ditentukan. Contoh: mengambil daftar produk.
    - POST: Mengirimkan data untuk diproses ke sumber daya yang ditentukan (misalnya, membuat sumber daya baru). Contoh: mengirim data formulir pendaftaran.
    - PUT : Memperbarui data yang ada di sumber daya yang ditentukan. Contoh: mengubah detail profil pengguna.
    - DELETE : Menghapus sumber daya yang ditentukan. Contoh: menghapus postingan.
    - PATCH : Memperbarui sebagian data di sumber daya.
2. URL (Uniform Resource Locator): Alamat unik dari sumber daya yang diminta.
3. Header: Berisi metadata tentang permintaan, seperti jenis konten yang diharapkan, autentikasi, atau informasi tentang klien.
4. Body (Opsional): Berisi data yang dikirimkan ke server, terutama untuk metode POST, PUT, atau PATCH.

Dan Respons HTTP yang dikirim server juga terdiri dari beberapa bagian :

1. Status Code: Kode numerik yang menunjukkan hasil dari permintaan. Beberapa kode status umum :
    - 200 OK: Permintaan berhasil diproses.
    - 201 Created: Sumber daya berhasil dibuat.
    - 400 Bad Request: Permintaan tidak valid dari sisi klien.
    - 401 Unauthorized: Klien tidak terautentikasi.
    - 403 Forbidden: Klien terautentikasi tetapi tidak memiliki izin.
    - 404 Not Found: Sumber daya yang diminta tidak ditemukan.
    - 500 Internal Server Error: Terjadi kesalahan di sisi server.
2. Header : Berisi metadata tentang respons, seperti jenis konten yang dikirim, ukuran file, atau informasi caching.
3. Body : Berisi data yang dikirim kembali ke klien, bisa berupa HTML, JSON, XML, gambar, atau jenis data lainnya.

#### 1.1.3.2 Prinsip-Prinsip REST API (Representational State Transfer Application Programming Interface)

REST bukan protokol, melainkan gaya arsitektur untuk membangun API web. API yang mengikuti prinsip-prinsip REST disebut RESTful API. Tujuan utama REST adalah membuat sistem yang dapat diskalakan, efisien, dan mudah dipelihara.

Prinsip-prinsip utama REST meliputi :

1. Client-Server : Pemisahan yang jelas antara klien dan server. Klien bertanggung jawab untuk UI, server untuk logika bisnis dan penyimpanan data. Ini memungkinkan kedua sisi untuk berkembang secara independen.
2. Stateless : Seperti HTTP, setiap permintaan dari klien ke server harus berisi semua informasi yang dibutuhkan server untuk memenuhi permintaan tersebut. Server tidak menyimpan "state" (status sesi) dari klien antar permintaan. Ini meningkatkan skalabilitas dan keandalan.
3. Cacheable : Respons dari server harus dapat dicache oleh klien untuk meningkatkan kinerja. Server harus menunjukkan apakah respons dapat dicache atau tidak.
4. Layered System : Klien tidak perlu tahu apakah mereka terhubung langsung ke server akhir atau ke middleware (seperti load balancer atau proxy). Ini meningkatkan fleksibilitas dan keamanan sistem.
5. Uniform Interface : Ini adalah prinsip paling krusial dan memiliki beberapa batasan : 
    - Resource Identification in Requests: Setiap "sumber daya" (data, misalnya pengguna, produk, postingan) harus memiliki pengenal unik, biasanya URL. Klien berinteraksi dengan sumber daya ini melalui URL.
    - Resource Manipulation Through Representations : Klien memanipulasi sumber daya dengan mengirimkan representasi (misalnya, data JSON atau XML) dari sumber daya tersebut ke server.
    - Self-descriptive Messages : Setiap pesan (permintaan atau respons) harus berisi informasi yang cukup untuk menafsirkan dirinya sendiri. Ini termasuk metode HTTP, URL, header, dan body.
    - Hypermedia as the Engine of Application State (HATEOAS) : Ini berarti respons API harus menyertakan tautan (hypermedia) yang memandu klien tentang tindakan selanjutnya yang dapat dilakukan. Meskipun ini adalah prinsip inti REST, seringkali tidak sepenuhnya diimplementasikan di banyak API "RESTful" praktis.

#### 1.1.3.3 Contoh Penerapan Protokol HTTP dan Konsep REST

Misalkan ada aplikasi e-commerce:
- Untuk mendapatkan daftar semua produk : Klien mengirimkan GET request ke /api/products. Server merespons dengan 200 OK dan body yang berisi array JSON produk.
- Untuk menambahkan produk baru : Klien mengirimkan POST request ke /api/products dengan body JSON yang berisi detail produk baru. Server merespons dengan 201 Created dan data produk yang baru dibuat.
- Untuk menghapus produk tertentu: Klien mengirimkan DELETE request ke /api/products/{id_produk}. Server merespons dengan 204 No Content (jika berhasil dihapus dan tidak ada konten untuk dikembalikan).

### 1.1.4 Stateless vs Stateful Applications

Dalam pengembangan backend, penting untuk memahami perbedaan antara aplikasi stateless dan stateful, terutama karena sifat HTTP yang stateless. Perbedaan ini memengaruhi bagaimana aplikasi dirancang, diskalakan, dan dikelola.

#### 1.1.4.1 Stateless Application

Aplikasi stateless adalah aplikasi di mana server tidak menyimpan informasi atau konteks apapun tentang sesi klien di antara permintaan-permintaan terpisah. Setiap permintaan dari klien ke server adalah independen dan harus mengandung semua informasi yang diperlukan server untuk memprosesnya.

Karakteristik :

- Setiap permintaan mandiri : Server memproses setiap permintaan tanpa perlu mengingat apa pun dari permintaan sebelumnya dari klien yang sama.
Tidak ada data sesi di server : Server tidak menyimpan data sesi pengguna di memori atau disknya. Jika ada data yang perlu dipertahankan antar permintaan (misalnya, status login), data tersebut harus disimpan di sisi klien (misalnya, dalam cookie, token, atau local storage) dan dikirimkan kembali ke server dengan setiap permintaan.
- Skalabilitas Horizontal Mudah : Karena setiap server tidak menyimpan status sesi,sehingga dapat dengan mudah menambahkan lebih banyak instance server untuk menangani peningkatan traffic (skalabilitas horizontal). Permintaan dari klien yang sama dapat diarahkan ke server mana pun yang tersedia tanpa masalah, karena setiap server dapat memprosesnya secara independen.
- Ketahanan (Resilience) Tinggi : Jika satu server mengalami down, permintaan klien dapat dengan mudah dialihkan ke server lain tanpa kehilangan data sesi.

Contoh :

- RESTful APIs : Sebagian besar API RESTful didesain stateless karena ini adalah salah satu prinsip inti REST.
- Website Statis : Server web yang hanya menyajikan file HTML, CSS, dan JavaScript statis.

Analogi :

Bayangkan seseorang menelepon layanan call center dan setiap kali seseorang menelepon, seseprang berbicara dengan agen yang berbeda dan harus menjelaskan masalah dari awal, karena agen tersebut tidak tahu apa yang sudah dibicarakan dengan agen sebelumnya.

#### 1.1.4.2 Statefull Application

Aplikasi stateful adalah aplikasi di mana server menyimpan informasi atau konteks tentang sesi klien di antara permintaan-permintaan yang berbeda. Server "mengingat" status klien dari interaksi sebelumnya.

Karakteristik dari :

- Ketergantungan pada sesi : Server mempertahankan data sesi untuk setiap klien. Data ini bisa disimpan dalam memori server, database sesi, atau cache bersama.
- Skalabilitas Horizontal Sulit : Jika status sesi disimpan di memori satu server, menambahkan server baru akan rumit karena permintaan dari klien yang sama harus selalu diarahkan ke server yang sama yang menyimpan sesi mereka (menggunakan sticky sessions atau session affinity). Jika server tersebut down, sesi klien akan hilang.
- Kompleksitas Manajemen Sesi : Membutuhkan mekanisme untuk mengelola sesi, seperti waktu kedaluwarsa sesi, replikasi sesi antar server (untuk ketahanan), atau penyimpanan sesi terpusat (misalnya, di Redis).
- Performa Awal yang Lebih Cepat : Untuk interaksi lanjutan dalam sesi yang sama, server mungkin tidak perlu meminta semua informasi dari klien lagi karena sudah menyimpannya.

Contoh :

- Aplikasi Web Tradisional (Server-Side Sessions): Banyak aplikasi web lama yang menggunakan sesi berbasis server untuk melacak status login, item keranjang belanja, atau preferensi pengguna.
- Server Game Online : Server yang perlu melacak posisi pemain, skor, dan status game secara real-time.
- Server WebSocket : Protokol WebSocket mempertahankan koneksi stateful antara klien dan server.

Analogi :

Seseorang menelepon call center yang sama dan selalu berbicara dengan agen yang sama, atau agen yang berbeda tetapi mereka memiliki riwayat lengkap percakapan sebelumnya dan tidak perlu menjelaskan dari awal.

#### 1.1.4.3 Mengapa Penting Membedakan Keduanya?

Pemilihan antara pendekatan stateless dan stateful sangat memengaruhi desain arsitektur aplikasi :

- Skalabilitas : Aplikasi stateless jauh lebih mudah untuk diskalakan secara horizontal karena bisa menambahkan server tanpa khawatir tentang migrasi atau replikasi sesi.
- Ketahanan : Aplikasi stateless lebih tahan terhadap kegagalan server tunggal.
- Kompleksitas : Aplikasi stateful seringkali lebih kompleks untuk dibangun dan dikelola, terutama dalam lingkungan terdistribusi.
- Performa : Tergantung pada kasus penggunaan, stateless bisa lebih cepat untuk setiap permintaan karena tidak ada overhead manajemen sesi, atau stateful bisa lebih cepat untuk interaksi berkelanjutan dalam satu sesi.

Dalam pengembangan backend modern, terutama dengan arsitektur mikroservis dan cloud-native, pendekatan stateless seringkali lebih diutamakan karena keunggulannya dalam skalabilitas dan ketahanan, yang sejalan dengan prinsip-prinsip REST. Jika status perlu dipertahankan, biasanya dilakukan melalui database eksternal atau cache terdistribusi yang dapat diakses oleh semua instance server, daripada disimpan di memori instance server itu sendiri.

## 1.2 Pengenalan Web Framework

### 1.2.1 Mengapa butuh framework (vs net/http standard library)

Dalam pengembangan backend Go, standar pustaka net/http menyediakan fungsionalitas dasar yang sangat kuat untuk membangun server web. Dapat membangun seluruh aplikasi web hanya dengan net/http. Namun, untuk aplikasi yang lebih kompleks dan produktif, web framework menjadi sangat berguna.

Berikut alasan mengapa butuh framework dibanding hanya menggunakan net/http standar :

1. Mengurangi Boilerplate Code : Pustaka net/http membutuhkan penulisan kode berulang untuk tugas-tugas umum seperti routing, parsing request body, penanganan middleware, atau validasi input. Framework menyediakan fungsionalitas ini secara bawaan, memungkinkan developer fokus pada logika bisnis inti.
2. Struktur Proyek Terorganisir : Framework seringkali menganjurkan atau bahkan memaksakan struktur direktori dan pola desain tertentu. Ini membantu menjaga kode tetap terorganisir, terutama dalam proyek besar dengan banyak developer, sehingga lebih mudah untuk dipahami dan dipelihara.
3. Fitur Bawaan (Built-in Features) : Sebagian besar framework dilengkapi dengan fitur-fitur yang sering dibutuhkan, seperti:
    - Routing Lanjutan : Mendukung pola route yang kompleks, parameter URL, dan grouping route.
    - Middleware : Mekanisme untuk menambahkan fungsionalitas ke request-response cycle secara global atau pada route tertentu (misalnya, logging, otentikasi, CORS).
    - Validasi Data : Membantu memvalidasi input dari klien.
    - Serialisasi/Deserialisasi JSON : Mempermudah konversi antara objek Go dan JSON.
    - Penanganan Error: Struktur yang lebih baik untuk mengelola dan merespons kesalahan.
4. Keamanan : Framework seringkali menyediakan lapisan keamanan bawaan atau praktik terbaik untuk melindungi aplikasi dari kerentanan umum seperti serangan XSS (Cross-Site Scripting) atau CSRF (Cross-Site Request Forgery).
5. Performa dan Optimasi: Banyak framework dioptimalkan untuk performa tinggi, dengan fitur seperti connection pooling, concurrency handling, dan caching. Meskipun net/http sangat efisien, framework menambahkan lapisan optimasi di atasnya.
6. Komunitas dan Ekosistem: Framework yang populer memiliki komunitas yang besar, yang berarti banyak resource, tutorial, dan plugin pihak ketiga yang dapat digunakan. Ini mempercepat proses pengembangan dan pemecahan masalah.

Meskipun net/http adalah dasar yang kuat dan dapat digunakan untuk membangun API yang sangat efisien, framework menyediakan abstraksi dan alat yang diperlukan untuk membangun aplikasi skala besar dengan lebih cepat dan mudah. Pilihan antara keduanya tergantung pada kompleksitas proyek dan kebutuhan developer. Untuk proyek kecil atau yang sangat spesifik, net/http mungkin cukup. Namun, untuk proyek yang lebih besar dan membutuhkan fitur lengkap, framework adalah pilihan yang lebih efisien.

Contoh Perbandingan Kode Sederhana (Endpoint "Halo Dunia")

Mari buat sebuah endpoint API sederhana yang merespons dengan pesan "Halo Dunia!" saat diakses : 

1.2.1-1.nonframework.go
```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"message": "Halo Dunia dari net/http!"}`)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Println("Server net/http berjalan di :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Dan dibandingkan juga dengan menggunakan Framework (Gin) :

1.2.1-2.framework.go
```go
package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() 

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Halo Dunia dari Gin!",
		})
	})

	log.Println("Server Gin berjalan di :8080")
	router.Run(":8080")
}
```
> Untuk mencoba kode bisa langsung membuka folder [`/golang-backend/source-code/chapter1`](../../source-code/chapter1)

Perbedaan Kunci yang Terlihat dari Contoh :

- Kemudahan Penulisan Respons JSON: Gin menyediakan c.JSON() yang jauh lebih ringkas untuk mengirim respons JSON dibandingkan dengan menulis w.Header().Set("Content-Type", "application/json"), w.WriteHeader(), dan fmt.Fprint() secara manual.
- Abstraksi Konteks : Gin menyediakan objek *gin.Context yang menyatukan fungsionalitas http.ResponseWriter dan *http.Request serta menambahkan banyak metode helper yang mempercepat pengembangan.
- Middleware Bawaan : gin.Default() sudah menyertakan middleware logging dan recovery, yang harus diimplementasikan secara manual dengan net/http jika dibutuhkan.

Meskipun contoh "Halo Dunia" ini sederhana, perbedaan dalam cara menulis response JSON sudah menunjukkan bagaimana framework mengurangi kode boilerplate dan membuat pengembangan lebih efisien. Dalam aplikasi yang lebih kompleks dengan routing yang banyak, validasi input, dan middleware otentikasi, keuntungan framework akan jauh lebih signifikan.

### 1.2.2 Perbandingan Framework Go: Gin vs Echo vs Fiber vs Chi

Ekosistem framework web di Go cukup kaya, dengan beberapa pilihan populer yang masing-masing memiliki kekuatan dan kelemahan. Memilih framework yang tepat tergantung pada kebutuhan spesifik proyek, preferensi performa, dan fitur yang dibutuhkan. Berikut perbandingan beberapa framework Go yang sering digunakan :

| Fitur / Framework | Gin | Echo | Fiber | Chi |
|-------------------|-----|------|-------|-----|
| **Tipe** | Micro-framework | Micro-framework | Web framework | Minimalist router |
| **Performa** | Sangat cepat, performa tinggi | Sangat cepat, performa tinggi | Sangat cepat, seringkali tercepat karena berbasis Fasthttp | Cepat, efisien |
| **Filosofi** | Cepat & ringan, fokus pada API RESTful | High performance, minimalis, extensible | Express.js-like API, performa ekstrem | Minimalis, middleware-first, composition over configuration |
| **Dependencies** | Ringan | Ringan | Menggunakan Fasthttp | Sangat ringan |
| **Kelebihan** | - Performa sangat baik<br>- Banyak middleware bawaan dan pihak ketiga<br>- Komunitas besar dan dokumentasi bagus<br>- API friendly | - Performa sangat baik<br>- API intuitif dan extensible<br>- Banyak fitur bawaan (validasi, renderer, middleware)<br>- Komunitas aktif | - Performa tercepat (untuk use case tertentu)<br>- Sintaks mirip Express.js (familiar bagi developer Node.js)<br>- Banyak fitur seperti templating, websocket, rate limit | - Sangat ringan dan modular<br>- Fokus pada routing yang kuat dan terstruktur<br>- Cocok untuk microservices atau API kecil<br>- Fleksibel dengan middleware |
| **Kekurangan** | - Beberapa keputusan desain mungkin terasa kaku<br>- Agak lebih besar dari Gin dalam ukuran biner | - Bergantung pada Fasthttp yang memiliki API berbeda dari net/http standar (potensi masalah kompatibilitas library)<br>- Komunitas lebih kecil dibanding Gin/Echo | - Hanya menyediakan routing, perlu menambahkan banyak komponen secara manual<br>- Kurang cocok untuk proyek yang butuh banyak fitur bawaan framework |  |
| **Use Case Terbaik** | - Membangun API RESTful yang cepat<br>- Proyek web berskala menengah hingga besar | - API RESTful performa tinggi<br>- Aplikasi web skala penuh<br>- Proyek yang mencari keseimbangan fitur dan performa | - API performa sangat tinggi<br>- Aplikasi yang membutuhkan respons cepat dan throughput tinggi (mirip Express.js) | - Microservices kecil<br>- API sederhana dan bersih<br>- Proyek yang memprioritaskan kontrol penuh dan minim dependensi |

Dan di dalam dokumentasi ini, akan terkunci pada penggunaan Framework Gin.

### 1.2.3 Arsitektur MVC dalam konteks Go

Meskipun Go tidak secara native menganut pola arsitektur MVC (Model-View-Controller) sekuat bahasa atau framework lain seperti Ruby on Rails atau Laravel, konsepnya tetap dapat diterapkan dan sering menjadi referensi untuk mengorganisir kode backend di Go.

MVC adalah pola arsitektur perangkat lunak yang memisahkan aplikasi menjadi tiga komponen utama untuk memisahkan kekhawatiran (separation of concerns) dan meningkatkan modularitas serta pemeliharaan :

- Model : Bertanggung jawab untuk data dan logika bisnis. Ini berinteraksi langsung dengan database, melakukan validasi data, dan menerapkan aturan bisnis. Model adalah representasi dari data yang digunakan aplikasi dan bagaimana data tersebut berinteraksi.
- View : Bertanggung jawab untuk presentasi data kepada pengguna. Dalam aplikasi web, ini seringkali adalah file HTML, CSS, dan JavaScript yang dilihat oleh pengguna. Dalam konteks API backend murni, "View" ini bisa berarti format data yang dikirimkan sebagai respons (misalnya, JSON atau XML).
- Controller : Bertindak sebagai perantara antara Model dan View. Controller menerima input dari pengguna (melalui View), memanggil Model untuk memproses data atau logika, dan kemudian memilih View yang tepat untuk menampilkan hasil kepada pengguna.

Dalam Go backend, terutama untuk API RESTful yang tidak menyajikan HTML secara langsung (artinya, frontend terpisah), konsep "View" sedikit berbeda.

1. Model dalam Go Backend : 
   - Direktori atau paket (package) yang berisi definisi struktur data (structs) yang merepresentasikan entitas database atau objek bisnis.
   - Berisi logika untuk berinteraksi dengan database (misalnya, fungsi untuk membuat, membaca, memperbarui, menghapus data - CRUD).
   - Dapat juga berisi logika validasi data dan aturan bisnis terkait entitas tersebut.
   - Contoh: models/user.go, models/product.go
2. View dalam Go Backend (untuk API RESTful):
   - Karena backend Go seringkali menyediakan API untuk frontend terpisah, "View" di sini lebih mengacu pada struktur respons data (biasanya JSON).
   - Bisa berupa struct terpisah yang digunakan untuk memformat data sebelum dikirimkan sebagai respons (misalnya, menyembunyikan field sensitif atau menggabungkan data dari berbagai model). Ini sering disebut sebagai "DTO" (Data Transfer Object) atau "Resource".
   - Contoh: Fungsi dalam handler yang mem-marshall struct ke JSON, atau struct khusus responses/user_response.go.
3. Controller dalam Go Backend :
   - Direktori atau paket yang berisi handler fungsi untuk HTTP request.
   - Menerima HTTP request dari router.
   - Menganalisis request (misalnya, membaca parameter URL, query, atau request body).
   - Memanggil metode yang sesuai dari Model untuk melakukan operasi database atau logika bisnis.
   - Memproses hasil dari Model.
   - Mengirimkan Respons HTTP (dengan status code dan body yang diformat, biasanya JSON) kembali ke klien.
   - Contoh: controllers/user_controller.go, controllers/product_controller.go.
4. Hubungan dengan Router :
   - Router (misalnya, yang disediakan oleh Gin, Echo, Fiber, atau Chi) bertanggung jawab untuk mengarahkan incoming HTTP request ke fungsi Controller yang tepat berdasarkan URL dan metode HTTP.

Struktur Proyek Contoh dengan MVC :
```
my-go-app/
├── main.go               // Titik masuk aplikasi
├── config/               // Konfigurasi aplikasi (database, port, dll)
├── routes/               // Definisi semua rute dan menghubungkannya ke controller
│   └── api.go
├── controllers/          // Handler HTTP, berinteraksi dengan models
│   ├── user_controller.go
│   └── product_controller.go
├── models/               // Definisi struct data dan logika interaksi database
│   ├── user.go
│   └── product.go
├── database/             // Setup koneksi database
│   └── db.go
├── middlewares/          // Custom HTTP middleware (autentikasi, logging)
│   └── auth_middleware.go
└── utils/                // Fungsi utilitas umum (helpers)
    └── validator.go`
```

Kelebihan Penerapan MVC (atau Pola Serupa) di Go : 
- Pemisahan Kekhawatiran (Separation of Concerns) : Setiap bagian memiliki tanggung jawab yang jelas, membuat kode lebih modular dan mudah dipahami.
- Pemeliharaan yang Lebih Mudah : Perubahan pada logika bisnis (Model) tidak langsung memengaruhi cara data disajikan (View/Respons) atau bagaimana request ditangani (Controller), dan sebaliknya.
- Pengujian (Testing) : Lebih mudah untuk menguji setiap komponen secara independen.
- Kolaborasi Tim : Developer dapat bekerja pada bagian yang berbeda dari aplikasi tanpa terlalu banyak konflik.

Meskipun Go tidak memaksakan MVC, mengadopsi prinsip pemisahan tanggung jawab seperti yang ada di MVC sangat dianjurkan untuk membangun aplikasi backend yang bersih, terukur, dan mudah dikelola.