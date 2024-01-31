package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/upalx/Store-API/internal/database"
	"github.com/upalx/Store-API/internal/service"
	"github.com/upalx/Store-API/internal/webserver"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3006)/store")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	webProductHandler := webserver.NewWebProductHandler(productService)

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/category/{id}", webCategoryHandler.GetCategory)
	router.Get("/categories", webCategoryHandler.GetCategories)
	router.Post("/category", webCategoryHandler.CreateCategory)

	router.Get("/product/{id}", webProductHandler.GetProduct)
	router.Get("/products", webProductHandler.GetProducts)
	router.Get("/product/category/{category_id}", webProductHandler.GetProductByCategoryID)
	router.Post("/product", webProductHandler.CreateProduct)

	fmt.Println("Server is running on port 8080")

	http.ListenAndServe(":8080", router)
}
