package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1 // Pass by value
	address2 := &address1 // pointer (Pass by reference)

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}
	address5 := address1

	var address3 Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	var address4 *Address = &address3

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)
	fmt.Println(address5)

	var address6 *Address = new(Address)

	fmt.Println(address6)
}
