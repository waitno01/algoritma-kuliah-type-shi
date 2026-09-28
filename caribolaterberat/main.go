package main

import "fmt"

func main() {
	var (
		bolaA        float64
		bolaB        float64
		bolaC        float64
		bolaD        float64
		kandidat1    float64
		kandidat2    float64
		bolaTerberat float64
	)

	fmt.Print("Masukkan berat bola A: ")
	fmt.Scan(&bolaA)
	fmt.Print("Masukkan berat bola B: ")
	fmt.Scan(&bolaB)
	fmt.Print("Masukkan berat bola C: ")
	fmt.Scan(&bolaC)
	fmt.Print("Masukkan berat bola D: ")
	fmt.Scan(&bolaD)

	if bolaA > bolaB {
		kandidat1 = bolaA
	} else {
		kandidat1 = bolaB
	}

	if bolaC > bolaD {
		kandidat2 = bolaC
	} else {
		kandidat2 = bolaD
	}

	if kandidat1 > kandidat2 {
		bolaTerberat = kandidat1
	} else {
		bolaTerberat = kandidat2
	}

	fmt.Println(bolaTerberat)
}
