package main

import "fmt"

func main() {
	name := "Bukan Agung"
	nameAgung := "Agung"

	if name == "Agung" {
		fmt.Println(name)
	} else if nameAgung == "Agung" {
		fmt.Println(nameAgung)
	} else {
		fmt.Println(name + " Bukan Agung")
	}

	if lengthName := len(name); lengthName > 10 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
