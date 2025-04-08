package product

type MockProductService struct {
	AddProductFunc func(p Product) (Product, error)
}

func (m *MockProductService) AddProduct(p Product) (Product, error) {
	return m.AddProductFunc(p)
}
