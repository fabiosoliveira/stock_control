package main

import (
	"net/http"

	"github.com/fabiosoliveira/stock_control/internal/controller"
	"github.com/fabiosoliveira/stock_control/internal/product"
)

func main() {
	productController := controller.NewProductController(product.NewProductRepositorySqlite())

	mux := http.NewServeMux()

	mux.HandleFunc("/", productController.Index)

	mux.HandleFunc("POST /product", productController.CreateProduct)

	mux.HandleFunc("DELETE /product/{id}", productController.DeleteProduct)

	mux.HandleFunc("GET /product/{id}", productController.GetProduct)

	http.ListenAndServe(":8080", mux)
}
