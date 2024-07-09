package main

import "fmt"

type Products struct {
	ProductID string
	Stock     float64
}

func main() {
	var (
		products            []Products
		productMap          = make(map[string]bool)
		productWithStockMap = make(map[string]float64)
	)

	n := []Products{
		{
			ProductID: "Product-1",
			Stock:     10,
		},
		{
			ProductID: "Product-2",
			Stock:     20,
		},
		{
			ProductID: "Product-1",
			Stock:     30,
		},
		{
			ProductID: "Product-3",
			Stock:     40,
		},
		{
			ProductID: "Product-3",
			Stock:     50,
		},
		{
			ProductID: "Product-4",
			Stock:     60,
		},
	}

	for _, v := range n {
		if _, ok := productMap[v.ProductID]; !ok {
			productMap[v.ProductID] = true

			products = append(products, v)
		} else {
			productWithStockMap[v.ProductID] = v.Stock
		}
	}

	for i := range products {
		if stock, ok := productWithStockMap[products[i].ProductID]; ok {
			products[i].Stock += stock
		}
	}

	// result
	// Product-1 = 40
	// Product-2 = 20
	// Product-3 = 90
	// Product-4 = 60
	fmt.Println("result: ", products)
}
