package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func blackListAdmin(name string) bool {
	return name == "admin"
}

func blackListroot(name string) bool {
	return name == "root"
}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)

	registerUser("root", func(name string) bool {
		return name == "admin"
	})
}
