package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("Agung Suprayitno", "agung"))
	fmt.Println(strings.Split("Agung Suprayitno", " "))
	fmt.Println(strings.ToLower("Agung Suprayitno"))
	fmt.Println(strings.ToUpper("Agung Suprayitno"))
	fmt.Println(strings.Trim("	Agung Suprayitno	", " "))
	fmt.Println(strings.Trim("agung agung agung", " eko"))
	fmt.Println(strings.ToTitle("agung agung agung"))

}
