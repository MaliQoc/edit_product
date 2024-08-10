package main

import (
	"fmt"
	"golesson/project"
)

func main() {
	/* product, _ := project.AddProduct()
	fmt.Println(product) */

	products, _ := project.GetAllProducts()

	for i := 0; i < len(products); i++ {
		fmt.Println(products[i])
	}
}