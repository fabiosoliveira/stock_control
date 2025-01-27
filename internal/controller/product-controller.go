package controller

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"sync"

	"github.com/fabiosoliveira/stock_control/internal/product"
)

type ProductController struct {
	repository     product.ProductRepository
	templateIndex  *template.Template
	templateCreate *template.Template
	mu             sync.RWMutex
}

func NewProductController(productRepository product.ProductRepository) *ProductController {
	return &ProductController{
		repository:     productRepository,
		templateIndex:  template.Must(template.ParseFiles("web/template/index.gohtml", "web/template/product-row.gohtml")),
		templateCreate: template.Must(template.ParseFiles("web/template/product-row.gohtml")),
	}
}

func (p *ProductController) Index(w http.ResponseWriter, r *http.Request) {
	p.mu.RLock()
	products, err := p.repository.GetAll()
	p.mu.RUnlock()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	p.templateIndex.Execute(w, products)
}

func (p *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stock, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
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
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		p.mu.Lock()
		err = p.repository.Update(id, name, price, stock)
		p.mu.Unlock()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		product.ID = id
	} else {

		p.mu.Lock()
		product, err = p.repository.Create(name, price, stock)
		p.mu.Unlock()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "text/html")
	if err := p.templateCreate.ExecuteTemplate(w, "ProductRow", product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (p *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	p.mu.Lock()
	err = p.repository.Remove(id)
	p.mu.Unlock()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (p *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	p.mu.RLock()
	product, err := p.repository.GetByID(id)
	p.mu.RUnlock()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
