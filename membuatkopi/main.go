package main

import "fmt"

func main() {
	var (
		kopiTersedia bool
		gulaTersedia bool
		stokKopi     int
		stokGula     int
	)

	// Stok is how "periksa persediaan" / "jika habis" is decided.
	// The pseudocode does not say how many servings are on hand.
	fmt.Print("Stok kopi (jumlah cangkir): ")
	fmt.Scan(&stokKopi)
	fmt.Print("Stok gula (jumlah cangkir): ")
	fmt.Scan(&stokGula)

	kopiTersedia = true
	gulaTersedia = true

	for kopiTersedia == true && gulaTersedia == true {
		fmt.Println("ambil cangkir")
		fmt.Println("masukkan kopi ke cangkir")
		fmt.Println("masukkan gula ke cangkir")
		fmt.Println("tuangkan air panas")
		fmt.Println("aduk menggunakan sendok")
		fmt.Println("sajikan kopi")

		stokKopi = stokKopi - 1
		stokGula = stokGula - 1

		fmt.Println("periksa persediaan kopi")
		fmt.Println("periksa persediaan gula")

		if stokKopi <= 0 {
			kopiTersedia = false
		}

		if stokGula <= 0 {
			gulaTersedia = false
		}

		fmt.Println()
	}

	fmt.Println("Pembuatan kopi dihentikan")
}
