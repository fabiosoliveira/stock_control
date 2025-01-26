package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/fabiosoliveira/stock_control/internal/product"
)

func main() {
	productRepository := product.NewProductRepository()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		products, err := productRepository.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tmpl := template.Must(template.ParseFiles("web/template/index.gohtml", "web/template/product-row.gohtml"))

		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, products)
	})

	mux.HandleFunc("POST /product", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		stock, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		idParam := r.FormValue("id")

		product := product.Product{
			Name:  name,
			Price: price,
			Stock: stock,
		}

		if idParam != "" {
			id, err := strconv.Atoi(idParam)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = productRepository.Update(id, name, price, stock)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			product.ID = id
		} else {

			product, err = productRepository.Create(name, price, stock)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		tmpl := template.Must(template.ParseFiles("web/template/product-row.gohtml"))

		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.ExecuteTemplate(w, "ProductRow", product); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("DELETE /product/{id}", func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = productRepository.Remove(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})

	mux.HandleFunc("GET /product/{id}", func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		product, err := productRepository.GetByID(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// json
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)

	})

	http.ListenAndServe(":8080", mux)
}
