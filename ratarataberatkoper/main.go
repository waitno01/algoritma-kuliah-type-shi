package main

import "fmt"

func main() {
	var (
		N          int
		i          int
		beratKoper float64
		totalBerat float64
		rataRata   float64
	)

	fmt.Print("Masukkan jumlah koper (N): ")
	fmt.Scan(&N)

	i = 1
	totalBerat = 0

	for i <= N {
		fmt.Printf("Masukkan berat koper ke-%d: ", i)
		fmt.Scan(&beratKoper)

		totalBerat = totalBerat + beratKoper
		i = i + 1
	}

	rataRata = totalBerat / float64(N)

	fmt.Println(totalBerat)
	fmt.Println(rataRata)
}
