package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() ([]Product, error) {
	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var products []Product
	json.Unmarshal(bodyBytes, &products)
	return products, nil
}

func AddProduct() (Product, error) {
	product := Product{ProductName: "Telephone", CategoryId: 1, UnitPrice: 6000.0}
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		return Product{}, err
	}

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	if err != nil {
		return Product{}, err
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var productResponse Product
	json.Unmarshal(bodyBytes, &productResponse)
	return productResponse, nil
}

func UpdateProduct(productId int, updatedProduct Product) (Product, error) {
	jsonProduct, err := json.Marshal(updatedProduct)
	if err != nil {
		return Product{}, err
	}

	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:3000/products/%d", productId), bytes.NewBuffer(jsonProduct))
	if err != nil {
		return Product{}, err
	}
	request.Header.Set("Content-Type", "application/json;charset=utf-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return Product{}, err
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var productResponse Product
	json.Unmarshal(bodyBytes, &productResponse)
	return productResponse, nil
}

func DeleteProduct(productId int) error {
	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:3000/products/%d", productId), nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete product, status code: %d", response.StatusCode)
	}

	return nil
}
