package main

import "fmt"

func main() {
	var (
		celcius    float64
		reamur     float64
		fahrenheit float64
		kelvin     float64
	)

	fmt.Print("Masukkan suhu dalam Celcius: ")
	fmt.Scan(&celcius)

	reamur = celcius * 4 / 5
	fahrenheit = celcius*9/5 + 32
	kelvin = celcius + 273.15

	fmt.Printf("%.2f\n", reamur)
	fmt.Printf("%.2f\n", fahrenheit)
	fmt.Printf("%.2f\n", kelvin)
}
