package main

import "fmt"

func main() {
	var (
		bola1 int
		bola2 int
		bola3 int
		temp  int
	)

	fmt.Print("Angka pada bola pertama: ")
	fmt.Scan(&bola1)
	fmt.Print("Angka pada bola kedua: ")
	fmt.Scan(&bola2)
	fmt.Print("Angka pada bola ketiga: ")
	fmt.Scan(&bola3)

	if bola1 > bola2 {
		temp = bola1
		bola1 = bola2
		bola2 = temp
	}

	if bola1 > bola3 {
		temp = bola1
		bola1 = bola3
		bola3 = temp
	}

	if bola2 > bola3 {
		temp = bola2
		bola2 = bola3
		bola3 = temp
	}

	fmt.Println(bola1, bola2, bola3)
}
