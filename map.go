package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Agung",
		"address": "Tangerang",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)

	book["title"] = "Belajar Go-Lang"
	book["author"] = "Agung"
	book["ups"] = "salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
