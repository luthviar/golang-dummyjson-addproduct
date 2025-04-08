package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// MarshalFunc is a function type for JSON marshaling
type MarshalFunc func(v interface{}) ([]byte, error)

// Default marshal function that uses the standard json.Marshal
var DefaultMarshal MarshalFunc = json.Marshal

type productService struct {
	apiURL  string
	client  *http.Client
	marshal MarshalFunc // Custom marshal function for testing
}

// NewProductService creates a new productService with default configuration
func NewProductService(apiURL string, client *http.Client) *productService {
	return &productService{
		apiURL:  apiURL,
		client:  client,
		marshal: DefaultMarshal,
	}
}

func (s *productService) AddProduct(p Product) (Product, error) {
	body, err := json.Marshal(p)
	if err != nil {
		return Product{}, fmt.Errorf("failed to marshal product: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, s.apiURL, bytes.NewBuffer(body))
	if err != nil {
		return Product{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return Product{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return Product{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result Product
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Product{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}
