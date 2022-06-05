package main

import "fmt"

func main() {
	name := "Agung"
	switch name {
	case "Agung":
		fmt.Println("Agung")
	case "Suprayitno":
		fmt.Println("Suprayitno")
	case "Bukan Agung":
		fmt.Println("Bukan Agung")
	default:
		fmt.Println("Halo")

	}

	switch nameLength := "Agung"; len(nameLength) > 10 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}
}
