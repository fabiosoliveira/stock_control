package main

import (
	"html/template"
	"net/http"

	"github.com/fabiosoliveira/stock_control/internal/product"
)

func main() {
	productRepository := product.NewProductRepository()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("web/index.html")

		products, err := productRepository.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, products)
	})

	http.ListenAndServe(":8080", mux)
}
