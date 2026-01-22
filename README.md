# Kasir API

Aplikasi API sederhana untuk sistem kasir dengan fitur manajemen categories dan produk.

## Tech Stack
- **Language**: Go 1.21+
- **API Documentation**: Swagger UI (auto-generated dengan swaggo/swag)
- **Data Storage**: In-memory (untuk development)

## Project Structure
```
.
├── main.go             # Handler dan route definitions
├── go.mod              # Go module file
├── go.sum              # Dependencies lock file
├── docs/               # Auto-generated API documentation
│   ├── docs.go         # Generated swagger Go code
│   ├── swagger.json    # OpenAPI spec (JSON format)
│   └── swagger.yaml    # OpenAPI spec (YAML format)
└── README.md           # File ini
```

## Quick Start

### 1. Run Server
```bash
go run main.go
```
Server akan berjalan di `http://localhost:8080`

### 2. Access Swagger UI
Buka browser ke:
```
http://localhost:8080/docs/
```

### 3. Health Check
```bash
curl http://localhost:8080/health | jq
```

## API Endpoints

### Categories
- `GET /categories` - Ambil semua kategori
- `POST /categories` - Buat kategori baru
- `GET /categories/{id}` - Ambil detail kategori
- `PUT /categories/{id}` - Update kategori
- `DELETE /categories/{id}` - Hapus kategori

### Produk
- `GET /api/produk` - Ambil semua produk
- `POST /api/produk` - Buat produk baru
- `GET /api/produk/{id}` - Ambil detail produk
- `PUT /api/produk/{id}` - Update produk
- `DELETE /api/produk/{id}` - Hapus produk

## Data Models

### Category
```json
{
  "id": 1,
  "name": "Makanan",
  "description": "Produk makanan dan minuman"
}
```

### Produk
```json
{
  "id": 1,
  "nama": "Indomie Goreng",
  "harga": 3500,
  "stok": 10
}
```

## Development

### Generate Swagger Documentation
Setelah mengubah handler atau comments, regenerate docs:
```bash
$(go env GOPATH)/bin/swag init
```

### Format Code
```bash
go fmt ./...
```

## Notes
- Data disimpan in-memory, akan hilang saat server restart
- Untuk production, gunakan persistent database (PostgreSQL, MySQL, dll)
- CORS belum dikonfigurasi - perlu tambahan untuk frontend

## TODO
- [ ] Add database integration (PostgreSQL)
- [ ] Add authentication/authorization
- [ ] Add validation middleware
- [ ] Add error handling middleware
- [ ] Add CORS support
- [ ] Add request logging

File tambahan
- `openapi.yaml` berisi spesifikasi OpenAPI minimal untuk endpoints kategori.
