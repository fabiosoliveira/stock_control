package product

type ProductRepository struct{}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (p *ProductRepository) GetAll() ([]Product, error) {
	rows, err := DB.Query("SELECT * FROM products")
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

func (p *ProductRepository) Create(name string, price float64, stock int) (Product, error) {
	result, err := DB.Exec("INSERT INTO products (name, price, stock) VALUES (?, ?, ?)", name, price, stock)
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

func (p *ProductRepository) Remove(id int) error {
	_, err := DB.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}
