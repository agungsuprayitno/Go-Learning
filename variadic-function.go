package main

import "fmt"

func main() {
	//	Variadic function
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println(VariadicFunction(10, 20, 30, 40, 50))
	fmt.Println(VariadicFunction(numbers...))
}

func VariadicFunction(numbers ...int) int {
	total := 1
	for _, value := range numbers {
		total += value
	}

	return total
}
