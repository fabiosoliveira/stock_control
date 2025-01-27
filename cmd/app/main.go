package main

import (
	"net/http"

	"github.com/fabiosoliveira/stock_control/internal/controller"
	"github.com/fabiosoliveira/stock_control/internal/middleware"
	"github.com/fabiosoliveira/stock_control/internal/product"
)

func main() {
	productController := controller.NewProductController(product.NewProductRepositorySqlite())

	mux := http.NewServeMux()

	mux.Handle("/", middleware.CachePage(productController.Index))

	mux.Handle("POST /product", middleware.CachePage(productController.CreateProduct))

	mux.Handle("DELETE /product/{id}", middleware.CachePage(productController.DeleteProduct))

	mux.HandleFunc("GET /product/{id}", productController.GetProduct)

	http.ListenAndServe(":8080", mux)
}
