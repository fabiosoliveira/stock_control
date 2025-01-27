package product

import "github.com/fabiosoliveira/stock_control/internal/config"

type ProductRepositorySqlite struct{}

func NewProductRepositorySqlite() *ProductRepositorySqlite {
	return &ProductRepositorySqlite{}
}

func (p *ProductRepositorySqlite) GetAll() ([]Product, error) {
	rows, err := config.DB.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (p *ProductRepositorySqlite) GetByID(id int) (Product, error) {
	var product Product
	err := config.DB.QueryRow("SELECT * FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
	return product, err
}

func (p *ProductRepositorySqlite) Create(name string, price float64, stock int) (Product, error) {
	result, err := config.DB.Exec("INSERT INTO products (name, price, stock) VALUES (?, ?, ?)", name, price, stock)
	if err != nil {
		return Product{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Product{}, err
	}

	product := Product{
		ID:    int(id),
		Name:  name,
		Price: price,
		Stock: stock,
	}

	return product, nil
}

func (p *ProductRepositorySqlite) Update(id int, name string, price float64, stock int) error {
	_, err := config.DB.Exec("UPDATE products SET name = ?, price = ?, stock = ? WHERE id = ?", name, price, stock, id)
	return err
}

func (p *ProductRepositorySqlite) Remove(id int) error {
	_, err := config.DB.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}
