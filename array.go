package main

import "fmt"

func main() {
	var arrayNames [3]string
	arrayNames[0] = "Agung"
	arrayNames[1] = "Suprayitno"
	arrayNames[2] = "Suprayitno2"
	fmt.Println(arrayNames)

	var values = [5]int{
		1,
		2,
		3,
		4,
		5,
	}

	fmt.Println(values)

	fmt.Println(len(arrayNames))
	fmt.Println(len(values))
}
