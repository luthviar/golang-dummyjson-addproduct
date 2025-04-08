package product

type ProductService interface {
	AddProduct(p Product) (Product, error)
}
