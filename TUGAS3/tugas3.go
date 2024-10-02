package main

import "fmt"

func main() {
	var r, luas float64
	const PHI float64 = 3.14

	fmt.Print("Masukkan jari-jari: ")
	fmt.Scanln(&r)

	luas = PHI * r * r

	fmt.Printf("Luas lingkaran: %.2f\n", luas)
}
