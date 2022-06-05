package main

import "fmt"

func main() {
	//	function as Parameter
	fmt.Println(functionAsParameter("Agung", spamFilter))
}

type Filter func(string) string

func functionAsParameter(name string, filter Filter) (string, string) {
	return "Hello", filter(name)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
