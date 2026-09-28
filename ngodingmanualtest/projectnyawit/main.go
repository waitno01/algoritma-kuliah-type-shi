package main

import "fmt"

func main() {
	var (
		jumlahpohonsawit int
		jumlahanggaran   int
		lahansawit       int
	)

	fmt.Print("jumlah pohon sawit: ")
	fmt.Scan(&jumlahpohonsawit)

	fmt.Print("jumlah anggaran: ")
	fmt.Scan(&jumlahanggaran)

	fmt.Print("lahan sawit: ")
	fmt.Scan(&lahansawit)

	fmt.Println("jumlahpohonsawit" + "lahansawit") 
}
