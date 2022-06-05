package main

import "fmt"

func main() {

	//	Return value
	name := returnValue()
	fmt.Println(name)
	fmt.Println(ReturnNamedMultipleValue())
}

func ReturnNamedMultipleValue() (firstName string, lastName string) {
	firstName = "Agung"
	lastName = "Suprayitno"
	return
}
