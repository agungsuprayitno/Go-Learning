package main

func main() {
	type NoKTP string

	var ktp NoKTP = "111111111"

	println(ktp)
	println(NoKTP("223434324234"))
}
