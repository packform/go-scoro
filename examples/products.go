package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	scoro "github.com/lxmx/go-scoro"
)

func main() {
	company := flag.String("company", "", "Company id")
	apiKey := flag.String("api_key", "", "Scoro API key")
	flag.Parse()

	if *company == "" || *apiKey == "" {
		fmt.Println("Please specify company and api_key")
		return
	}

	credentials := scoro.Credentials{ApiKey: *apiKey, CompanyID: *company}

	fmt.Println("List products: ")
	listProducts(credentials)

	fmt.Println("Create product: ")
	product := createProduct(credentials)

	fmt.Println("Modify product: ")
	product = modifyProduct(credentials, product)

	fmt.Println("Remove product: ")
	removeProduct(credentials, product)
}

func listProducts(credentials scoro.Credentials) {
	products, err := scoro.Products(credentials).List(nil, 0, 3)
	if err != nil {
		panic(err)
	}

	printObject(products)
}

func createProduct(credentials scoro.Credentials) scoro.Product {
	product := scoro.Product{
		Code:         "435345",
		Description:  scoro.MakeStrings("go-scoro example product description", scoro.DefaultLang),
		IsActive:     scoro.Bool{Value: true},
		ModifiedDate: scoro.Time{Time: time.Now()},
		Names:        scoro.MakeStrings("Go scoro product", scoro.DefaultLang),
	}

	result, err := scoro.Products(credentials).Modify(product)
	if err != nil {
		panic(err)
	}

	printObject(result)
	return *result
}

func modifyProduct(credentials scoro.Credentials, product scoro.Product) scoro.Product {
	product.Description = scoro.MakeStrings("go-scoro changed product description", scoro.DefaultLang)

	result, err := scoro.Products(credentials).Modify(product)
	if err != nil {
		panic(err)
	}

	printObject(result)
	return *result
}

func removeProduct(credentials scoro.Credentials, product scoro.Product) {
	err := scoro.Products(credentials).Delete(*product.Id)
	if err != nil {
		panic(err)
	}

	fmt.Println("OK!")
}

func printObject(obj interface{}) {
	jsObj, _ := json.MarshalIndent(obj, "", "\t")
	fmt.Println(string(jsObj))
}
