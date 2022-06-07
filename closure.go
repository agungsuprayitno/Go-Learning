package main

import "fmt"

func main() {
	counter := 0

	increement := func() {
		fmt.Println("Increement", counter)
		counter++
	}

	increement()
	increement()

	fmt.Println(counter)
}
