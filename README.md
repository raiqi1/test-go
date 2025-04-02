# Product Management API

A RESTful API for product management built with Go, PostgreSQL, and Redis.

## Persyaratan Sistem

- Go (versi 1.x atau lebih tinggi)
- PostgreSQL
- Postgres driver untuk Go (`github.com/lib/pq`)
- Redis (untuk caching)
- Go Redis client (`github.com/go-redis/redis/v8`)
- Gin web framework (`github.com/gin-gonic/gin`)
- Godotenv untuk environment variables (`github.com/joho/godotenv`)

## Instalasi

### 1. Instalasi Go
Pastikan Go telah terinstal di sistem Anda. Anda dapat mengunduhnya di [https://golang.org/dl/](https://golang.org/dl/).

### 2. Menginstal PostgreSQL
Pastikan PostgreSQL telah terinstal dan berjalan di sistem Anda. Jika PostgreSQL belum terinstal, Anda bisa mengunduhnya dari [situs resmi PostgreSQL](https://www.postgresql.org/download/).

### 3. Menginstal Redis
Pastikan Redis telah terinstal dan berjalan di sistem Anda. Jika Redis belum terinstal, Anda bisa mengunduhnya dari [situs resmi Redis](https://redis.io/download).

### 4. Instalasi Dependensi
Jalankan perintah berikut untuk menginisialisasi proyek Go dan menginstal semua dependensi yang diperlukan:

```bash
# Inisialisasi proyek Go
go mod init product-management

# Instal dependensi
go get github.com/lib/pq
go get github.com/gin-gonic/gin
go get github.com/joho/godotenv
go get github.com/go-redis/redis/v8
go get gorm.io/gorm
go get gorm.io/driver/postgres

# Selesaikan dependensi
go mod tidy
```

### 5. Tambahkan file .env
Buat file `.env` di direktori root proyek dan tambahkan konfigurasi berikut:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=product_management

# Redis Configuration
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# Application Port
PORT=8080
```

## Menjalankan Aplikasi

1. Pastikan PostgreSQL dan Redis server berjalan
2. Jalankan aplikasi dengan perintah:

```bash
go run main.go
```

## API Endpoints

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET    | /products | Mendapatkan semua produk |
| GET    | /products/:id | Mendapatkan produk berdasarkan ID |
| POST   | /products | Membuat produk baru |
| PUT    | /products/:id | Memperbarui produk |
| DELETE | /products/:id | Menghapus produk |

## Fitur

- CRUD operasi untuk manajemen produk
- Caching dengan Redis untuk meningkatkan performa
- Respons API yang menunjukkan sumber data (database atau cache)