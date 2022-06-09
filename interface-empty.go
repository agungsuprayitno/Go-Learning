package main

import "fmt"

type EmptyInterface interface {
}

func sayHello(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Hello"
	}
}

func main() {
	fmt.Println(sayHello(34))
}
