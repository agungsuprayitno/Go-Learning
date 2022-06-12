package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "put your database host")
	var username *string = flag.String("username", "root", "put your database username")
	var password *string = flag.String("password", "root", "put your database password")

	flag.Parse()

	fmt.Println(*host, *username, *password)

	fmt.Println("Host:", *host)
	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)

	// contoh command
	//	go run package.go -host=agung -username=agung -password=123456
}
