package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Agung")
	data.PushBack("Suprayitno")
	data.PushBack("Satria")
	data.PushBack("Jalu")
	data.PushFront("Sofi")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	//	iterasi data dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	//	iterasi data dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}

}
