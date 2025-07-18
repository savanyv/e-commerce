# 🛒 E-Commerce Backend API (Golang + Echo + GORM)

Backend API sederhana untuk manajemen produk dan brand menggunakan **Golang**, **Echo Framework**, dan **GORM**, dengan pendekatan arsitektur **Domain-Driven Design (DDD)**.

---

## 🚀 Fitur API

- **Manajemen Brand**
  - 🔹 Tambah brand → `POST /brands`
  - 🔹 Ambil semua brand + produk → `GET /brands`
  - 🔹 Hapus brand (jika tidak digunakan produk) → `DELETE /brands/:id`

- **Manajemen Produk**
  - 🔹 Tambah produk baru → `POST /products`
  - 🔹 Ambil semua produk dengan pagination → `GET /products?page=1&limit=10`
  - 🔹 Ambil detail produk → `GET /products/:id`
  - 🔹 Mengupdate produk → `PUT /products/:id`
  - 🔹 Hapus produk → `DELETE /products/:id`

---

## 📦 Teknologi yang Digunakan

- **Golang**
- **Echo** (HTTP framework)
- **GORM** (ORM)
- **PostgreSQL** (Database)
- **Validator** dari `go-playground/validator.v10`

---

## 🛠 Instalasi & Menjalankan Server

1. **Clone repository**
```bash
git clone https://github.com/savanyv/e-commerce.git
cd e-commerce
```

2. **Setup .env file**
```bash
# Postgres Database
HOST_DB=localhost
PORT_DB=5432
USER_DB=your-username
PASS_DB=your-password
NAME_DB=e-commerce

# Port Server
PORT_SERVER=8000
```

3. **Install Dependencies**
```bash
go mod tidy
```

4. **Jalankan Aplikasi**
```bash
go run cmd/main.go
```

Server berjalan di `localhost:8000/api`

---

## 📂 Struktur Direktori

```
e-commerce/
├── cmd/                     # Entry point aplikasi
│   └── main.go
├── config/                  # Load konfigurasi
│   └── config.go
├── internal/
│   ├── app/                 # Inisialisasi Echo + route
│   ├── database/            # Inisialisasi PostgreSQL
│   ├── delivery/
│   │   └── handlers/        # HTTP handler untuk brand & product
│   ├── routes/              # Routing modular
│       ├── brand_routes.go
│       ├── product_routes.go
│       └── routes.go
├── dto/                     # DTO untuk request/response
│   ├── brand_dto.go
│   └── product_dto.go
├── helpers/                 # Validasi custom
│   └── validator_helper.go
├── models/                  # Representasi tabel DB
│   ├── brand.go
│   └── product.go
├── repository/              # Layer akses DB
│   ├── brand_repo.go
│   └── product_repo.go
├── usecase/                 # Business logic
│   ├── brand_usecase.go
│   └── product_usecase.go
├── .env
├── .env.sample
├── go.mod
├── go.sum
└── readme.md
```

---

## 📖 Dokumentasi API

### 🔹 `POST /brands`
```json
// Request
{
  "name": "Logitech"
}

// Response
{
  "message": "brand created successfully"
}
```

### 🔹 `GET /brands`
```json
// Response
{
  "data": [
    {
      "id": 2,
      "name": "Nike",
      "products": [
        {
          "id": 2,
          "name": "Sepatu Sport",
          "price": 1000000,
          "quantity": 2
        },
        {
          "id": 1,
          "name": "Sepatu Nike",
          "price": 2000000,
          "quantity": 4
        }
      ]
    },
    {
      "id": 3,
      "name": "Logitech",
      "products": [
        {
          "id": 3,
          "name": "Mouse Logitech Superlight 2",
          "price": 1000000,
          "quantity": 2
        }
      ]
    },
    {
      "id": 4,
      "name": "Logitech",
      "products": null
    },
    {
      "id": 5,
      "name": "Asus",
      "products": null
    }
  ],
  "message": "brands retrieved successfully"
}
```

### 🔹 `DELETE /brands/:id`
```json
// Response
{
  "message": "brand deleted successfully"
}
```

---

### 🔹 `POST /products`
```json
// Request
{
  "name": "Mouse Logitech Superlight 2",
  "price": 1000000,
  "quantity": 2,
  "id_brand": 1
}

// Response
{
  "message": "product created successfully"
}
```

### 🔹 `GET /products?page=1&limit=2`
```json
// Response
{
  "data": [
    {
      "id": 2,
      "name": "Sepatu Sport",
      "price": 1000000,
      "quantity": 2,
      "brand": {
        "id": 2,
        "name": "Nike"
      }
    },
    {
      "id": 3,
      "name": "Mouse Logitech Superlight 2",
      "price": 1000000,
      "quantity": 2,
      "brand": {
        "id": 3,
        "name": "Logitech"
      }
    }
  ],
  "message": "products retrieved successfully",
  "total": 3
}
```

### 🔹 `GET /products/:id`
```json
// Response
{
  "data": {
    "id": 1,
    "name": "Sepatu Nike",
    "price": 2000000,
    "quantity": 4,
    "brand": {
      "id": 2,
      "name": "Nike"
    }
  },
  "message": "product retrieved successfully"
}
```

### 🔹 `PUT /products/:id`
```json
// Request
{
  "name": "Sepatu Nike",
  "price": 2000000,
  "quantity": 4,
  "id_brand": 2
}

// Response
{
  "message": "product updated successfully"
}
```

### 🔹 `DELETE /products/:id`
```json
// Response
{
  "message": "product deleted successfully"
}
```
