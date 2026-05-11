# API Spec — Book Management

**Base URL**: `http://localhost:8080`

---

## Model

### Book (Response)
```json
{
  "id": 1,
  "title": "string",
  "author": "string",
  "year": 2024,
  "stock": 5
}
```

### BookInput (Request Body)
```json
{
  "title": "string",
  "author": "string",
  "year": 2024,
  "stock": 0
}
```

| Field    | Tipe | Wajib | Keterangan              |
|----------|------|-------|-------------------------|
| `title`  | string | ✅   | Judul buku              |
| `author` | string | ✅   | Penulis buku            |
| `year`   | int   | ✅    | Tahun terbit            |
| `stock`  | int   | ❌    | Jumlah stok (default: 0) |

---

## Endpoints

### 1. Tambah Buku

**Request**
```
POST /books
Content-Type: application/json

{
  "title": "Belajar Go",
  "author": "John Doe",
  "year": 2024,
  "stock": 10
}
```

**Response — `201 Created`**
```json
{
  "id": 1,
  "title": "Belajar Go",
  "author": "John Doe",
  "year": 2024,
  "stock": 10
}
```

**Response — `400 Bad Request`** (validasi gagal)
```json
{
  "error": "Key: 'BookInput.Title' Error:Field validation for 'Title' failed on the 'required' tag"
}
```

---

### 2. Lihat Semua Buku

**Request**
```
GET /books
```

**Response — `200 OK`**
```json
[
  {
    "id": 1,
    "title": "Belajar Go",
    "author": "John Doe",
    "year": 2024,
    "stock": 10
  },
  {
    "id": 2,
    "title": "Pemrograman Web",
    "author": "Jane Doe",
    "year": 2023,
    "stock": 3
  }
]
```

---

### 3. Detail Buku

**Request**
```
GET /books/:id
```

**Response — `200 OK`**
```json
{
  "id": 1,
  "title": "Belajar Go",
  "author": "John Doe",
  "year": 2024,
  "stock": 10
}
```

**Response — `400 Bad Request`** (ID bukan angka)
```json
{
  "error": "ID harus angka"
}
```

**Response — `404 Not Found`**
```json
{
  "error": "buku dengan ID 99 tidak ditemukan"
}
```

---

### 4. Update Buku

**Request**
```
PUT /books/:id
Content-Type: application/json

{
  "title": "Belajar Go Edisi 2",
  "author": "John Doe",
  "year": 2025,
  "stock": 7
}
```

**Response — `200 OK`**
```json
{
  "id": 1,
  "title": "Belajar Go Edisi 2",
  "author": "John Doe",
  "year": 2025,
  "stock": 7
}
```

**Response — `400 Bad Request`** (validasi gagal)
```json
{
  "error": "Key: 'BookInput.Title' Error:Field validation for 'Title' failed on the 'required' tag"
}
```

**Response — `404 Not Found`**
```json
{
  "error": "buku dengan ID 99 tidak ditemukan"
}
```

---

### 5. Hapus Buku

**Request**
```
DELETE /books/:id
```

**Response — `200 OK`**
```json
{
  "message": "Buku berhasil dihapus"
}
```

**Response — `400 Bad Request`** (ID bukan angka)
```json
{
  "error": "ID harus angka"
}
```

**Response — `404 Not Found`**
```json
{
  "error": "buku dengan ID 99 tidak ditemukan"
}
```

---

## Ringkasan

| Method | Endpoint         | Deskripsi      | Status |
|--------|------------------|----------------|--------|
| POST   | `/books`         | Tambah buku    | ✅     |
| GET    | `/books`         | Semua buku     | ✅     |
| GET    | `/books/:id`     | Detail buku    | ✅     |
| PUT    | `/books/:id`     | Update buku    | ✅     |
| DELETE | `/books/:id`     | Hapus buku     | ✅     |
