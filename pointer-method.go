package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr ." + man.Name
}

func main() {
	agung := Man{"Agung"}
	agung.Married()

	fmt.Println(agung.Name)
}
