package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Package OS")
	args := os.Args
	fmt.Println("Argument", args)

	// sebagai contoh cara run argument Os
	//	go run package-os.go agung suprayitno

	fmt.Println(args[0]) // akan muncul default argument bawaan operating system
	fmt.Println(args[1]) // akan muncul Agung
	fmt.Println(args[2]) // akan muncul Suprayitno

	username := os.Getenv("USERNAME") // username operating system

	fmt.Println(username)
}
