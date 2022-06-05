package main

import "fmt"

func main() {
	counter := 0

	for counter < 10 {
		fmt.Println("Perulangan Ke ", counter)
		counter++
	}

	for counter2 := 0; counter2 < 10; counter2++ {
		fmt.Println("Perulangan Ke ", counter2)
	}

	names := []string{"Agung", "Suprayitno", "Satria", "Jalu"}

	for index, name := range names {
		fmt.Println("index", index, "name", name)
	}

	for _, name := range names {
		fmt.Println("name without index variable", name)
	}

	person := make(map[string]string)
	person["name"] = "Agung"
	person["title"] = "Developer"

	fmt.Println(person)
}
