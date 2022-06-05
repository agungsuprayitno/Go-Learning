package main

import "fmt"

func main() {
	arrayMonths := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	slice1 := arrayMonths[3:6]
	slice1[1] = "GantiSlice1"

	slice2 := arrayMonths[1:]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(arrayMonths)

	arrayDays := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	sliceDays := []string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	fmt.Println(arrayDays)
	fmt.Println(sliceDays)

}
