package main

import (
	"Microservices/product-api/sdk/client"
	"Microservices/product-api/sdk/client/products"
	"fmt"
	"testing"
)

func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("Localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := products.NewListProductsParams()

	prod, err := c.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(prod)
}
