package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi() {
	fmt.Println("Hello, My Name is", customer.Name)
}

func (customer Customer) sayHuu() {
	fmt.Println("Huuu From ", customer.Name)
}

func main() {
	var agung Customer
	agung.Name = "Agung Suprayitno"
	agung.Address = "Tangerang"
	agung.Age = 30

	fmt.Println(agung.Name)
	fmt.Println(agung.Address)
	fmt.Println(agung.Age)

	agung1 := Customer{
		Name:    "Agung 1",
		Address: "Tangerang 1",
		Age:     30,
	}

	agung2 := Customer{"agung 2", "Tangerang 2", 30}

	fmt.Println(agung1)
	fmt.Println(agung2)

	rully := Customer{
		Name: "Rully",
	}

	rully.sayHi()

	rully.sayHuu()

}
