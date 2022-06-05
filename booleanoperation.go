package main

func main() {
	nilaiAkhir := 90
	absensi := 80

	lulusNilaiAkhir := nilaiAkhir >= 80
	lulusAbsensi := absensi >= 80

	lulus := lulusNilaiAkhir && lulusAbsensi

	println(lulus)
}
