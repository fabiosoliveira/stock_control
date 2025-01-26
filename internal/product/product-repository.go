package product

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetByID(id int) (Product, error)
	Create(name string, price float64, stock int) (Product, error)
	Update(id int, name string, price float64, stock int) error
	Remove(id int) error
}
