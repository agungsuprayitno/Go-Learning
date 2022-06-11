package main

import "fmt"

func Random() interface{} {
	return "Agung"
}

func main() {
	result := Random()

	//	jika data int code ini akan menimbulkan 'panic', karna error tidak dihandle
	// resultString := result.(string)
	// fmt.Println(resultString)

	//	better pakai switch, karna tidak akan menimbulkan panic
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}

}
