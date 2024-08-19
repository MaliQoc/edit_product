package main

import (
	"fmt"
	"golesson/project"
	"log"
)

func main() {

	products, err := project.GetAllProducts()
	if err != nil {
		log.Fatalf("Ürünler alınamadı: %v", err)
	}

	fmt.Println("Tüm Ürünler:")
	for i := 0; i < len(products); i++ {
		fmt.Printf("ID: %d, Ürün Adı: %s, Kategori ID: %d, Fiyat: %.2f\n", products[i].Id, products[i].ProductName, products[i].CategoryId, products[i].UnitPrice)
	}

	/*
		product, err := project.AddProduct()
		if err != nil {
			log.Fatalf("Ürün eklenemedi: %v", err)
		}
		fmt.Printf("\nYeni eklenen ürün: ID: %d, Ürün Adı: %s, Kategori ID: %d, Fiyat: %.2f\n", product.Id, product.ProductName, product.CategoryId, product.UnitPrice)


		updatedProduct := project.Product{ProductName: "Updated Telephone", CategoryId: 1, UnitPrice: 5500.0}
		updatedProductResponse, err := project.UpdateProduct(product.Id, updatedProduct)
		if err != nil {
			log.Fatalf("Ürün güncellenemedi: %v", err)
		}
		fmt.Printf("\nGüncellenen ürün: ID: %d, Ürün Adı: %s, Kategori ID: %d, Fiyat: %.2f\n", updatedProductResponse.Id, updatedProductResponse.ProductName, updatedProductResponse.CategoryId, updatedProductResponse.UnitPrice)


		err = project.DeleteProduct(product.Id)
		if err != nil {
			log.Fatalf("Ürün silinemedi: %v", err)
		} else {
			fmt.Printf("\nÜrün başarıyla silindi: ID: %d\n", product.Id)
		}
	*/
}
