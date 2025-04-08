package main

import (
	"dummyjson/product"
	"fmt"
	"net/http"
)

func main() {
	svc := product.NewProductService("https://dummyjson.com/products/add", http.DefaultClient)

	newProduct := product.Product{
		Title:       "BMW Pencil 11",
		Description: "A luxury pencil by BMW 12",
		Price:       1213,
		Brand:       "BMW 14",
		Category:    "stationery 15",
	}

	added, err := svc.AddProduct(newProduct)
	if err != nil {
		fmt.Println("Failed to add product:", err)
		return
	}

	fmt.Println("Product added successfully:", added)
}
