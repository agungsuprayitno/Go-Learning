package main

import "fmt"

func main() {

	//	Return value
	name := returnValue()
	fmt.Println(name)
	fmt.Println(returnMultipleValue())
}

func returnValue() string {
	return "Return Agung"
}
