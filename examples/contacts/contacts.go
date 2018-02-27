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
	subdomain := flag.String("subdomain", "", "Subdomain id")
	apiKey := flag.String("api_key", "", "Scoro API key")
	flag.Parse()

	if *company == "" || *apiKey == "" || *subdomain == "" {
		fmt.Println("Please specify company, api_key and subdomain")
		return
	}

	credentials := scoro.Credentials{ApiKey: *apiKey, CompanyID: *company, Subdomain: *subdomain}

	fmt.Println("List contacts: ")
	listContacts(credentials)

	fmt.Println("Create contact: ")
	contact := createContact(credentials)

	fmt.Println("Modify contact: ")
	contact = modifyContact(credentials, contact)

	fmt.Println("Remove contact: ")
	removeContact(credentials, contact)
}

func listContacts(credentials scoro.Credentials) {
	contacts, err := scoro.Contacts(credentials).List(nil, 0, 3)
	if err != nil {
		panic(err)
	}

	printObject(contacts)
}

func createContact(credentials scoro.Credentials) scoro.Contact {
	contact := scoro.Contact{
		Name:         "Viktor",
		Lastname:     "Ladochkin",
		Birthday:     scoro.Date{Time: time.Date(1984, 4, 3, 0, 0, 0, 0, time.UTC)},
		ContactType:  "person",
		IsClient:     scoro.Bool{Value: true},
		ModifiedDate: scoro.Time{Time: time.Now()},
		Sex:          "M",
	}

	result, err := scoro.Contacts(credentials).Modify(contact)
	if err != nil {
		panic(err)
	}

	printObject(result)
	return *result
}

func modifyContact(credentials scoro.Credentials, contact scoro.Contact) scoro.Contact {
	contact.Addresses = []scoro.Address{
		scoro.Address{
			City:    "Tomsk",
			Country: "Russia",
		},
	}

	result, err := scoro.Contacts(credentials).Modify(contact)
	if err != nil {
		panic(err)
	}

	printObject(result)
	return *result
}

func removeContact(credentials scoro.Credentials, contact scoro.Contact) {
	err := scoro.Contacts(credentials).Delete(*contact.ContactID)
	if err != nil {
		panic(err)
	}

	fmt.Println("OK!")
}

func printObject(obj interface{}) {
	jsObj, _ := json.MarshalIndent(obj, "", "\t")
	fmt.Println(string(jsObj))
}
