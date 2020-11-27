package main

import (
	"assignment1/domain"
	"assignment1/mapstore"
	"fmt"
)

var controller *CustomerController

func init() {
	controller = &CustomerController{
		store: mapstore.NewMapStore()}
}

func main() {

	shiras := domain.Customer{ID: "cust101", Name: "Shiras", Email: "shiras.hassan@gmail.com"}
	ajeesh := domain.Customer{ID: "cust102", Name: "Ajeesh", Email: "ajeesh@gmail.com"}
	navindra := domain.Customer{ID: "cust103", Name: "Navindra Kushwaha", Email: "navindra.kushwaha@gmail.com"}
	fawaz := domain.Customer{ID: "cust104", Name: "Fawaz", Email: "fawazabdullakp@gmail.com"}
	vineet := domain.Customer{ID: "cust105", Name: "Vineet", Email: "vineet@gmail.com"}

	fmt.Println("----------------------------------------------")
	fmt.Println("Add 5 Customers (Verifying create and getall functionality)")

	controller.Add(shiras)   // Create new Customer
	controller.Add(ajeesh)   // Create new Customer
	controller.Add(navindra) // Create new Customer
	controller.Add(fawaz)    // Create new Customer
	controller.Add(vineet)   // Create new Customer
	customers, err := controller.GetAll()
	fmt.Println("Current active customers")
	for k, v := range customers {
		fmt.Printf("  (%d) Customer details: %v \n", k, v)
	}

	fmt.Println("----------------------------------------------")
	fmt.Println("Verifying validations functionality")

	proxy := shiras
	fmt.Println("1. Adding same customer.")
	err = controller.Add(proxy) // Create new Customer with same id.
	fmt.Println("  response:", err)

	proxy.ID = "1"
	fmt.Println("2. Modifying invalid customer.")
	err = controller.Modify(proxy)
	fmt.Println("  response:", err)

	fmt.Println("3. Removing invalid customer.")
	controller.Remove(proxy)
	fmt.Println("  response:", err)

	fmt.Println("4. Querying invalid customer.")
	_, err = controller.Get(proxy.ID)
	fmt.Println("  response:", err)

	fmt.Println("----------------------------------------------")
	fmt.Println("Verifying update functionality")

	shiras.Name = "Shiras Hassan"
	controller.Modify(shiras)
	proxy, _ = controller.Get(shiras.ID)
	fmt.Println("   Modified value:", proxy)
	fmt.Println("----------------------------------------------")
	fmt.Println("Verifying delete functionality")
	fmt.Println("Removing a customer")
	controller.Remove(shiras)
	fmt.Println("Querying removed customer")
	_, err = controller.Get(shiras.ID)
	fmt.Println(err)
	fmt.Println("----------------------------------------------")
	fmt.Println("Verifying get all functionality")

	customers, err = controller.GetAll()
	fmt.Println("Active customers")
	for k, v := range customers {
		fmt.Printf("(%d) Customer details: %v \n", k, v)
	}
}
