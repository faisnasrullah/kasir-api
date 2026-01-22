// @title Kasir API
// @version 1.0
// @description API sederhana untuk kasir - categories dan produk
// @host localhost:8080
// @BasePath /

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "kasir-api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// Produk represents a product in the cashier system
type Produk struct {
	ID    int    `json:"id" example:"1"`
	Nama  string `json:"nama" example:"Indomie Goreng"`
	Harga int    `json:"harga" example:"3500"`
	Stok  int    `json:"stok" example:"10"`
}

// In-memory storage (sementara, nanti ganti database)
var produk = []Produk{
	{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
	{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40},
	{ID: 3, Nama: "kecap", Harga: 12000, Stok: 20},
}

// Category represents a product category
type Category struct {
	ID          int    `json:"id" example:"1"`
	Name        string `json:"name" example:"Makanan"`
	Description string `json:"description" example:"Produk makanan dan minuman"`
}

// In-memory categories
var categories = []Category{
	{ID: 1, Name: "Makanan", Description: "Produk makanan dan minuman"},
	{ID: 2, Name: "Minuman", Description: "Segala jenis minuman"},
}

// CategoryInput used for create/update payloads
type CategoryInput struct {
	Name        string `json:"name" example:"Alat Tulis"`
	Description string `json:"description" example:"Pulpen, pensil"`
}

// ProdukInput used for create/update produk
type ProdukInput struct {
	Nama  string `json:"nama" example:"Vit 1000ml"`
	Harga int    `json:"harga" example:"3000"`
	Stok  int    `json:"stok" example:"40"`
}

// @Summary Get all categories
// @Tags Categories
// @Produce application/json
// @Success 200 {array} Category
// @Router /categories [get]
func getCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

// @Summary Create a category
// @Tags Categories
// @Accept application/json
// @Produce application/json
// @Param category body CategoryInput true "Category input"
// @Success 201 {object} Category
// @Router /categories [post]
func createCategory(w http.ResponseWriter, r *http.Request) {
	var c Category
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	c.ID = len(categories) + 1
	categories = append(categories, c)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

// @Summary Get a category by ID
// @Tags Categories
// @Produce application/json
// @Param id path int true "Category ID"
// @Success 200 {object} Category
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [get]
func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}
	for _, c := range categories {
		if c.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(c)
			return
		}
	}
	http.Error(w, "Category not found", http.StatusNotFound)
}

// @Summary Update a category
// @Tags Categories
// @Accept application/json
// @Produce application/json
// @Param id path int true "Category ID"
// @Param category body CategoryInput true "Category input"
// @Success 200 {object} Category
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [put]
func updateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}
	var upd Category
	if err := json.NewDecoder(r.Body).Decode(&upd); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	for i := range categories {
		if categories[i].ID == id {
			upd.ID = id
			categories[i] = upd
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(upd)
			return
		}
	}
	http.Error(w, "Category not found", http.StatusNotFound)
}

// @Summary Delete a category
// @Tags Categories
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [delete]
func deleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}
	for i, c := range categories {
		if c.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "category deleted"})
			return
		}
	}
	http.Error(w, "Category not found", http.StatusNotFound)
}

// @Summary Get a produk by ID
// @Tags Produk
// @Produce application/json
// @Param id path int true "Produk ID"
// @Success 200 {object} Produk
// @Failure 404 {object} map[string]string
// @Router /api/produk/{id} [get]
func getProdukByID(w http.ResponseWriter, r *http.Request) {
	// Parse ID dari URL path
	// URL: /api/produk/123 -> ID = 123
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	// Cari produk dengan ID tersebut
	for _, p := range produk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	// Kalau tidak found
	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// @Summary Get all produk
// @Tags Produk
// @Produce application/json
// @Success 200 {array} Produk
// @Router /api/produk [get]
func getProduks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produk)
}

// @Summary Create a produk
// @Tags Produk
// @Accept application/json
// @Produce application/json
// @Param produk body ProdukInput true "Produk input"
// @Success 201 {object} Produk
// @Router /api/produk [post]
func createProduk(w http.ResponseWriter, r *http.Request) {
	var produkBaru Produk
	err := json.NewDecoder(r.Body).Decode(&produkBaru)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	produkBaru.ID = len(produk) + 1
	produk = append(produk, produkBaru)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(produkBaru)
}

// @Summary Update a produk
// @Tags Produk
// @Accept application/json
// @Produce application/json
// @Param id path int true "Produk ID"
// @Param produk body ProdukInput true "Produk input"
// @Success 200 {object} Produk
// @Failure 404 {object} map[string]string
// @Router /api/produk/{id} [put]
func updateProduk(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	// ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	// get data dari request
	var updateProduk Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// loop produk, cari id, ganti sesuai data dari request
	for i := range produk {
		if produk[i].ID == id {
			updateProduk.ID = id
			produk[i] = updateProduk

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduk)
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// @Summary Delete a produk
// @Tags Produk
// @Param id path int true "Produk ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/produk/{id} [delete]
func deleteProduk(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	// loop produk cari ID, dapet index yang mau dihapus
	for i, p := range produk {
		if p.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			produk = append(produk[:i], produk[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

func main() {
	// GET localhost:8080/api/produk/{id}
	// PUT localhost:8080/api/produk/{id}
	// DELETE localhost:8080/api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProdukByID(w, r)
		} else if r.Method == "PUT" {
			updateProduk(w, r)
		} else if r.Method == "DELETE" {
			deleteProduk(w, r)
		}
	})

	// Category routes
	// GET /categories
	// POST /categories
	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getCategories(w, r)
		} else if r.Method == "POST" {
			createCategory(w, r)
		}
	})

	// GET /categories/{id}
	// PUT /categories/{id}
	// DELETE /categories/{id}
	http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getCategoryByID(w, r)
		} else if r.Method == "PUT" {
			updateCategory(w, r)
		} else if r.Method == "DELETE" {
			deleteCategory(w, r)
		}
	})

	// GET localhost:8080/api/produk
	// POST localhost:8080/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProduks(w, r)
		} else if r.Method == "POST" {
			createProduk(w, r)
		}
	})

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	// Swagger UI - serve at /docs/
	http.Handle("/docs/", httpSwagger.WrapHandler)

	// Get port from environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running di :%s\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
