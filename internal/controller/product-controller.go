package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/fabiosoliveira/stock_control/internal/product"
)

type ProductController struct {
	product.ProductRepository
}

func NewProductController(productRepository product.ProductRepository) *ProductController {
	return &ProductController{ProductRepository: productRepository}
}

func (repository *ProductController) Index(w http.ResponseWriter, r *http.Request) {
	products, err := repository.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("web/template/index.gohtml", "web/template/product-row.gohtml"))

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, products)
}

func (repository *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
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

		err = repository.Update(id, name, price, stock)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		product.ID = id
	} else {

		product, err = repository.Create(name, price, stock)
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
}

func (repository *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = repository.Remove(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (repository *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	product, err := repository.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
