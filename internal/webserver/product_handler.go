package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/upalx/Store-API/internal/entity"
	"github.com/upalx/Store-API/internal/service"
)

type WebProductHandler struct {
	ProductService *service.ProductService
}

func NewWebProductHandler(productService *service.ProductService) *WebProductHandler {
	return &WebProductHandler{ProductService: productService}
}

func (wph *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := wph.ProductService.GetProducts()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (wph *WebProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	product, err := wph.ProductService.GetProduct(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (wph *WebProductHandler) GetProductByCategoryID(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "productID")

	if productID == "" {
		http.Error(w, "productID is required.", http.StatusBadRequest)
		return
	}

	products, err := wph.ProductService.GetProductsByCategoryID(productID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (wph WebProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	println("IS TRYING CREATE A PRODUCT...")

	var product entity.Product

	println("AFTER VAR")

	err := json.NewDecoder(r.Body).Decode(&product)

	println("AFTER ERROR HAND")

	if err != nil {
		println("HAS AN ERROR", w)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	println("IS TRYING GET INTO SERVICE...")
	result, err := wph.ProductService.CreateProduct(product.Name, product.Description, product.CategoryID, product.ImageURL, product.Price)


	if err != nil {
		println("HAS AN ERROR ON CREATE THE PRODUCT ON DB...")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}
