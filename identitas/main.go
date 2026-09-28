package main

import "fmt"

func main() {
	var (
		nim          string
		namaLengkap  string
		semester     int
		programStudi string
		fakultas     string
		kampus       string
		kelompok     string
		asal         string
	)

	nim = "110052600007"
	namaLengkap = "Yuro Zezario Widodo"
	semester = 1
	programStudi = "Informatika"
	fakultas = "Informatika"
	kampus = "Telkom University Jakarta"
	kelompok = "Gugus 45 Kalimalang"
	asal = "Depok"

	fmt.Println("NIM           :", nim)
	fmt.Println("Nama lengkap  :", namaLengkap)
	fmt.Println("Semester      :", semester)
	fmt.Println("Program studi :", programStudi)
	fmt.Println("Fakultas      :", fakultas)
	fmt.Println("Kampus        :", kampus)
	fmt.Println("Kelompok      :", kelompok)
	fmt.Println("Asal          :", asal)
}
