package main

import "fmt"

func main() {
	// constanta tidak akan di komplain oleh go, walaupun tidak dipakai. karna kemungkinana akan dipakai di file lain
	const name = "Agung"
	const (
		firstName = "Agung"
		lastName  = "Suprayitno"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
