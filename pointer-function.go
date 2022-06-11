package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	alamat := Address{"Tangerang", "Banten", ""}
	var alamatPointer *Address = &Address{"Lebak", "Banten", ""}
	ChangeCountryToIndonesia(&alamat)
	ChangeCountryToIndonesia(alamatPointer)

	fmt.Println(alamat)
	fmt.Println(alamatPointer)
}
