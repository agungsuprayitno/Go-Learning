package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", boolean)
	}

	fmt.Println(strconv.Itoa(23))
	fmt.Println(strconv.Atoi("2000000"))
}
