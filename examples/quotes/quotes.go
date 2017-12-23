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

	fmt.Println("Create product: ")
	product := createProduct(credentials)

	fmt.Println("Create quote: ")
	quote := createQuote(credentials, product)

	fmt.Println("Remove quote: ")
	removeQuote(credentials, quote)

	fmt.Println("Remove product: ")
	removeProduct(credentials, product)
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

func createQuote(credentials scoro.Credentials, product scoro.Product) scoro.Quote {
	quote := scoro.Quote{
		Currency:    "USD",
		Description: "Sample description",
		OwnerID:     1,
		CompanyID:   36,
		Lines: []scoro.QuoteLine{
			scoro.QuoteLine{
				Amount:    scoro.NewDecimalFromFloat(100.12),
				Comment:   scoro.MakeStrings("Test comment", scoro.DefaultLang),
				ProductID: *product.Id,
				Sum:       scoro.NewDecimalFromFloat(1001.2),
				UnitPrice: scoro.NewDecimalFromFloat(10),
			},
		},
	}

	result, err := scoro.Quotes(credentials).Modify(quote)
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

func removeQuote(credentials scoro.Credentials, quote scoro.Quote) {
	err := scoro.Quotes(credentials).Delete(*quote.Id)
	if err != nil {
		panic(err)
	}

	fmt.Println("OK!")
}

func printObject(obj interface{}) {
	jsObj, _ := json.MarshalIndent(obj, "", "\t")
	fmt.Println(string(jsObj))
}
