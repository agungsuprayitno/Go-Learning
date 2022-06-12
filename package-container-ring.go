package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)

	//	push data baru ke Ring
	for i := 0; i < data.Len(); i++ {
		data.Value = "Data " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	//	Print all data on Ring
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})

}
